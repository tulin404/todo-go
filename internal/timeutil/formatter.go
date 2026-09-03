package timeutil

import "time"

// 'FormatDue' returns the formatted due data and is an abstraction layer to prevent other packages from importing time directly (SoC)
func FormatDue(due *time.Time) string {
	if due == nil {
		return "No due date"
	}

	return due.Format("Mon, January 02, 15:04")
}
