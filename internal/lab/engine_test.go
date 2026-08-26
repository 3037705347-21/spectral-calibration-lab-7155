package lab

import (
	"fmt"
	"sync"
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

func TestConcurrentObservationSubmissionsPreserveTraceability(t *testing.T) {
	engine := NewEngine()
	const total = 24
	start := make(chan struct{})
	errs := make(chan error, total)
	var wg sync.WaitGroup
	wg.Add(total)

	for index := 0; index < total/2; index++ {
		index := index
		go func() {
			defer wg.Done()
			<-start
			_, err := engine.Submit(ObservationInput{
				ProfileID:  "thermal-stability",
				Values:     []float64{10, 10.1, 9.9, 10},
				CapturedBy: fmt.Sprintf("operator-%d", index),
			})
			errs <- err
		}()
	}
	for index := total / 2; index < total; index++ {
		index := index
		go func() {
			defer wg.Done()
			<-start
			_, err := engine.Submit(ObservationInput{
				ProfileID:  "thermal-stability",
				Values:     []float64{10, 10.1, 9.9, 10},
				CapturedBy: fmt.Sprintf("operator-%d", index),
			})
			errs <- err
		}()
	}

	close(start)
	wg.Wait()
	for index := 0; index < total; index++ {
		if err := <-errs; err != nil {
			t.Fatalf("Submit() error = %v", err)
		}
	}

	if got := engine.RunCount(); got != total {
		t.Fatalf("RunCount() = %d, want %d", got, total)
	}
	records := engine.trace.ProfileRecords("thermal-stability")
	if len(records) != total {
		t.Fatalf("traceability records = %d, want %d", len(records), total)
	}
	for _, record := range records {
		if !TraceabilityComplete(record) {
			t.Fatalf("incomplete traceability record: %+v", record)
		}
	}
}
