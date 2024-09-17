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

func NewDayView(name string, x, y, w, h int, body string) *DayView {
	return &DayView{Name: name, X: x, Y: y, W: w, H: h, Body: body}
}

func (dv *DayView) SetProperties(x, y, w, h int) {
	dv.X = x
	dv.Y = y
	dv.W = w
	dv.H = h
}

func (dv *DayView) Update(g *gocui.Gui) error {
	v, err := g.SetView(dv.Name, dv.X, dv.Y, dv.X+dv.W, dv.Y+dv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
	}

	v.Title = dv.Body

	// for _, v := range dv.EventViews {
    //     err := v.Layout(g)
    //     if err != nil {
    //         return err
    //     }
	// }

	return nil
}
