package views

import (
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

type DayView struct {
	*BaseView

	Day *types.Day
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

	dv.updateChildViewProperties(g)

	if err = dv.UpdateChildren(g); err != nil {
		return err
	}

	return nil
}

func (dv *DayView) updateChildViewProperties(g *gocui.Gui) error {
    if err := dv.ClearChildren(g); err != nil {
        return err
    }

	for _, v := range dv.Day.Events {
		ev := NewEvenView(v.Name, v)

		ev.X = dv.X + 1
        ev.Y = dv.Y + types.TimeToPosition(v.Time, dv.TimeView.Body) + 1
		ev.W = dv.W - 2
        ev.H = types.DurationToHeight(v.DurationHour)

		dv.AddChild(v.Name, ev)
	}

    return nil
}
