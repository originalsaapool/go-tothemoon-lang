package domain

import "time"

type Timer struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Duration  time.Duration `json:"duration"`
	StartedAt time.Time     `json:"start"`
	StoppedAt time.Time     `json:"end"`
}

type TimerDuration struct {
	Number   int
	String   string
	Duration time.Duration
}
