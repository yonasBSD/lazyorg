package views

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)


type TodoView struct {
	*BaseView

	Calendar *types.Calendar
}

func NewTodoView(c *types.Calendar) *TodoView {
	tv := &TodoView{
        BaseView: NewBaseView("todo"),
        Calendar: c,
	}

	return tv
}

func (tv *TodoView) Update(g *gocui.Gui) error {
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
        v.Title = "TODO"
	}

    tv.updateBody(v)

	return nil
}

func (tv *TodoView) updateBody(v *gocui.View) {

}
