package lab

import "time"

type Profile struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ReferenceCenter float64 `json:"reference_center"`
	Tolerance       float64 `json:"tolerance"`
	MinimumSamples  int     `json:"minimum_samples"`
	AcceptScore     float64 `json:"accept_score"`
	ReviewScore     float64 `json:"review_score"`
	Unit            string  `json:"unit"`
	Description     string  `json:"description"`
}

type ObservationInput struct {
	ProfileID  string    `json:"profile_id"`
	Values     []float64 `json:"values"`
	CapturedBy string    `json:"captured_by"`
}

type ObservationResult struct {
	Status string `json:"status"`
	Run    Run    `json:"run"`
}

type Run struct {
	ID         string    `json:"id"`
	ProfileID  string    `json:"profile_id"`
	Values     []float64 `json:"values"`
	Mean       float64   `json:"mean"`
	Spread     float64   `json:"spread"`
	Drift      float64   `json:"drift"`
	Score      float64   `json:"score"`
	State      string    `json:"state"`
	CapturedAt time.Time `json:"captured_at"`
	CapturedBy string    `json:"captured_by"`
	Notes      []string  `json:"notes"`
}

type QualitySummary struct {
	ProfileID         string    `json:"profile_id"`
	ProfileName       string    `json:"profile_name"`
	TotalRuns         int       `json:"total_runs"`
	LatestScore       float64   `json:"latest_score"`
	AverageScore      float64   `json:"average_score"`
	RecommendedAction string    `json:"recommended_action"`
	ObservedAt        time.Time `json:"observed_at"`
	Signals           []string  `json:"signals"`
}

type Engine struct {
	catalog *Catalog
	ledger  *Ledger
	trace   *TraceabilityIndex
	clock   Clock
}

type Clock interface{ Now() time.Time }
type WallClock struct{}

func (WallClock) Now() time.Time { return time.Now().UTC() }
