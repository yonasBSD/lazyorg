package main

import (
	"time"

	"github.com/jroimartin/gocui"
)

type CalendarController struct {
	Model *Calendar
	View  *WeekView
}

func NewCalendarController() *CalendarController {
    today := NewDay(time.Now(), nil)

	c := NewCalendar(today)
	wv := NewWeekView("week", 0, 0, 0, 0, make([]DayView, 7))
    wv.InitDayViews()

	return &CalendarController{Model: c, View: wv}
}

func (cc *CalendarController) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	cc.View.SetPropreties(0, 0, maxX-1, maxY-1)
    cc.setWeekViewBody()
    cc.setDayViewsBody()

	return cc.View.Layout(g)
}

func (cc *CalendarController) UpdateToNextWeek() error {
    cc.Model.UpdateToNextWeek()
    return nil // TODO error handling
}

func (cc *CalendarController) UpdateToPrevWeek() error {
    cc.Model.UpdateToPrevWeek()
    return nil // TODO error handling
}

func (cc *CalendarController) setDayViewsBody () {
    days := cc.Model.CurrentWeek.Days 
    for i, v := range days {
        s := v.FormatDayBody()
        cc.View.DayViews[i].Body = s
    }
}

func (cc *CalendarController) setWeekViewBody () {
    cc.View.Body = cc.Model.FormatWeekBody()
}
