package calendar

import (
	"fmt"
	"strings"
	"time"

	"github.com/HubertBel/lazyorg/internal/utils"
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
	return fmt.Sprintf("%s | %s", e.FormatDurationTime(), e.Name)
}

func (e *Event) FormatDurationTime() string {
	startTimeString := utils.FormatHourFromTime(e.Time)

    duration := time.Duration(e.DurationHour * float64(time.Hour))
    endTime := e.Time.Add(duration)
	endTimeString := utils.FormatHourFromTime(endTime)

    return fmt.Sprintf("%s-%s", startTimeString, endTimeString)
}

func (e *Event) FormatBody() string {
	var sb strings.Builder

    sb.WriteString("\n")
    sb.WriteString(fmt.Sprintf("\n%s | %s\n", e.FormatDurationTime(), e.Location))
    sb.WriteString("\nDescription :\n")
    sb.WriteString( "--------------\n")
    sb.WriteString(e.Description)

	return sb.String()
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
