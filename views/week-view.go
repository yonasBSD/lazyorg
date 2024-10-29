package views

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

var weekdayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

const (
	padding = 1
)

type WeekView struct {
	*BaseView

	Calendar *types.Calendar

	TimeView *TimeView
}

func NewWeekView(c *types.Calendar, tv *TimeView) *WeekView {
	wv := &WeekView{
		BaseView: NewBaseView("week"),
		Calendar: c,
		TimeView: tv,
	}

	wv.AddChild(weekdayNames[0], NewDayView(weekdayNames[0], c.CurrentWeek.Days[0], tv))
	wv.AddChild(weekdayNames[1], NewDayView(weekdayNames[1], c.CurrentWeek.Days[1], tv))
	wv.AddChild(weekdayNames[2], NewDayView(weekdayNames[2], c.CurrentWeek.Days[2], tv))
	wv.AddChild(weekdayNames[3], NewDayView(weekdayNames[3], c.CurrentWeek.Days[3], tv))
	wv.AddChild(weekdayNames[4], NewDayView(weekdayNames[4], c.CurrentWeek.Days[4], tv))
	wv.AddChild(weekdayNames[5], NewDayView(weekdayNames[5], c.CurrentWeek.Days[5], tv))
	wv.AddChild(weekdayNames[6], NewDayView(weekdayNames[6], c.CurrentWeek.Days[6], tv))

	return wv
}

func (wv *WeekView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		wv.Name,
		wv.X,
		wv.Y,
		wv.X+wv.W,
		wv.Y+wv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	wv.updateChildViewProperties()

	if err = wv.UpdateChildren(g); err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) updateChildViewProperties() {
	x := wv.X
	w := wv.W/7 - padding

	for _, weekday := range weekdayNames {
		if dayView, ok := wv.GetChild(weekday); ok {

			dayView.SetProperties(
				x,
				wv.Y+1,
				w,
				wv.H-2,
			)
		}

		x += w + padding
	}
}
