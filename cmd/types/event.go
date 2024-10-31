package types

import (
	"fmt"
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

func (e *Event) FormatTimeAndName() string {
	startTimeString := fmt.Sprintf("%02dh%02d", e.Time.Hour(), e.Time.Minute())

    duration := time.Duration(e.DurationHour * float64(time.Hour))
    endTime := e.Time.Add(duration)
	endTimeString := fmt.Sprintf("%02dh%02d", endTime.Hour(), endTime.Minute())

	s := fmt.Sprintf("%s-%s | %s", startTimeString, endTimeString, e.Name)

	return s
}

func (e *Event) FormatTitle() string {
	return ""
}

func (e *Event) FormatBody() string {
	return ""
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
