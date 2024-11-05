package views

import (
	"github.com/HubertBel/go-organizer/internal/calendar"
	"github.com/jroimartin/gocui"
)

type EventView struct {
    *BaseView

    Event *calendar.Event
}

func NewEvenView(name string, e *calendar.Event) *EventView {
	return &EventView {
        BaseView: NewBaseView(name),

        Event: e,
    }
}

func (ev *EventView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		ev.Name,
		ev.X,
		ev.Y,
		ev.X+ev.W,
		ev.Y+ev.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
        v.Title = ev.Event.Name
	}

	return nil
}
