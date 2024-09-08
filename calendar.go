package main

import (
	"strconv"
	"time"
)

type Calendar struct {
	startDay time.Time
	endDay   time.Time
	daysName  []string
}

func (c *Calendar) initWeek() {
	currentDay := time.Now()
    currentWeekDay := currentDay.Weekday()

	diffToSunday := currentWeekDay
	diffToSaturday := 6 - currentWeekDay

	c.startDay = currentDay.AddDate(0, 0, -int(diffToSunday))
	c.endDay = currentDay.AddDate(0, 0, int(diffToSaturday))

	c.daysName = make([]string, 7)

	for i := range c.daysName {
        d := c.startDay.AddDate(0, 0, i)
		c.daysName[i] = d.Weekday().String() + " - " + strconv.Itoa(d.Day())
	}
}
