package views

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

const (
	TimeViewWidth   = 9
	TitleViewHeight = 3
)

type MainView struct {
	*BaseView

	Calendar *types.Calendar
}

func NewMainView(c *types.Calendar) *MainView {
	mv := &MainView{
		BaseView: NewBaseView("main"),
        Calendar: c,
	}

	mv.AddChild("time", NewTimeView())
	mv.AddChild("week", NewWeekView(c))
	mv.AddChild("title", NewTitleView(c))

	return mv
}

func (mv *MainView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		mv.Name,
		mv.X,
		mv.Y,
		mv.X+mv.W,
		mv.Y+mv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
        v.FgColor = gocui.AttrBold
	}

	mv.updateChildViewProperties()

	if err = mv.UpdateChildren(g); err != nil {
		return err
	}

    if err = mv.updateCursorPosition(g); err != nil {
        return err
    }

	return nil
}

func (mv *MainView) updateCursorPosition(g *gocui.Gui) error {
    if view, ok := mv.GetChild("time"); ok {
        if timeView, ok := view.(*TimeView); ok {
            y := types.TimeToPosition(mv.Calendar.CurrentDay.Date, timeView.Body)

            g.SetCurrentView(weekdayNames[mv.Calendar.CurrentDay.Date.Weekday()])
            g.CurrentView().SetCursor(1, y)
        } else {
            return gocui.ErrUnknownView
        }
    } else {
        return gocui.ErrUnknownView
    }

    return nil
}

func (mv *MainView) updateChildViewProperties() {
	if timeView, ok := mv.GetChild("time"); ok {
		y := mv.Y + TitleViewHeight + 1
		timeView.SetProperties(
			mv.X+1,
			y,
			TimeViewWidth,
			mv.H-y-1,
		)
	}

	if weekView, ok := mv.GetChild("week"); ok {
		y := mv.Y + TitleViewHeight
		weekView.SetProperties(
			mv.X+TimeViewWidth+1,
			y,
			mv.W-TimeViewWidth-1,
			mv.H-y,
		)
	}

	if titleView, ok := mv.GetChild("title"); ok {
		titleView.SetProperties(
			mv.X,
			mv.Y,
			mv.W,
			TitleViewHeight,
		)
	}
}
