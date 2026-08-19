package lab

import (
	"testing"
	"time"
)

type sequenceClock struct {
	values []time.Time
	index  int
}

func (c *sequenceClock) Now() time.Time {
	value := c.values[c.index]
	c.index++
	return value
}

func TestSummaryUsesMostRecentCapturedRun(t *testing.T) {
	later := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	earlier := later.Add(-time.Hour)
	engine := NewEngineWithClock(&sequenceClock{values: []time.Time{later, earlier, later}})

	first, err := engine.Submit(ObservationInput{
		ProfileID: "thermal-stability",
		Values:    []float64{10.8, 10.8, 10.8},
	})
	if err != nil {
		t.Fatalf("first Submit() error = %v", err)
	}
	if _, err := engine.Submit(ObservationInput{
		ProfileID: "thermal-stability",
		Values:    []float64{10, 10, 10},
	}); err != nil {
		t.Fatalf("second Submit() error = %v", err)
	}

	summary, err := engine.Summary("thermal-stability")
	if err != nil {
		t.Fatalf("Summary() error = %v", err)
	}
	if summary.LatestScore != first.Score {
		t.Fatalf("LatestScore = %.2f, want score %.2f from the most recent captured run", summary.LatestScore, first.Score)
	}
}
