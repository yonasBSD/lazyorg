package types

import (
	"fmt"
	"time"
)

type Event struct {
	Name     string
	Time     time.Time
	DurationHour float64
}

func NewEvent(name string, time time.Time, duration float64) *Event {
	return &Event{Name: name, Time: time, DurationHour: duration}
}

func (e *Event) FormatTime() string {
    return fmt.Sprintf("%02dh%02d", e.Time.Hour(), e.Time.Minute())
}
