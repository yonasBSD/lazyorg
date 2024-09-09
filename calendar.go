package main

import "strconv"

type Calendar struct {
    CurrentDay *Day
    CurrentWeek *Week
}

func NewCalendar(currentDay *Day) *Calendar {
    c := &Calendar{CurrentDay: currentDay, CurrentWeek: &Week{}}
    c.CurrentWeek.InitDays()
    c.UpdateWeek()

    return c
}

func (c *Calendar) setWeekLimits () {
    d := c.CurrentDay.Date

	diffToSunday := d.Weekday()
	diffToSaturday := 6 - d.Weekday()

    c.CurrentWeek.StartDate = d.AddDate(0, 0, -int(diffToSunday))
    c.CurrentWeek.EndDate = d.AddDate(0, 0, int(diffToSaturday))
}

func (c *Calendar) FormatWeekBody () string {
	startDay := c.CurrentWeek.StartDate
	endDay := c.CurrentWeek.EndDate
	month := endDay.Month().String()

	return "Week: " + strconv.Itoa(startDay.Day()) + " to " + strconv.Itoa(endDay.Day()) + ", " + month
}

func (c *Calendar) UpdateWeek() {
    c.setWeekLimits()

	for i := range c.CurrentWeek.Days {
        d := c.CurrentWeek.StartDate.AddDate(0, 0, i)
        c.CurrentWeek.Days[i].Date = d
        // Add event logic
    }
}

func (c *Calendar) UpdateToNextWeek() {
    c.CurrentDay.Date = c.CurrentDay.Date.AddDate(0, 0, 7)
    c.UpdateWeek()
}

func (c *Calendar) UpdateToPrevWeek() {
    c.CurrentDay.Date = c.CurrentDay.Date.AddDate(0, 0, -7)
    c.UpdateWeek()
}
