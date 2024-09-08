package main

import (
	"strconv"
	"time"
)

type Calendar struct {
    currentDay time.Time
	startDay time.Time
	endDay   time.Time
	daysName  []string
}

func NewCalendar() *Calendar {
    c := &Calendar{
        daysName: make([]string, 7),
    }
    c.initWeek()
    return c
}

func (c *Calendar) initWeek() {
	c.currentDay = time.Now()
    c.updateWeek()
}

func (c *Calendar) updateWeek() {
    currentWeekDay := c.currentDay.Weekday()

	diffToSunday := currentWeekDay
	diffToSaturday := 6 - currentWeekDay

	c.startDay = c.currentDay.AddDate(0, 0, -int(diffToSunday))
	c.endDay = c.currentDay.AddDate(0, 0, int(diffToSaturday))

	for i := range c.daysName {
        d := c.startDay.AddDate(0, 0, i)
        c.daysName[i] = d.Weekday().String() + " - " + strconv.Itoa(d.Day())
    }
}

func (c *Calendar) nextWeek() {
    c.currentDay = c.currentDay.AddDate(0, 0, 7)
    c.updateWeek()
}

func (c *Calendar) prevWeek() {
    c.currentDay = c.currentDay.AddDate(0, 0, -7)
    c.updateWeek()
}
