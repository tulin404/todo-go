package timeutil

import (
	"fmt"
	"time"
)

// 'FromNow' receives a input (string) and returns the sum of this input and the current time
func FromNow(input string) (*time.Time, error) {
	duration, err := time.ParseDuration(input)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	t := time.Now().Add(duration)
	return &t, nil
}
