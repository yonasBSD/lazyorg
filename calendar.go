package main

import (
	"strconv"
	"time"
)

type Calendar struct {
	startWeek time.Time
	endWeek   time.Time
	daysName  []string
}

func (c *Calendar) initWeek() {
	currentDay := time.Now()
	currentWeekDay := currentDay.Weekday()

	diffToSunday := currentWeekDay
	diffToSaturday := 6 - currentWeekDay

	c.startWeek = currentDay.AddDate(0, 0, -int(diffToSunday))
	c.endWeek = currentDay.AddDate(0, 0, int(diffToSaturday))

	c.daysName = make([]string, 7)

	for i := range c.daysName {
        d := c.startWeek.AddDate(0, 0, i)
		c.daysName[i] = d.Weekday().String() + "\n" + strconv.Itoa(d.Day())
	}
}
