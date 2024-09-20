package views

import (
	"github.com/jroimartin/gocui"
)

type EvenPopupView struct {
	Name string
	X, Y int
	W, H int
}

func NewEvenPopup(name string, x, y, w, h int) *EvenPopupView {
    return &EvenPopupView{Name: name, X: x, Y: y, W: w, H: h}
}

func (epv *EvenPopupView) SetProperties(x, y, w, h int) {
    epv.X = x
    epv.Y = y
    epv.W = w
    epv.H = h
}

func (epv *EvenPopupView) Update(g *gocui.Gui) error {

	v, err := g.SetView(epv.Name, epv.X, epv.Y, epv.X+epv.W, epv.Y+epv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
            return err
		}
        v.Title = "New event"
        g.SetCurrentView(epv.Name)
	}

	return nil
}

func (epv *EvenPopupView) RemovePopup(g *gocui.Gui) error {
    err := g.DeleteView(epv.Name)
    if err != nil {
        return err
    }

    return nil
}


