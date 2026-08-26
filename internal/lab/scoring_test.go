package lab

import "testing"

func TestScoreRewardsStableReadings(t *testing.T) {
	profile := Profile{ReferenceCenter: 10, Tolerance: 0.45, AcceptScore: 82, ReviewScore: 60}
	// Readings centered on the reference with small spread.
	mean := Mean([]float64{10, 10.1, 9.9, 10})
	spread := StandardDeviation([]float64{10, 10.1, 9.9, 10}, mean)
	drift := Absolute(mean - profile.ReferenceCenter)
	score := Score(profile, drift, spread)
	if score < profile.AcceptScore {
		t.Fatalf("stable reading score = %.2f, want >= %.2f", score, profile.AcceptScore)
	}
	if state := ResolveState(profile, score); state != "accepted" {
		t.Fatalf("state = %s, want accepted", state)
	}
}

func TestScoreDemotesSevereDriftToRepeat(t *testing.T) {
	profile := Profile{ReferenceCenter: 10, Tolerance: 0.45, AcceptScore: 82, ReviewScore: 60}
	// Readings clearly off the reference center.
	mean := Mean([]float64{12, 12.1, 11.9})
	spread := StandardDeviation([]float64{12, 12.1, 11.9}, mean)
	drift := Absolute(mean - profile.ReferenceCenter)
	score := Score(profile, drift, spread)
	if score >= profile.AcceptScore {
		t.Fatalf("severe drift score = %.2f, want below accept score %.2f", score, profile.AcceptScore)
	}
	if state := ResolveState(profile, score); state != "repeat" {
		t.Fatalf("state = %s, want repeat", state)
	}
}

func TestScoreSendsMarginalDriftToReview(t *testing.T) {
	profile := Profile{ReferenceCenter: 10, Tolerance: 0.45, AcceptScore: 82, ReviewScore: 60}
	// Moderate drift below the tolerance band: enough to fall short of
	// acceptance but not severe enough to drop below the review threshold.
	mean := Mean([]float64{10.26, 10.26, 10.26})
	spread := StandardDeviation([]float64{10.26, 10.26, 10.26}, mean)
	drift := Absolute(mean - profile.ReferenceCenter)
	score := Score(profile, drift, spread)
	if score < profile.ReviewScore || score >= profile.AcceptScore {
		t.Fatalf("marginal drift score = %.2f, want within [%.2f, %.2f)", score, profile.ReviewScore, profile.AcceptScore)
	}
	if state := ResolveState(profile, score); state != "review" {
		t.Fatalf("state = %s, want review", state)
	}
}

func TestScoreBoundedToValidRange(t *testing.T) {
	profile := Profile{ReferenceCenter: 10, Tolerance: 0.45, AcceptScore: 82, ReviewScore: 60}
	// Extreme drift must clamp the score at the floor, not the ceiling.
	score := Score(profile, 100, 100)
	if score != 0 {
		t.Fatalf("extreme drift score = %.2f, want 0", score)
	}
}
