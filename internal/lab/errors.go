package lab

import (
	"errors"
	"fmt"
)

var ErrInvalidMeasurement = errors.New("invalid measurement")

type InvalidMeasurementError struct {
	Index  int
	Value  float64
	Reason string
}

func (e *InvalidMeasurementError) Error() string {
	return fmt.Sprintf("measurement %d (%.4f): %s", e.Index, e.Value, e.Reason)
}

func (e *InvalidMeasurementError) Unwrap() error {
	return ErrInvalidMeasurement
}

func IsInvalidMeasurement(err error) bool {
	return errors.Is(err, ErrInvalidMeasurement)
}
