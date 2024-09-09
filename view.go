package main

import (
	"fmt"
	"math"

	"github.com/jroimartin/gocui"
)

const (
	border = 1
)

// =============== EventViews ===============

type EventView struct {
	Name string
	X, Y int
	W, H int
	Body string
}

func (ev *EventView) Layout(g *gocui.Gui) error {
	v, err := g.SetView(ev.Name, ev.X, ev.Y, ev.X+ev.W, ev.Y+ev.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, ev.Body)
	}
	return nil
}

func (ev *EventView) SetPropreties(x, y, w, h int) {
	ev.X = x
	ev.Y = y
	ev.W = w
	ev.H = h
}

// func (w *Day) updateEventsView(g *gocui.Gui) {
// 	week, err := g.View("week") // Not good name hardcoded
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	buffer := week.Buffer()
//
// 	for _, v := range w.events {
// 		y := v.timeToPosisition(buffer) + 1
// 		h := v.durationToPosition()
// 		v.setPropreties(w.x+1, y, w.w-2, h)
// 		v.Layout(g)
// 		if y+h > w.y+w.h || y < w.y {
// 			g.DeleteView(v.name)
// 		}
// 	}
// }

// func (e *Event) durationToPosition() int {
// 	return int(e.duration * 2)
// }
//
// func (e *Event) timeToPosisition(buffer string) int {
// 	lines := strings.Split(buffer, "\n")
//
// 	for i, line := range lines {
// 		if line == e.time {
// 			return i
// 		}
// 	}
// 	return 0
// }

// =============== DayView ===============

type DayView struct {
	Name       string
	X, Y       int
	W, H       int
	Body       string
	EventViews []EventView
}

func NewDayView(name string, x, y, w, h int, body string, eventViews []EventView) *DayView {
	return &DayView{Name: name, X: x, Y: y, W: w, H: h, Body: body, EventViews: eventViews}
}

func (dv *DayView) Layout(g *gocui.Gui) error {
	v, err := g.SetView(dv.Name, dv.X, dv.Y, dv.X+dv.W, dv.Y+dv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, dv.Body)
	}
	// dv.updateEventsView(g)

	return nil
}

func (dv *DayView) SetPropreties(x, y, w, h int) {
	dv.X = x
	dv.Y = y
	dv.W = w
	dv.H = h
}

// =============== WeekView ===============

type WeekView struct {
	Name     string
	X, Y     int
	W, H     int
	Body     string
	DayViews []DayView
}

func NewWeekView(name string, x, y, w, h int, dayViews []DayView) *WeekView {
	return &WeekView{Name: name, X: x, Y: y, W: w, H: h, DayViews: dayViews}
}

func (wv *WeekView) SetPropreties(x, y, w, h int) {
	wv.X = x
	wv.Y = y
	wv.W = w
	wv.H = h
}

func (wv *WeekView) InitDayViews() {
	wv.DayViews[0] = *NewDayView("sunday", 0, 0, 0, 0, "", nil)
	wv.DayViews[1] = *NewDayView("monday", 0, 0, 0, 0, "", nil)
	wv.DayViews[2] = *NewDayView("tuesday", 0, 0, 0, 0, "", nil)
	wv.DayViews[3] = *NewDayView("wednesday", 0, 0, 0, 0, "", nil)
	wv.DayViews[4] = *NewDayView("thursday", 0, 0, 0, 0, "", nil)
	wv.DayViews[5] = *NewDayView("friday", 0, 0, 0, 0, "", nil)
	wv.DayViews[6] = *NewDayView("saturday", 0, 0, 0, 0, "", nil)
}

func (wv *WeekView) calculateDayViewWidth() int {
	width := math.Round(float64((wv.W-7)/7)) - border
	return int(width)
}

func (wv *WeekView) updateDayViews(g *gocui.Gui, x0 int, y0 int) {
	width := wv.calculateDayViewWidth()
	height := wv.H - y0 - 1
	x := x0
	for _, v := range wv.DayViews {
		g.DeleteView(v.Name)

		v.SetPropreties(x, y0, width, height)
		v.Layout(g)
		x += width + border
	}
}


func (wv *WeekView) Layout(g *gocui.Gui) error {
	v, err := g.SetView(wv.Name, wv.X, wv.Y, wv.X+wv.W, wv.Y+wv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, wv.Body)
	}

	y0 := 2
	wv.updateDayViews(g, 7, y0)
	wv.writeBody(v, y0)

	return nil
}

func (wv *WeekView) writeBody(v *gocui.View, y0 int) {
	height := wv.H - y0 - 3

	v.Clear()

	fmt.Fprintln(v, wv.Body)
	fmt.Fprintln(v)
	fmt.Fprintln(v)

	wv.writeTimes(v, height)
}

func (wv *WeekView) writeTimes(v *gocui.View, height int) {
	initialTime := 12 - height/4
	halfTime := 0

	for i := 0; i < int(height); i++ {
		fmt.Fprintf(v, "%02dh%02d\n", initialTime, halfTime)

		if halfTime == 0 {
			halfTime = 30
		} else {
			initialTime++
			halfTime = 0
		}
	}
}
