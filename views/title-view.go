package views

import (
	"fmt"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

type TitleView struct {
    *BaseView

    Calendar *types.Calendar
}

func NewTitleView(c *types.Calendar) *TitleView {
    tv := &TitleView{
        BaseView: NewBaseView("title"),
        Calendar: c,
    }

    return tv
}

func (tv *TitleView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		tv.Name,
		tv.X,
		tv.Y,
		tv.X+tv.W,
		tv.Y+tv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
        v.FgColor = gocui.AttrBold | gocui.ColorCyan
        v.Wrap = true
	}

    tv.updateBody(v)
    
	return nil
}

func (tv *TitleView) updateBody(v *gocui.View) {
    today := time.Now()
    selectedWeek := tv.Calendar.FormatWeekBody()
    todayString := fmt.Sprintf("%s %d - %02dh%02d", today.Month().String(), today.Day(), today.Hour(), today.Minute())

    title := fmt.Sprintf("%s | %s", selectedWeek, todayString)

    v.Clear()
	fmt.Fprintln(v, title)
}
