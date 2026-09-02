package timeutil

import (
	"fmt"
	"time"
)

func FromNow(input string) (time.Time, error) {
	duration, err := time.ParseDuration(input)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to read input: %w", err)
	}

	return time.Now().Add(duration), nil
}
