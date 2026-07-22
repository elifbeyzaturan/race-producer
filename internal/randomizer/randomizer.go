package randomizer

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/elifbeyzaturan/race-producer/internal/model"
)

const tickDurationSec = 4.0

func NewRace(r *rand.Rand) model.Race {
	discipline := Disciplines[r.Intn(len(Disciplines))]
	lanes := assignLanes(r, discipline)

	return model.Race{
		ID:         fmt.Sprintf("race-%d", time.Now().UnixMilli()),
		Status:     "IN_PROGRESS",
		StartedAt:  time.Now(),
		Discipline: discipline,
		Lanes:      lanes,
		Tick:       0,
		ElapsedMs:  0,
	}
}

func assignLanes(r *rand.Rand, discipline model.Discipline) []model.Lane {
	countries := make([]model.Country, len(Countries))
	copy(countries, Countries)
	r.Shuffle(len(countries), func(i, j int) {
		countries[i], countries[j] = countries[j], countries[i]
	})

	lanes := make([]model.Lane, 8)
	for i, country := range countries {
		lane := model.Lane{
			Number:         i + 1,
			Country:        country.Name,
			Position:       i + 1,
			TotalDistance:  discipline.Distance,
			LegFinishTimes: []float64{},
		}

		if discipline.IsRelay {
			team := make([]model.Swimmer, len(country.Swimmers))
			copy(team, country.Swimmers)
			r.Shuffle(len(team), func(a, b int) {
				team[a], team[b] = team[b], team[a]
			})
			lane.Team = team
			first := team[0]
			lane.CurrentLegSwimmer = &first
		} else {
			swimmer := country.Swimmers[r.Intn(len(country.Swimmers))]
			lane.Swimmer = &swimmer
		}

		lanes[i] = lane
	}
	return lanes
}

func Tick(race *model.Race, r *rand.Rand) {
	race.Tick++
	race.ElapsedMs += int64(tickDurationSec * 1000)

	for i := range race.Lanes {
		if race.Lanes[i].Finished {
			continue
		}

		if race.Discipline.IsRelay {
			tickRelay(&race.Lanes[i], r, race.Discipline, race.ElapsedMs)
		} else {
			tickIndividual(&race.Lanes[i], r, race.Discipline.Style, race.ElapsedMs)
		}
	}

	updatePositions(race)
}

func tickIndividual(lane *model.Lane, r *rand.Rand, style string, elapsedMs int64) {
	speed := randomSpeed(r, style)
	covered := speed * tickDurationSec

	lane.DistanceCovered += covered

	if lane.DistanceCovered >= float64(lane.TotalDistance) {
		overshoot := lane.DistanceCovered - float64(lane.TotalDistance)
		timeToFinish := (covered - overshoot) / speed
		lane.FinishTimeMs = elapsedMs - int64(tickDurationSec*1000) + int64(timeToFinish*1000)
		lane.DistanceCovered = float64(lane.TotalDistance)
		lane.Finished = true
	}
}

// tickRelay handles mid-tick leg transitions:
// when a swimmer finishes their leg within a tick, the next swimmer
// immediately starts with the remaining time in that tick.
func tickRelay(lane *model.Lane, r *rand.Rand, discipline model.Discipline, elapsedMs int64) {
	remainingSec := tickDurationSec
	tickStartMs := elapsedMs - int64(tickDurationSec*1000)

	for remainingSec > 0 && lane.CurrentLegIndex < discipline.Legs {
		style := getLegStyle(discipline, lane.CurrentLegIndex)
		speed := randomSpeed(r, style)

		legDistRemaining := float64(discipline.LegDist) - lane.CurrentLegDist
		timeToFinishLeg := legDistRemaining / speed

		if timeToFinishLeg <= remainingSec {
			lane.DistanceCovered += legDistRemaining
			lane.CurrentLegDist = 0

			legFinishMs := tickStartMs + int64((tickDurationSec-remainingSec+timeToFinishLeg)*1000)
			lane.LegFinishTimes = append(lane.LegFinishTimes, float64(legFinishMs)/1000.0)

			remainingSec -= timeToFinishLeg
			lane.CurrentLegIndex++

			if lane.CurrentLegIndex < discipline.Legs {
				next := lane.Team[lane.CurrentLegIndex]
				lane.CurrentLegSwimmer = &next
			}
		} else {
			lane.CurrentLegDist += speed * remainingSec
			lane.DistanceCovered += speed * remainingSec
			remainingSec = 0
		}
	}

	if lane.CurrentLegIndex >= discipline.Legs && !lane.Finished {
		lane.DistanceCovered = float64(lane.TotalDistance)
		lane.FinishTimeMs = elapsedMs - int64(remainingSec*1000)
		lane.CurrentLegSwimmer = nil
		lane.Finished = true
	}
}

func getLegStyle(discipline model.Discipline, legIndex int) string {
	if discipline.Style == "medley" {
		return MedleyRelayOrder[legIndex%len(MedleyRelayOrder)]
	}
	return discipline.Style
}

func updatePositions(race *model.Race) {
	sorted := make([]model.Lane, len(race.Lanes))
	copy(sorted, race.Lanes)

	sort.Slice(sorted, func(i, j int) bool {
		li, lj := sorted[i], sorted[j]
		if li.Finished && lj.Finished {
			return li.FinishTimeMs < lj.FinishTimeMs
		}
		if li.Finished {
			return true
		}
		if lj.Finished {
			return false
		}
		return li.DistanceCovered > lj.DistanceCovered
	})

	posMap := make(map[int]int)
	for pos, lane := range sorted {
		posMap[lane.Number] = pos + 1
	}

	for i := range race.Lanes {
		race.Lanes[i].Position = posMap[race.Lanes[i].Number]
	}
}

func IsFinished(race *model.Race) bool {
	for _, lane := range race.Lanes {
		if !lane.Finished {
			return false
		}
	}
	return true
}

func randomSpeed(r *rand.Rand, style string) float64 {
	sr := SpeedRanges[style]
	return sr[0] + r.Float64()*(sr[1]-sr[0])
}
