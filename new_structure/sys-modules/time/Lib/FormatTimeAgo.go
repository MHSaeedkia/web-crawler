package Lib

import (
	"fmt"
	"time"
)

func FormatTimeAgo(t time.Time) string {
	duration := time.Since(t)
	if duration < time.Minute {
		return "moments ago"
	} else if duration < time.Hour {
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	} else if duration < 24*time.Hour {
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	}
	return fmt.Sprintf("%d days ago", int(duration.Hours()/24))
}
