package views

import (
	"fmt"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

type DayView struct {
	*BaseView

	Day      *types.Day
	TimeView *TimeView
}

func NewDayView(name string, d *types.Day, tv *TimeView) *DayView {
	dv := &DayView{
		BaseView: NewBaseView(name),
		Day:      d,
		TimeView: tv,
	}

	return dv
}

func (dv *DayView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		dv.Name,
		dv.X,
		dv.Y,
		dv.X+dv.W,
		dv.Y+dv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if dv.Day.Date.YearDay() == time.Now().YearDay() {
		v.BgColor = gocui.ColorBlack
	} else {
		v.BgColor = gocui.ColorDefault
	}

	v.Title = dv.Day.FormatDayBody()

	if err = dv.updateChildViewProperties(g); err != nil {
		return err
	}

	if err = dv.UpdateChildren(g); err != nil {
		return err
	}

	return nil
}

func (dv *DayView) updateChildViewProperties(g *gocui.Gui) error {
	eventViews := make(map[string]*EventView)
	for pair := dv.children.Oldest(); pair != nil; pair = pair.Next() {
		if eventView, ok := pair.Value.(*EventView); ok {
			eventViews[eventView.GetName()] = eventView
		}
	}

	for _, event := range dv.Day.Events {
		x := dv.X + 1
		y := dv.Y + types.TimeToPosition(event.Time, dv.TimeView.Body) + 1
		w := dv.W - 2
		h := types.DurationToHeight(event.DurationHour)

		if (y + h) >= (dv.Y + dv.H) {
            newHeight := (dv.Y + dv.H) - y
            if newHeight <= 0{
                continue
            }
            h = newHeight
		}
		if y <= dv.Y {
			continue
		}

		viewName := fmt.Sprintf("%s-%d", event.Name, event.Id)
		if existingView, exists := eventViews[viewName]; exists {
			existingView.X, existingView.Y, existingView.W, existingView.H = x, y, w, h
			delete(eventViews, viewName)
		} else {
			ev := NewEvenView(viewName, event)
			ev.X, ev.Y, ev.W, ev.H = x, y, w, h
			dv.AddChild(viewName, ev)
		}
	}

	for viewName := range eventViews {
		if err := g.DeleteView(viewName); err != nil && err != gocui.ErrUnknownView {
			return err
		}
		dv.children.Delete(viewName)
	}

	return nil
}
