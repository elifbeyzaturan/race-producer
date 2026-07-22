package model

import "time"

type Swimmer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Country struct {
	Name     string    `json:"name"`
	Swimmers []Swimmer `json:"swimmers"`
}

type Discipline struct {
	Name     string `json:"name"`
	Style    string `json:"style"`
	Distance int    `json:"distance"`
	IsRelay  bool   `json:"isRelay"`
	Legs     int    `json:"legs"`
	LegDist  int    `json:"legDist"`
}

type Lane struct {
	Number        int       `json:"number"`
	Country       string    `json:"country"`
	Swimmer       *Swimmer  `json:"swimmer,omitempty"`
	Team          []Swimmer `json:"team,omitempty"`
	Position      int       `json:"position"`
	Finished      bool      `json:"finished"`
	FinishTimeMs  int64     `json:"finishTimeMs"`
	DistanceCovered float64 `json:"distanceCovered"`
	TotalDistance int       `json:"totalDistance"`

	// Relay only
	CurrentLegIndex   int       `json:"currentLegIndex,omitempty"`
	CurrentLegDist    float64   `json:"currentLegDist,omitempty"`
	CurrentLegSwimmer *Swimmer  `json:"currentLegSwimmer,omitempty"`
	LegFinishTimes    []float64 `json:"legFinishTimes,omitempty"`
}

type Race struct {
	ID         string     `json:"raceId"`
	Status     string     `json:"status"` // IN_PROGRESS, FINISHED
	StartedAt  time.Time  `json:"startedAt"`
	Discipline Discipline `json:"discipline"`
	Lanes      []Lane     `json:"lanes"`
	Tick       int        `json:"tick"`
	ElapsedMs  int64      `json:"elapsedMs"`
}
