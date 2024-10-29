package views

import (
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

var (
	MainViewWidthRatio = 0.8
	SideViewWidthRatio = 0.2
)

const (
    TitleViewHeight = 3

    PopupWidth  = LabelWidth + FieldWidth
    PopupHeight = 16
)

type AppView struct {
	*BaseView

	Database *types.Database
	Calendar *types.Calendar
}

func NewAppView(g *gocui.Gui, db *types.Database) *AppView {
    now := time.Now()
    t := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())

	c := types.NewCalendar(types.NewDay(t))

	av := &AppView{
		BaseView: NewBaseView("app"),
		Database: db,
		Calendar: c,
	}

	av.AddChild("title", NewTitleView(c))
	av.AddChild("popup", NewEvenPopup(g, c, db))
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

	if err = av.Calendar.UpdateEventsFromDatabase(av.Database); err != nil {
		return err
	}

	av.updateChildViewProperties()

	if err = av.UpdateChildren(g); err != nil {
		return err
	}

	if err = av.updateCurrentView(g); err != nil {
		return err
	}

	return nil
}

func (av *AppView) HideSideView(g *gocui.Gui) {
	SideViewWidthRatio = 0.0
	MainViewWidthRatio = 1.0

    av.children.Delete("side")
    g.DeleteView("side")
}

func (av *AppView) ShowSideView() {
	SideViewWidthRatio = 0.2
	MainViewWidthRatio = 0.8

	av.AddChild("side", NewSideView())
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

func (av *AppView) ShowPopup(g *gocui.Gui) error {
	if view, ok := av.GetChild("popup"); ok {
		view.SetProperties(
			av.X+(av.W/2-PopupWidth/2),
			av.Y+(av.H/2-PopupHeight/2),
			PopupWidth,
			PopupHeight,
		)
		if popupView, ok := view.(*EventPopupView); ok {
			return popupView.Show(g)
		}
	}
	return nil
}

func (av *AppView) updateChildViewProperties() {
	mainViewWidth := int(float64(av.W-1) * MainViewWidthRatio)
	sideViewWidth := int(float64(av.W) * SideViewWidthRatio)

	if titleView, ok := av.GetChild("title"); ok {
		titleView.SetProperties(
			av.X+sideViewWidth+1,
			av.Y,
			mainViewWidth,
			TitleViewHeight,
		)
	}

	if mainView, ok := av.GetChild("main"); ok {
		mainView.SetProperties(
			av.X+sideViewWidth+1,
			TitleViewHeight+1,
			mainViewWidth,
			av.H-TitleViewHeight-1,
		)
	}

	if sideView, ok := av.GetChild("side"); ok {
		sideView.SetProperties(
			av.X,
			av.Y,
			sideViewWidth,
			av.H,
		)
	}
}

func (av *AppView) updateCurrentView(g *gocui.Gui) error {
	if view, ok := av.GetChild("popup"); ok {
		if popupView, ok := view.(*EventPopupView); ok {
			if popupView.IsVisible {
				return nil
			}
		}
	}

	var mainView View
	var ok bool

	g.Cursor = true

	if mainView, ok = av.GetChild("main"); !ok {
		return gocui.ErrUnknownView
	}

	if view, ok := mainView.GetChild("time"); ok {
		if timeView, ok := view.(*TimeView); ok {
			y := types.TimeToPosition(av.Calendar.CurrentDay.Date, timeView.Body)
            if y == -1 {
                y = 0
            }

			g.SetCurrentView(weekdayNames[av.Calendar.CurrentDay.Date.Weekday()])
			g.CurrentView().SetCursor(1, y)

		} else {
			return gocui.ErrUnknownView
		}
	} else {
		return gocui.ErrUnknownView
	}

	return nil
}
