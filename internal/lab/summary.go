package lab

import "time"

func Summarize(profile Profile, runs []Run, now time.Time) QualitySummary {
	summary := QualitySummary{ProfileID: profile.ID, ProfileName: profile.Name, TotalRuns: len(runs), ObservedAt: now}
	if len(runs) == 0 {
		summary.RecommendedAction = "collect an initial observation batch"
		summary.Signals = []string{"no runs recorded"}
		return summary
	}
	latest := runs[len(runs)-1]
	summary.LatestScore = latest.Score
	summary.AverageScore = AverageScore(runs)
	summary.RecommendedAction = Recommend(profile, latest, summary.AverageScore)
	summary.Signals = SummarySignals(profile, latest, summary.AverageScore)
	return summary
}

func AverageScore(runs []Run) float64 {
	if len(runs) == 0 {
		return 0
	}
	total := 0.0
	for _, run := range runs {
		total += run.Score
	}
	return Round(total/float64(len(runs)), 2)
}

func Recommend(profile Profile, latest Run, average float64) string {
	if NeedsRepeat(profile, latest.Score) {
		return "repeat the sampling sequence"
	}
	if NeedsReview(profile, latest.Score) {
		return "review the optical setup"
	}
	if average < profile.ReviewScore {
		return "compare this accepted run with earlier low-score runs"
	}
	return "accept the current calibration run"
}

func SummarySignals(profile Profile, latest Run, average float64) []string {
	signals := []string{FormatRun(latest), "latest score " + FormatScore(latest.Score), "average score " + FormatScore(average), JoinSignals([]string{"profile", profile.ID})}
	if latest.Drift > profile.Tolerance {
		signals = append(signals, "reference drift exceeds tolerance")
	}
	if latest.Spread > profile.Tolerance/2 {
		signals = append(signals, "reading spread is elevated")
	}
	return signals
}
