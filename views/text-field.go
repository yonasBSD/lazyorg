package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type TextField struct {
	Name string
	X, Y int
	W, H int
	Body string

    IsSelected bool
}

func NewTextField(name string, x, y, w, h int, body string, isSelected bool) *TextField {
	return &TextField{Name: name, X: x, Y: y, W: w, H: h, Body: body, IsSelected: isSelected}
}

func (tf *TextField) SetProperties(x, y, w, h int) {
	tf.X = x
	tf.Y = y
	tf.W = w
	tf.H = h
}

func (tf *TextField) Update(g *gocui.Gui) error {
	v, err := g.SetView(tf.Name, tf.X, tf.Y, tf.X+tf.W, tf.Y+tf.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = tf.Name
        v.Editable = true
	}

    if (tf.IsSelected) {
        g.SetCurrentView(tf.Name)
    }

    fmt.Fprintln(v, tf.Body)

	return nil
}

type EventPopupView struct {
	Name string
	X, Y int
	W, H int

	TextFields []TextField

    IsVisible bool
}
