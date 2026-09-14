package daemon

import "time"

func NewTimer(due time.Time) *time.Timer {
    duration := time.Until(due)

    duration = max(0, duration)

    return time.NewTimer(duration)
}
