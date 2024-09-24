package views

import (
	"github.com/jroimartin/gocui"
)

type DayView struct {
	Properties UiProperties
	Body       string
	EventViews []EventView
}

func NewDayView(properties UiProperties, body string) *DayView {
	return &DayView{Properties: properties, Body: body}
}

func (dv *DayView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		dv.Properties.Name,
		dv.Properties.X,
		dv.Properties.Y,
		dv.Properties.X+dv.Properties.W,
		dv.Properties.Y+dv.Properties.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
	}

	v.Title = dv.Body

	return nil
}
