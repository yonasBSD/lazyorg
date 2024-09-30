package types

import (
	"strconv"
	"time"
)

type Day struct {
	Date   time.Time
	Events []Event
}

func NewDay(date time.Time) *Day {
	return &Day{Date: date}
}

func (d *Day) FormatDayBody() string {
	return d.Date.Weekday().String() + "-" + strconv.Itoa(d.Date.Day())
}

