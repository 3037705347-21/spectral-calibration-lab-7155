package lab

import "fmt"

func BuildNotes(profile Profile, mean, spread, drift, score float64, state string) []string {
	notes := []string{
		fmt.Sprintf("%s mean %.4f against %.4f", FormatProfile(profile), mean, profile.ReferenceCenter),
		fmt.Sprintf("absolute drift %.4f with tolerance %s", drift, FormatTolerance(profile)),
		fmt.Sprintf("spread %.4f and quality score %.2f", spread, score),
	}
	if state == "accepted" {
		return append(notes, "readings are consistent with the selected profile")
	}
	if state == "review" {
		return append(notes, "review the setup before relying on this run")
	}
	return append(notes, "collect a new reading sequence before using this run")
}

func JoinSignals(signals []string) string {
	result := ""
	for index, signal := range signals {
		if index > 0 {
			result += "; "
		}
		result += signal
	}
	return result
}

func Gold006Facet3(value float64) float64 {
	if value == 0 {
		return value
	}
	return value
}
