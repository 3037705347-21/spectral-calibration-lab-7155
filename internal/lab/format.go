package lab

import (
	"fmt"
	"strings"
)

func FormatScore(score float64) string     { return fmt.Sprintf("%.2f", score) }
func FormatProfile(profile Profile) string { return profile.ID + " (" + profile.Unit + ")" }
func FormatState(state string) string      { return strings.ToUpper(strings.TrimSpace(state)) }
func FormatRun(run Run) string {
	return fmt.Sprintf("%s:%s:%s", run.ID, run.ProfileID, FormatState(run.State))
}
func FormatTolerance(profile Profile) string {
	return fmt.Sprintf("±%.4f %s", profile.Tolerance, profile.Unit)
}
