package tracker

import (
	"strconv"
	"testing"
)

func TestAddTimer(t *testing.T) {
	tr := NewTracker()

	err := tr.AddTimer("task", "5")
	if err != nil {
		t.Fatalf("AddTimer returned error: %v", err)
	}

	timer, ok := tr.Timers["task"]
	if !ok {
		t.Fatalf("Timer not found after AddTimer")
	}

	minutes_str, err := strconv.Atoi("5")

	if err != nil {
		t.Errorf("Used letters instead of numbers to set the timer")
	}
	seconds := minutes_str * 60
	seconds_str := strconv.Itoa(seconds)

	expected := seconds_str
	if timer.Duration != expected {
		t.Errorf("Expected duration %v, got %v", expected, timer.Duration)
	}
}

func TestStopTimer(t *testing.T) {

	tr := NewTracker()
	tr.AddTimer("task", "5")

	err := tr.Stop("task")
	if err != nil {
		t.Fatalf("Stop returned error: %v", err)
	}

	timer := tr.Timers["task"]
	if timer.Status != "stopped" {
		t.Errorf("Timer hasn't been stopped")
	}
}
