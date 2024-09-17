package main

import (
	"bytes"
	"strings"

	"github.com/HubertBel/go-organizer/cmd/database"
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/HubertBel/go-organizer/views"

	"time"

	"github.com/jroimartin/gocui"
)

type CalendarController struct {
	model    *types.Calendar
	view     *views.WeekView
	Database *database.Database
}

func NewCalendarController() *CalendarController {
	today := types.NewDay(time.Now(), nil)

	c := types.NewCalendar(today)
	wv := views.NewWeekView("week", 0, 0, 0, 0)

	return &CalendarController{model: c, view: wv}
}

func (cc *CalendarController) InitDatabase(path string) error {
	cc.Database = &database.Database{}
	return cc.Database.InitDatabase(path)
}

func (cc *CalendarController) Layout(g *gocui.Gui) error {
	err := cc.updateEvents()
	if err != nil {
		return err
	}

	cc.updateViewPropreties(g)

	cc.setWeekViewBody()
	cc.setDayViewsBody()

	cc.updateEventViews()

	return cc.view.Layout(g)
}

func (cc *CalendarController) UpdateToNextWeek() error {
	cc.model.UpdateToNextWeek()
	return nil // TODO error handling
}

func (cc *CalendarController) UpdateToPrevWeek() error {
	cc.model.UpdateToPrevWeek()
	return nil // TODO error handling
}

func (cc *CalendarController) updateEvents() error {
	for i := range cc.model.CurrentWeek.Days {
		day := &cc.model.CurrentWeek.Days[i]
		var err error
		day.Events, err = cc.Database.GetEventsByDate(day.Date)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cc *CalendarController) updateEventViews() error {
    buffer := cc.view.Body

	for i, day := range cc.model.CurrentWeek.Days {
		var eventViews []views.EventView
		dayView := &cc.view.DayViews[i]
		for _, event := range day.Events {
			x := dayView.X + 1
			y := timeToPosisition(buffer, event.FormatTime())
			w := dayView.W - 2
			h := durationToPosition(event.DurationHour)
			eventViews = append(eventViews, *views.NewEvenView(event.Name, x, y, w, h, event.Name))
		}
		dayView.EventViews = eventViews
	}

    return nil
}

func (cc *CalendarController) AddTestEvents() error {
    // t := time.Date(2024, time.September, 8, 12, 0, 0, 0, time.Now().Location())
	// e := types.NewEvent("Archi", t, 2.0)

	// var err error
	// _, err = cc.Database.AddEvent(e)
	// if err != nil {
	// 	return err
	// }

    // t = time.Date(2024, time.September, 12, 12, 0, 0, 0, time.Now().Location())
	// e = types.NewEvent("Tennis", t, 1.5)
	// _, err = cc.Database.AddEvent(e)
	// if err != nil {
	// 	return err
	// }

    t := time.Date(2024, time.September, 9, 15, 30, 0, 0, time.Now().Location())
    e := types.NewEvent("Test", t, 1)
    _, err := cc.Database.AddEvent(e)
	if err != nil {
		return err
	}

	return nil
}

func (cc *CalendarController) updateViewPropreties(g *gocui.Gui) {
	maxX, maxY := g.Size()

	x0 := 0
	y0 := 0
	w := (maxX - 1) - x0
	h := (maxY - 1) - y0

	cc.view.Update(x0, y0, w, h)
}

func (cc *CalendarController) setDayViewsBody() {
	days := cc.model.CurrentWeek.Days
	for i, v := range days {
		cc.view.DayViews[i].Body = v.FormatDayBody()
	}
}

func (cc *CalendarController) setWeekViewBody() {
    var b bytes.Buffer

	h := cc.view.H - cc.view.Y - 5

	b.WriteString(cc.model.FormatWeekBody())
    b.WriteString("\n")
    b.WriteString("\n")
	cc.view.WriteTimes(&b, h)

	cc.view.Body = b.String()
}

func durationToPosition(d float64) int {
	return int(d * 2)
}

func timeToPosisition(buffer string, t string) int {
	lines := strings.Split(buffer, "\n")

	for i, v := range lines {
		if strings.Contains(v, t) {
			return i + 1
		}
	}
	return 0
}
