package views

import (

	"github.com/jroimartin/gocui"
)

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
        v.Wrap = true
        v.Autoscroll = true
	}

    v.Title = dv.Body
	// dv.updateEventsView(g)

	return nil
}

func (dv *DayView) SetPropreties(x, y, w, h int) {
	dv.X = x
	dv.Y = y
	dv.W = w
	dv.H = h
}

