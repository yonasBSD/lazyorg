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
}

func NewTextField(name string, x, y, w, h int, body string) *TextField {
	return &TextField{Name: name, X: x, Y: y, W: w, H: h, Body: body}
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
		fmt.Fprintln(v, tf.Body)
	}

	return nil
}

type EvenPopupView struct {
	Name string
	X, Y int
	W, H int

	TextFields []TextField
}

func NewEvenPopup(name string, x, y, w, h int) *EvenPopupView {
	epv := &EvenPopupView{Name: name, X: x, Y: y, W: w, H: h,
		TextFields: []TextField{
			*NewTextField("Name", 0, 0, 0, 0, ""),
			*NewTextField("Time", 0, 0, 0, 0, ""),
			*NewTextField("Location", 0, 0, 0, 0, ""),
			*NewTextField("Duration", 0, 0, 0, 0, ""),
			*NewTextField("Frequency", 0, 0, 0, 0, ""),
			*NewTextField("Occurence", 0, 0, 0, 0, ""),
			*NewTextField("Description", 0, 0, 0, 0, ""),
		},
	}

	return epv
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

	epv.setTextFieldsProperties()
	epv.updateTextFields(g)

	return nil
}

func (epv *EvenPopupView) RemovePopup(g *gocui.Gui) error {
	for _, v := range epv.TextFields {
		err := g.DeleteView(v.Name)
		if err != nil {
			return err
		}
	}

	err := g.DeleteView(epv.Name)
	if err != nil {
		return err
	}

	return nil
}

func (epv *EvenPopupView) UpdateViewOnTop(g *gocui.Gui) error {
	if _, err := g.SetViewOnTop(epv.Name); err != nil {
		return err
	}

	for _, v := range epv.TextFields {
		if _, err := g.SetViewOnTop(v.Name); err != nil {
			return err
		}
	}

	return nil
}

func (epv *EvenPopupView) setTextFieldsProperties() {

	b := 1

	x := epv.X + 2
	y := epv.Y + 1
	w := (epv.W/3) - (2*b)
	h := (epv.H/3) - (2*b)

	// Name
	epv.TextFields[0].SetProperties(x, y, w, h)

	// Time
	x += w + b
	epv.TextFields[1].SetProperties(x, y, w, h)

	// Location
	x += w + b
	epv.TextFields[2].SetProperties(x, y, w, h)

	// Duration
	x = epv.X + 2
	y += h + b
	epv.TextFields[3].SetProperties(x, y, w, h)

	// Frequency
	x += w + b
	epv.TextFields[4].SetProperties(x, y, w, h)

	// Occurence
	x += w + b
	epv.TextFields[5].SetProperties(x, y, w, h)

	// Description
	x = epv.X + 2
	y += h + b
    h = ((epv.Y+epv.H) - y) - 1
	epv.TextFields[6].SetProperties(x, y, epv.W-4, h)
}

func (epv *EvenPopupView) updateTextFields(g *gocui.Gui) error {
	for _, v := range epv.TextFields {
		if err := v.Update(g); err != nil {
			return err
		}
	}

	return nil
}
