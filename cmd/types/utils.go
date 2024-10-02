package types

import (
	"fmt"
	"strings"
	"time"
)

func DurationToHeight(d float64) int {
	return int(d * 2)
}

func FormatHour(t time.Time) string {
	return fmt.Sprintf("%02d:%02d", t.Hour(), t.Minute())
}

func TimeToPosition(t time.Time, s string) int {

    time := FormatHour(t)
	lines := strings.Split(s, "\n")

	for i, v := range lines {
		if strings.Contains(v, time) {
			return i
		}
	}

	return -1
}

