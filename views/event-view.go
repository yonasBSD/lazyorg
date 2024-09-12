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

func NewEvenView(name string, x, y, w, h int, body string) *EventView {
	return &EventView{Name: name, X: x, Y: y, W: w, H: h, Body: body}
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
