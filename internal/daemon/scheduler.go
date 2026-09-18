package daemon

import "time"

// 'newTimer' receives a due date and returns a time.Timer that matches it
func NewTimer(due time.Time) *time.Timer {
	duration := time.Until(due)

	duration = max(0, duration)

	return time.NewTimer(duration)
}
