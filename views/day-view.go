package views

import (
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

type DayView struct {
	*BaseView

	Day *types.Day
}

func NewDayView(name string, d *types.Day) *DayView {
	dv := &DayView{
		BaseView: NewBaseView(name),
		Day:      d,
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

	return nil
}
