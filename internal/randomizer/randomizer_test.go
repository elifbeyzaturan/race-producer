package randomizer

import (
	"math/rand"
	"testing"
)

func TestHasEightLanes(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	race := NewRace(r)

	if len(race.Lanes) != 8 {
		t.Errorf("expected 8 lanes, got %d", len(race.Lanes))
	}
}

func TestStatusIsInProgress(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	race := NewRace(r)

	if race.Status != "IN_PROGRESS" {
		t.Errorf("expected IN_PROGRESS, got %s", race.Status)
	}
}

func TestTickAndElapsedMsStartAtZero(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	race := NewRace(r)

	if race.Tick != 0 {
		t.Errorf("expected tick 0, got %d", race.Tick)
	}
	if race.ElapsedMs != 0 {
		t.Errorf("expected elapsedMs 0, got %d", race.ElapsedMs)
	}
}

func TestTickIncreasesDistance(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	race := NewRace(r)
	before := race.Lanes[0].DistanceCovered

	Tick(&race, r)

	if race.Lanes[0].DistanceCovered <= before {
		t.Errorf("expected distance to increase after tick")
	}
}

func TestTickUpdatesElapsedMs(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	race := NewRace(r)

	Tick(&race, r)

	if race.ElapsedMs != 4000 {
		t.Errorf("expected 4000ms after one tick, got %d", race.ElapsedMs)
	}
}

func TestFinishedWhenAllLanesDone(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	race := NewRace(r)

	for i := range race.Lanes {
		race.Lanes[i].Finished = true
	}

	if !IsFinished(&race) {
		t.Errorf("expected race to be finished when all lanes are done")
	}
}

func TestNotFinishedWhenOneLaneRemains(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	race := NewRace(r)

	for i := range race.Lanes {
		race.Lanes[i].Finished = true
	}
	race.Lanes[0].Finished = false

	if IsFinished(&race) {
		t.Errorf("expected race to not be finished when one lane remains")
	}
}

func TestPositionsUpdatedCorrectly(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	race := NewRace(r)

	race.Lanes[0].DistanceCovered = 90.0
	race.Lanes[1].DistanceCovered = 50.0
	race.Lanes[2].DistanceCovered = 70.0

	Tick(&race, r)

	if race.Lanes[0].Position >= race.Lanes[2].Position {
		t.Errorf("lane with more distance should have better position")
	}
}

func TestRelayLegTransition(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	discipline := Disciplines[12] // 4x100m Freestyle Relay
	lanes := assignLanes(r, discipline)

	lane := &lanes[0]
	lane.CurrentLegDist = 99.0 // 1m kaldı, bu tick'te bacak bitmeli

	tickRelay(lane, r, discipline, 4000)

	if len(lane.LegFinishTimes) == 0 {
		t.Errorf("expected leg to finish and record finish time")
	}
	if lane.CurrentLegIndex != 1 {
		t.Errorf("expected leg index to advance to 1, got %d", lane.CurrentLegIndex)
	}
}

func TestIndividualLaneHasSwimmer(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	discipline := Disciplines[0] // Freestyle 50m - individual
	lanes := assignLanes(r, discipline)

	for _, lane := range lanes {
		if lane.Swimmer == nil {
			t.Errorf("expected swimmer in lane %d, got nil", lane.Number)
		}
	}
}

func TestRelayLaneHasFourSwimmers(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	discipline := Disciplines[12] // 4x100m Freestyle Relay
	lanes := assignLanes(r, discipline)

	for _, lane := range lanes {
		if len(lane.Team) != 4 {
			t.Errorf("expected 4 swimmers in relay lane %d, got %d", lane.Number, len(lane.Team))
		}
	}
}

func TestLaneNumbersAreOneToEight(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	discipline := Disciplines[0]
	lanes := assignLanes(r, discipline)

	for i, lane := range lanes {
		expected := i + 1
		if lane.Number != expected {
			t.Errorf("expected lane number %d, got %d", expected, lane.Number)
		}
	}
}

func TestIndividualFinishedWhenDistanceReached(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	discipline := Disciplines[0] // Freestyle 50m
	lanes := assignLanes(r, discipline)
	lane := &lanes[0]
	lane.DistanceCovered = 49.0 // 1m kaldı

	tickIndividual(lane, r, discipline.Style, 4000)

	if !lane.Finished {
		t.Errorf("expected lane to be finished after reaching total distance")
	}
	if lane.FinishTimeMs == 0 {
		t.Errorf("expected finishTimeMs to be set")
	}
}

func TestRelayFinishedWhenAllLegsComplete(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	discipline := Disciplines[12] // 4x100m Freestyle Relay
	lanes := assignLanes(r, discipline)
	lane := &lanes[0]

	// Son bacakta 1m kaldı
	lane.CurrentLegIndex = 3
	lane.CurrentLegDist = 99.0

	tickRelay(lane, r, discipline, 4000)

	if !lane.Finished {
		t.Errorf("expected relay lane to be finished after all legs complete")
	}
}

func TestFinishedLaneRanksBeforeUnfinished(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	race := NewRace(r)

	race.Lanes[0].Finished = true
	race.Lanes[0].FinishTimeMs = 3000
	race.Lanes[1].Finished = false
	race.Lanes[1].DistanceCovered = 99.0

	updatePositions(&race)

	if race.Lanes[0].Position != 1 {
		t.Errorf("finished lane should rank first, got position %d", race.Lanes[0].Position)
	}
}
