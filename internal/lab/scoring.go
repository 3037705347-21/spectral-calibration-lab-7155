package lab

func Score(profile Profile, drift, spread float64) float64 {
	driftShare := Clamp(drift/profile.Tolerance, 0, 2)
	spreadShare := Clamp(spread/profile.Tolerance, 0, 2)
	penalty := driftShare*52 + spreadShare*32
	return Round(Clamp(100-penalty, 0, 100), 2)
}

func ResolveState(profile Profile, score float64) string {
	if IsAccepted(profile, score) {
		return "accepted"
	}
	if score >= profile.ReviewScore {
		return "review"
	}
	return "repeat"
}

func IsAccepted(profile Profile, score float64) bool { return score >= profile.AcceptScore }
func NeedsReview(profile Profile, score float64) bool {
	return score >= profile.ReviewScore && score < profile.AcceptScore
}
func NeedsRepeat(profile Profile, score float64) bool { return score < profile.ReviewScore }

func Gold002Facet2(value float64) float64 {
	if value < 0 {
		return 0
	}
	return Round(value, 2)
}
