package main

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/HubertBel/go-organizer/views"

	"time"

	"github.com/jroimartin/gocui"
)

type CalendarController struct {
	Model *types.Calendar
	View  *views.WeekView
}

func NewCalendarController() *CalendarController {
    today := types.NewDay(time.Now(), nil)

	c := types.NewCalendar(today)
    wv := views.NewWeekView("week", 0, 0, 0, 0, make([]views.DayView, 7))

	return &CalendarController{Model: c, View: wv}
}

func (cc *CalendarController) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

    x0 := 0
    y0 := 1

    x1 := (maxX-1)-x0
    y1 := (maxY-1)-y0

	cc.View.SetPropreties(x0, y0, x1, y1)
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
        cc.View.DayViews[i].Body = v.FormatDayBody()
    }
}

func (cc *CalendarController) setWeekViewBody () {
    cc.View.Body = cc.Model.FormatWeekBody()
}
