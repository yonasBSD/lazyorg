package views

import (
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

const (
	MainViewWidth = 0.8
)

type AppView struct {
	*BaseView

	Calendar *types.Calendar
}

func NewAppView(g *gocui.Gui) *AppView {
	c := types.NewCalendar(types.NewDay(time.Now()))

	av := &AppView{
		BaseView: NewBaseView("app"),
		Calendar: c,
	}

	av.AddChild("main", NewMainView(c))
	av.AddChild("side", NewSideView())

	return av
}

func (av *AppView) Layout(g *gocui.Gui) error {
	return av.Update(g)
}

func (av *AppView) Update(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	av.SetProperties(0, 0, maxX-1, maxY-1)

	v, err := g.SetView(
		av.Name,
		av.X,
		av.Y,
		av.X+av.W,
		av.Y+av.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	av.updateChildViewProperties()

	if err = av.UpdateChildren(g); err != nil {
		return err
	}

	return nil
}

func (av *AppView) UpdateToNextWeek() {
	av.Calendar.UpdateToNextWeek()
}

func (av *AppView) UpdateToPrevWeek() {
	av.Calendar.UpdateToPrevWeek()
}

func (av *AppView) UpdateToNextDay() {
	av.Calendar.UpdateToNextDay()
}

func (av *AppView) UpdateToPrevDay() {
	av.Calendar.UpdateToPrevDay()
}

func (av *AppView) UpdateToNextTime() {
	av.Calendar.UpdateToNextTime()
}

func (av *AppView) UpdateToPrevTime() {
	av.Calendar.UpdateToPrevTime()
}

func (av *AppView) updateChildViewProperties() {
	if mainView, ok := av.GetChild("main"); ok {
		w := int(float64(av.W) * MainViewWidth)
		mainView.SetProperties(
			av.X+(av.W-w),
			av.Y,
			w,
			av.H,
		)
	}

	if weekView, ok := av.GetChild("side"); ok {
		weekView.SetProperties(
			av.X,
			av.Y,
			int(float64(av.W)*(1-MainViewWidth)),
			av.H,
		)
	}
}
