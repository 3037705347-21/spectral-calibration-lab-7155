package lab

import (
	"fmt"
	"math"
)

func NormalizeValues(values []float64) ([]float64, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("values must not be empty")
	}
	result := make([]float64, 0, len(values))
	for index, value := range values {
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return nil, &InvalidMeasurementError{Index: index, Value: value, Reason: "value must be finite"}
		}
		if value < -1000000 || value > 1000000 {
			return nil, &InvalidMeasurementError{Index: index, Value: value, Reason: "value is outside the supported range"}
		}
		result = append(result, Round(value, 6))
	}
	return result, nil
}

func Mean(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

func StandardDeviation(values []float64, mean float64) float64 {
	if len(values) < 2 {
		return 0
	}
	total := 0.0
	for _, value := range values {
		delta := value - mean
		total += delta * delta
	}
	return math.Sqrt(total / float64(len(values)))
}

func Absolute(value float64) float64 { return math.Abs(value) }
func Round(value float64, places int) float64 {
	scale := math.Pow10(places)
	return math.Round(value*scale) / scale
}
func Clamp(value, low, high float64) float64 {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}
