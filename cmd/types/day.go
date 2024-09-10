package types

import (
	"strconv"
	"time"
)

type Day struct {
	Date   time.Time
	Events []Event
}

func NewDay(date time.Time, events []Event) *Day {
	return &Day{Date: date, Events: events}
}

func (d *Day) FormatDayBody() string {
	return d.Date.Weekday().String() + "-" + strconv.Itoa(d.Date.Day())
}
