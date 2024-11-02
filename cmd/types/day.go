package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Day struct {
	Date   time.Time
	Events []*Event
}

func NewDay(date time.Time) *Day {
	return &Day{Date: date}
}

func (d *Day) FormatTitle() string {
	return d.Date.Weekday().String() + "-" + strconv.Itoa(d.Date.Day())
}

func (d *Day) FormatTimeAndHour() string {
	s := fmt.Sprintf("%s %d | %02dh%02d", d.Date.Month().String(), d.Date.Day(), d.Date.Hour(), d.Date.Minute())
	return s
}

func (d *Day) FormatBody() string {
	var sb strings.Builder

    sb.WriteString("\n")
    sb.WriteString("\nEvents :\n")
    sb.WriteString( "---------\n")
	for _, v := range d.Events {
		s := "-> " + v.FormatTimeAndName() + "\n"
		sb.WriteString(s)
	}

	return sb.String()
}

func (d *Day) SortEventsByTime() {
    sort.Slice(d.Events, func(i, j int) bool {
        return d.Events[i].Time.Before(d.Events[j].Time)
    })
}
