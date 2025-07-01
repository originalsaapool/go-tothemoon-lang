package tracker

import (
	"testing"
)

func TestAddTimer(t *testing.T) {
	tr := NewTracker()

	err := tr.AddTimer("task", "5")
	if err != nil {
		t.Fatalf("AddTimer returned error: %v", err)
	}

	_, ok := tr.Timers["task"]
	if !ok {
		t.Fatalf("Timer not found after AddTimer")
	}

}

func TestStopTimer(t *testing.T) {

	tr := NewTracker()
	tr.AddTimer("task", "5")

	err := tr.Stop("task")
	if err != nil {
		t.Fatalf("Stop returned error: %v", err)
	}
	err = tr.Stop("task")
	if err == nil {
		t.Errorf("Expected error on second stop, got nil")
	}
}
