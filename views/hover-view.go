package views

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)


type HoverView struct {
	*BaseView

	Calendar *types.Calendar
    CurrentView View
}

func NewHoverView(c *types.Calendar) *HoverView {
	hv := &HoverView{
        BaseView: NewBaseView("hover"),
        Calendar: c,
	}

	return hv
}

func (hv *HoverView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		hv.Name,
		hv.X,
		hv.Y,
		hv.X+hv.W,
		hv.Y+hv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

    if hv.CurrentView != nil {
        v.Title = hv.CurrentView.GetName()
    }

    hv.updateBody(v)

	return nil
}

func (hv *HoverView) updateBody(v *gocui.View) {

}
