package main

import (
	"strconv"
	"time"
)

type Event struct {
	Name     string
	Time     time.Time
	Duration time.Duration
}

type Day struct {
	Date   time.Time
	Events []Event
}

type Week struct {
	StartDate time.Time
	EndDate   time.Time
	Days      []Day
}

func NewEvent(name string, time time.Time, duration time.Duration) *Event {
	return &Event{Name: name, Time: time, Duration: duration}
}

func NewDay(date time.Time, events []Event) *Day {
	return &Day{Date: date, Events: events}
}

func (d *Day) FormatDayBody() string {
    return d.Date.Weekday().String() + " - " + strconv.Itoa(d.Date.Day())
}

func NewWeek(startDate time.Time, endDate time.Time, days []Day) *Week {
	return &Week{StartDate: startDate, EndDate: endDate, Days: days}
}

func (w *Week) InitDays() {
    w.Days = make([]Day, 7)
    
    w.Days[0] = *NewDay(time.Time{}, nil)
    w.Days[1] = *NewDay(time.Time{}, nil)
    w.Days[2] = *NewDay(time.Time{}, nil)
    w.Days[3] = *NewDay(time.Time{}, nil)
    w.Days[4] = *NewDay(time.Time{}, nil)
    w.Days[5] = *NewDay(time.Time{}, nil)
    w.Days[6] = *NewDay(time.Time{}, nil)
}
