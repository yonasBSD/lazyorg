package views

import "github.com/jroimartin/gocui"

type SideView struct {
    *BaseView

    // TODO
}

func NewSideView() *SideView {
    sv := &SideView{
        BaseView: NewBaseView("side"),
    }

    return sv
}

func (sv *SideView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		sv.Name,
		sv.X,
		sv.Y,
		sv.X+sv.W,
		sv.Y+sv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
        v.FgColor = gocui.AttrBold
	}

	return nil
}
