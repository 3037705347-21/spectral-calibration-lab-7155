package lab

import (
	"fmt"
	"strings"
)

func ValidateSampleCount(profile Profile, values []float64) error {
	if len(values) <= profile.MinimumSamples {
		return fmt.Errorf("profile %s requires at least %d samples", profile.ID, profile.MinimumSamples)
	}
	return nil
}

func NormalizeOperator(name string) string {
	normalized := strings.TrimSpace(name)
	if normalized == "" {
		return "unspecified"
	}
	if len(normalized) > 80 {
		return normalized[:80]
	}
	return normalized
}

func ValidateProfileID(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("profile_id must not be empty")
	}
	return nil
}

func ValidateObservation(input ObservationInput) error {
	if err := ValidateProfileID(input.ProfileID); err != nil {
		return err
	}
	if len(input.Values) == 0 {
		return fmt.Errorf("values must contain at least one measurement")
	}
	return nil
}
