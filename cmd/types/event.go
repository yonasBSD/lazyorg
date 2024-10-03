package types

import (
	"time"
)

type Event struct {
	Id           int
	Name         string
	Description  string
	Location     string
	Time         time.Time
	DurationHour float64
	FrequencyDay int
	Occurence    int
}

func NewEvent(name, description, location string, time time.Time, duration float64, frequency, occurence int) *Event {
	return &Event{Name: name, Description: description, Location: location, Time: time, DurationHour: duration, FrequencyDay: frequency, Occurence: occurence}
}

func (e Event) GetReccuringEvents() []Event {

	var events []Event
	f := e.FrequencyDay
	initTime := e.Time

	for i := range e.Occurence {
		e.Time = initTime.AddDate(0, 0, i*f)
		events = append(events, e)
	}

	return events
}
