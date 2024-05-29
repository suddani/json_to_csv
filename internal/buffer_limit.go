package json_to_csv

import (
	"math"
)

type BufferLimit struct {
	Default int
	Max     int
	Valid   bool
}

func NewBufferLimit(size int, maxSize int) BufferLimit {
	defaultValue := 0
	maxValue := 0
	if size == 0 && maxSize == 0 {
		return BufferLimit{
			Default: defaultValue,
			Max:     maxValue,
			Valid:   false,
		}
	}

	if size == 0 {
		defaultValue = int(math.Min(float64(64*1024), float64(maxSize)))
	} else {
		defaultValue = size
	}

	if maxSize == 0 {
		maxValue = int(math.Max(float64(defaultValue), float64(64*1024)))
	} else {
		maxValue = int(math.Max(float64(defaultValue), float64(maxSize)))
	}

	return BufferLimit{
		Default: defaultValue,
		Max:     maxValue,
		Valid:   true,
	}
}
