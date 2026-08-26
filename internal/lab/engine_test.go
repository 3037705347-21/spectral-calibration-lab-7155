package lab

import (
	"testing"
	"time"
)

type fixedClock struct{ value time.Time }

func (c fixedClock) Now() time.Time { return c.value }

func TestSubmitAcceptsStableThermalObservation(t *testing.T) {
	engine := NewEngineWithClock(fixedClock{value: time.Date(2026, 8, 19, 0, 0, 0, 0, time.UTC)})
	run, err := engine.Submit(ObservationInput{ProfileID: "thermal-stability", Values: []float64{10, 10.1, 9.9, 10}, CapturedBy: "tester"})
	if err != nil {
		t.Fatalf("Submit() error = %v", err)
	}
	if run.State != "accepted" {
		t.Fatalf("state = %s", run.State)
	}
}

func TestSubmitRejectsTooFewSamples(t *testing.T) {
	engine := NewEngine()
	_, err := engine.Submit(ObservationInput{ProfileID: "thermal-stability", Values: []float64{10, 10.1}})
	if err == nil {
		t.Fatal("expected sample validation error")
	}
}

func TestSummaryRecommendsInitialObservation(t *testing.T) {
	engine := NewEngine()
	summary, err := engine.Summary("thermal-stability")
	if err != nil {
		t.Fatalf("Summary() error = %v", err)
	}
	if summary.TotalRuns != 0 {
		t.Fatalf("TotalRuns = %d", summary.TotalRuns)
	}
}

func TestHistoryReturnsIndependentRunSnapshots(t *testing.T) {
	engine := NewEngine()
	if _, err := engine.Submit(ObservationInput{
		ProfileID:  "thermal-stability",
		Values:     []float64{10, 10.1, 9.9, 10},
		CapturedBy: "operator",
	}); err != nil {
		t.Fatalf("Submit() error = %v", err)
	}

	history, err := engine.History("thermal-stability")
	if err != nil {
		t.Fatalf("History() error = %v", err)
	}
	history[0].Values[0] = 999
	history[0].Notes[0] = "changed outside the engine"

	again, err := engine.History("thermal-stability")
	if err != nil {
		t.Fatalf("History() second call error = %v", err)
	}
	if again[0].Values[0] == 999 {
		t.Fatal("mutating returned values changed the stored run")
	}
	if again[0].Notes[0] == "changed outside the engine" {
		t.Fatal("mutating returned notes changed the stored run")
	}
}
