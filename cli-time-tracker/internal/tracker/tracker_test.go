package tracker

import (
	"cli-time-tracker/internal/domain"
	"strconv"
	"testing"
)

func TestAddTimer(t *testing.T) {

	tr := &Tracker{
		Timers: make(map[string]*domain.Timer),
	}
	name := "to do something"
	minutes := "5"
	tr.AddTimer(name, minutes)

	_, ok := tr.Timers[name]
	if !ok {
		t.Errorf("Can't find added timer in tracker")
	}

	if len(tr.Timers) == 0 {
		t.Errorf("No new timers added")
	}
	minutes_str, err := strconv.Atoi(minutes)

	if err != nil {
		t.Errorf("Used letters instead of numbers to set the timer")
	}
	seconds := minutes_str * 60
	seconds_str := strconv.Itoa(seconds)
	if tr.Timers[name].Duration != seconds_str {
		t.Errorf("Duration from input and in timer are different")
	}
}

func TestStopTimer(t *testing.T) {

	tr := &Tracker{
		Timers: make(map[string]*domain.Timer),
	}
	name := "to do something"
	minutes := "5"
	tr.AddTimer(name, minutes)

	tr.Stop(name)

	if tr.Timers[name].Status != "stopped" {
		t.Errorf("Timer hasn't been stopped")
	}
}
