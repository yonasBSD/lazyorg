package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type EventView struct {
	Properties UiProperties
	Body       string
}

func NewEvenView(properties UiProperties, body string) *EventView {
	return &EventView{Properties: properties, Body: body}
}

func (ev *EventView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		ev.Properties.Name,
		ev.Properties.X,
		ev.Properties.Y,
		ev.Properties.X+ev.Properties.W,
		ev.Properties.Y+ev.Properties.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		fmt.Fprintln(v, ev.Body)
	}

	return nil
}
