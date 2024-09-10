package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

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
