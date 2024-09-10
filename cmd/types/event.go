package types

import "time"

type Event struct {
	Name     string
	Time     time.Time
	Duration time.Duration
}

func NewEvent(name string, time time.Time, duration time.Duration) *Event {
	return &Event{Name: name, Time: time, Duration: duration}
}
