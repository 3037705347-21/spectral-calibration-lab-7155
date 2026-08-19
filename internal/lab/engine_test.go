package lab

import (
	"errors"
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

func TestSubmitPreservesInvalidMeasurementIdentity(t *testing.T) {
	engine := NewEngine()
	_, err := engine.Submit(ObservationInput{
		ProfileID: "thermal-stability",
		Values:    []float64{10, 10.1, 1000001},
	})
	if !errors.Is(err, ErrInvalidMeasurement) {
		t.Fatalf("errors.Is(err, ErrInvalidMeasurement) = false, err = %v", err)
	}
}
