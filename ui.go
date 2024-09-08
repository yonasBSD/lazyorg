package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
)

type Event struct {
	name     string
	x, y     int
	w, h     int
	body     string
	time     string
	duration float64
}

type Day struct {
	name   string
	x, y   int
	w, h   int
	body   string
	events []*Event
}

type Week struct {
	name string
	x, y int
	w, h int
	body string
    calendar Calendar
	days []*Day
}

func newEvent(name string, body string, time string, duration float64) *Event {
	return &Event{name: name, x: 0, y: 0, w: 0, h: 0, body: body, time: time, duration: duration}
}

func (e *Event) durationToPosition() int {
	return int(e.duration * 2)
}

func (e *Event) timeToPosisition(buffer string) int {
	lines := strings.Split(buffer, "\n")

	for i, line := range lines {
		if line == e.time {
			return i
		}
	}
	return 0
}

func (w *Event) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, w.body)
	}
	return nil
}

func (e *Event) setPropreties(x, y, w, h int) {
	e.x = x
	e.y = y
	e.w = w
	e.h = h
}

func newDay(name string, events []*Event) *Day {
	return &Day{name: name, x: 0, y: 0, w: 0, h: 0, body: "", events: events}
}

func (w *Day) updateEventsView(g *gocui.Gui) {
	week, err := g.View("w1")
	if err != nil {
		log.Fatal(err)
	}

	buffer := week.Buffer()

	for _, v := range w.events {
		y := v.timeToPosisition(buffer) + 1
		h := v.durationToPosition()
		v.setPropreties(w.x+1, y, w.w-2, h)
        v.Layout(g)
        if y+h > w.y+w.h || y < w.y {
            g.DeleteView(v.name)
        }
	}
}

func (w *Day) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, w.body)
	}
	w.updateEventsView(g)

	return nil
}

func (d *Day) setPropreties(x, y, w, h int, body string) {
	d.x = x
	d.y = y
	d.w = w
	d.h = h
    d.body = body
}

func newWeek(name string, days []*Day, body string, calendar Calendar) *Week {
	return &Week{name: name, x: 0, y: 0, w: 0, h: 0, body: body, days: days, calendar: calendar}
}

func (we *Week) setPropreties(x, y, w, h int, body string) {
	we.x = x
	we.y = y
	we.w = w
	we.h = h
    we.body = body
}

func (w *Week) getDayDimensions() (width, border int) {
	numberOfDay := len(w.days)
	n := math.Round(float64((w.w - 7) / numberOfDay))
	b := 1.0
	wi := n - b

	return int(wi), int(b)
}

func (w *Week) updateDaysView(g *gocui.Gui, x0 int) {
	width, border := w.getDayDimensions()
	y0 := 2
	x := x0
	for i, v := range w.days {
		v.setPropreties(x, y0, width, w.h-y0-1, w.calendar.daysName[i])
		v.Layout(g)
		x += width + border
	}
}

func (w *Week) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

    startDay := w.calendar.startDay.Day()
    endDay := w.calendar.endDay.Day()

    month := w.calendar.endDay.Month().String()

    body := "Week: " + strconv.Itoa(startDay) + " to " + strconv.Itoa(endDay) + ", " + month
	w.setPropreties(0, 0, maxX-1, maxY-1, body)

	view, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	w.updateDaysView(g, 7)
	w.writeTime(view)

	return nil
}

func (w *Week) writeTime(v *gocui.View) {
	height := w.days[0].h

	v.Clear()

	fmt.Fprintln(v, w.body)
	fmt.Fprintln(v)
	fmt.Fprintln(v)

	initialTime := 13 - height/4
	halfTime := 0
	for i := 0; i < int(height)-2; i++ {
		fmt.Fprintf(v, "%02dh%02d\n", initialTime, halfTime)
		if halfTime == 0 {
			halfTime = 30
		} else {
			initialTime++
			halfTime = 0
		}
	}
}
