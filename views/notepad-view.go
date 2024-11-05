package views

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)


type NotepadView struct {
	*BaseView
}

func NewNotepadView(c *types.Calendar) *NotepadView {
	tv := &NotepadView{
        BaseView: NewBaseView("notepad"),
	}

	return tv
}

func (tv *NotepadView) Update(g *gocui.Gui) error {
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
        v.Title = "Notepad"
        v.Editable = true
        v.FgColor = gocui.AttrBold
	}

	return nil
}
