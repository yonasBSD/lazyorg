package types

import (
	"strconv"
	"time"
)

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

	return month + " " + strconv.Itoa(startDay.Day()) + " to " + strconv.Itoa(endDay.Day())
}

func (c *Calendar) UpdateWeek() {
    c.setWeekLimits()

	for i := range c.CurrentWeek.Days {
        d := c.CurrentWeek.StartDate.AddDate(0, 0, i)
        c.CurrentWeek.Days[i].Date = d
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

func (c *Calendar) GetDayFromTime(time time.Time) Day {
    for _, v := range c.CurrentWeek.Days {
        vYear, vMonth, vDay := v.Date.Date()
        tYear, tMonth, tDay := time.Date()
        if vYear == tYear && vMonth == tMonth && vDay == tDay {
            return v
        }
    }
    return Day{}
}
