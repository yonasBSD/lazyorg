package views

import "github.com/jroimartin/gocui"

type TimeView struct {
    Name string
    X int
    Y int
    W int
    H int
    Body string
}

func NewTimeView(name string, x, y, w, h int, body string) *TimeView {
    return &TimeView{Name: name, X: x, Y: y, W: w, H: h, Body: body}
}

func (tv *TimeView) Update(g *gocui.Gui) error {
    // g.SetView
    // etc..
    return nil
}
