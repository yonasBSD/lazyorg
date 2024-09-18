package views

import (
	"fmt"
	"math"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

const (
	border = 1
)

type WeekView struct {
	Name string
	X, Y int
	W, H int
	Body string

	TimeView   TimeView
	DayViews   []DayView
	EventViews []EventView

	Database *types.Database

	Calendar *types.Calendar
}

func NewWeekView(db *types.Database) *WeekView {
	wv := &WeekView{Name: "week", Database: db}
	wv.TimeView = *NewTimeView("time", 0, 0, 0, 0, "")
	wv.EventViews = make([]EventView, 0)
	wv.Calendar = types.NewCalendar(types.NewDay(time.Now(), nil))
	wv.initDayViews()
	return wv
}

func (wv *WeekView) SetProperties(x, y, w, h int) {
	wv.X = x
	wv.Y = y
	wv.W = w
	wv.H = h
}

func (wv *WeekView) Layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	wv.SetProperties(0, 0, maxX-1, maxY-1)
	wv.Body = wv.Calendar.FormatWeekBody()

	v, err := g.SetView(wv.Name, wv.X, wv.Y, wv.X+wv.W, wv.Y+wv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Wrap = true
		v.Autoscroll = true
	}

	v.Clear()
	fmt.Fprintln(v, wv.Body)

	err = wv.updateTimeView(g)
	if err != nil {
		return err
	}

	err = wv.updateDayViews(g)
	if err != nil {
		return err
	}

	err = wv.updateEvents(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) AddTestEvents() error {
	t := time.Date(2024, time.September, 16, 10, 30, 0, 0, time.Now().Location())
	e := types.NewEvent("Archi1", t, 2.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 16, 13, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Astro1", t, 1.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 17, 16, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Tennis", t, 2.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 18, 13, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Astro2", t, 2.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 18, 15, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Russe", t, 3.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 19, 9, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Robotique", t, 3.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 19, 13, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Archi2", t, 2.0)

	wv.Database.AddEvent(e)

	t = time.Date(2024, time.September, 20, 10, 30, 0, 0, time.Now().Location())
	e = types.NewEvent("Robotique lab", t, 2.0)

	wv.Database.AddEvent(e)

	return nil
}

func (wv *WeekView) UpdateToNextWeek() {
	wv.Calendar.UpdateToNextWeek()
}

func (wv *WeekView) UpdateToPrevWeek() {
	wv.Calendar.UpdateToPrevWeek()
}

func (wv *WeekView) initDayViews() {
	wv.DayViews = make([]DayView, 7)
	wv.DayViews[0] = *NewDayView("sunday", 0, 0, 0, 0, "")
	wv.DayViews[1] = *NewDayView("monday", 0, 0, 0, 0, "")
	wv.DayViews[2] = *NewDayView("tuesday", 0, 0, 0, 0, "")
	wv.DayViews[3] = *NewDayView("wednesday", 0, 0, 0, 0, "")
	wv.DayViews[4] = *NewDayView("thursday", 0, 0, 0, 0, "")
	wv.DayViews[5] = *NewDayView("friday", 0, 0, 0, 0, "")
	wv.DayViews[6] = *NewDayView("saturday", 0, 0, 0, 0, "")
}

func (wv *WeekView) updateEvents(g *gocui.Gui) error {
	err := wv.Calendar.UpdateEventsFromDatabase(wv.Database)
	if err != nil {
		return err
	}

	wv.clearEventViews(g)
	wv.createEventViews()

	err = wv.updateEventViews(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) updateTimeView(g *gocui.Gui) error {
	x := wv.X
	y := wv.Y + 3
	w := 8
	h := wv.H - wv.Y - 4

	wv.TimeView.SetProperties(x, y, w, h)
	err := wv.TimeView.Update(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) updateDayViews(g *gocui.Gui) error {
	x := wv.X + wv.TimeView.W + 1
	y := wv.Y + 3
	w := calculateDayViewWidth(x, wv.W)
	h := wv.H - wv.Y - 4

	for i := range wv.DayViews {
		wv.DayViews[i].SetProperties(x, y, w, h)
		wv.DayViews[i].Body = wv.Calendar.CurrentWeek.Days[i].FormatDayBody()

		err := wv.DayViews[i].Update(g)
		if err != nil {
			return err
		}

		x += w + border
	}
	return nil
}

func (wv *WeekView) clearEventViews(g *gocui.Gui) error {

	for _, v := range wv.EventViews {
		err := g.DeleteView(v.Name)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
	}
	wv.EventViews = nil

	return nil
}

func (wv *WeekView) createEventViews() {
	var eventViews []EventView

	for i, day := range wv.Calendar.CurrentWeek.Days {
		for _, event := range day.Events {
			x := wv.DayViews[i].X + 1
			y := wv.TimeView.TimeToPosition(event.FormatHour())
			w := wv.DayViews[i].W - 2
			h := wv.TimeView.DurationToHeight(event.DurationHour)

			if y > wv.DayViews[i].Y && y+h < wv.DayViews[i].Y+wv.DayViews[i].H {
				eventViews = append(eventViews, *NewEvenView(event.Name, x, y, w, h, event.Name))
			}
		}
	}

	wv.EventViews = eventViews
}

func (wv *WeekView) updateEventViews(g *gocui.Gui) error {

	for i := range wv.EventViews {
		err := wv.EventViews[i].Update(g)
		if err != nil {
			return err
		}
	}

	return nil
}

func calculateDayViewWidth(x, w int) int {
	width := math.Round(float64((w-x)/7)) - border
	return int(width)
}
