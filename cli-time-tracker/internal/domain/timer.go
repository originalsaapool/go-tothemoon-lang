package domain

import "time"

type Timer struct {
	ID       string `json:"id"`
	Start    string `json:"start"`
	Duration string `json:"duration"`
	Status   string `json:"status"`
	Name     string `json:"name"`
	// ID       string `json: "id"`
}

type TimerDuration struct {
	Number   int
	String   string
	Duration time.Duration
}
