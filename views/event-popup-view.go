package views

import (
	"github.com/jroimartin/gocui"
)

func NewEvenPopup(name string, x, y, w, h int) *EvenPopupView {
	epv := &EvenPopupView{Name: name, X: x, Y: y, W: w, H: h,
		TextFields: []TextField{
            *NewTextField("Name", 0, 0, 0, 0, "", true),
			*NewTextField("Time", 0, 0, 0, 0, "", false),
			*NewTextField("Location", 0, 0, 0, 0, "", false),
			*NewTextField("Duration", 0, 0, 0, 0, "1.0", false),
			*NewTextField("Frequency", 0, 0, 0, 0, "7", false),
			*NewTextField("Occurence", 0, 0, 0, 0, "1", false),
			*NewTextField("Description", 0, 0, 0, 0, "", false),
		},
        IsVisible: false,
	}

	return epv
}

func (epv *EvenPopupView) SetProperties(x, y, w, h int) {
	epv.X = x
	epv.Y = y
	epv.W = w
	epv.H = h
}

func (epv *EvenPopupView) Show() {
    epv.IsVisible = true
}

func (epv *EvenPopupView) Hide(g *gocui.Gui) error {
    var err error

    if epv.IsVisible {
        err = epv.RemovePopup(g)
    }

    epv.IsVisible = false

    return err
}

func (epv *EvenPopupView) Update(g *gocui.Gui) error {

	v, err := g.SetView(epv.Name, epv.X, epv.Y, epv.X+epv.W, epv.Y+epv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "New event"
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

func (epv *EvenPopupView) NextField() {
    currentField := 0

    for i, v := range epv.TextFields {
        if v.IsSelected {
            currentField = i
        }
    }

    epv.TextFields[currentField].IsSelected = false
    if currentField == len(epv.TextFields)-1 {
        epv.TextFields[0].IsSelected = true
    } else {
        epv.TextFields[currentField+1].IsSelected = true
    }
}

func (epv *EvenPopupView) PrevField() {
    currentField := 0

    for i, v := range epv.TextFields {
        if v.IsSelected {
            currentField = i
        }
    }

    epv.TextFields[currentField].IsSelected = false
    if currentField == 0 {
        epv.TextFields[len(epv.TextFields)-1].IsSelected = true
    } else {
        epv.TextFields[currentField-1].IsSelected = true
    }
}

func (epv *EvenPopupView) UpdateViewOnTop(g *gocui.Gui) error {
	if _, err := g.SetViewOnTop(epv.Name); err != nil {
		return err
	}

	for i := range epv.TextFields {
		if _, err := g.SetViewOnTop(epv.TextFields[i].Name); err != nil {
			return err
		}
	}

	return nil
}

func (epv *EvenPopupView) setTextFieldsProperties() {
	const (
		margin       = 2
		padding      = 1
		columnsCount = 3
		rowsCount    = 3
	)

	innerWidth := epv.W - 2*margin
	innerHeight := epv.H - 2*margin

	fieldWidth := (innerWidth / columnsCount) - padding
	fieldHeight := (innerHeight / rowsCount) - padding

	layout := []struct {
		field    *TextField
		row, col int
		spanRows int
	}{
		{&epv.TextFields[0], 0, 0, 1}, // Name
		{&epv.TextFields[1], 0, 1, 1}, // Time
		{&epv.TextFields[2], 0, 2, 1}, // Location
		{&epv.TextFields[3], 1, 0, 1}, // Duration
		{&epv.TextFields[4], 1, 1, 1}, // Frequency
		{&epv.TextFields[5], 1, 2, 1}, // Occurrence
		{&epv.TextFields[6], 2, 0, 1}, // Description
	}

	for _, item := range layout {
		x := epv.X + margin + (item.col * (int(fieldWidth) + padding))
		y := epv.Y + margin + (item.row * (fieldHeight + padding))
		w := fieldWidth
		h := fieldHeight * item.spanRows

		if item.field == &epv.TextFields[6] {
			w = innerWidth
			h = innerHeight - (2 * fieldHeight) - padding
		}

		maxX := epv.X + epv.W - 1
		maxY := epv.Y + epv.H - 1

		if x < epv.X {
			x = epv.X
		}
		if y < epv.Y {
			y = epv.Y
		}
		if x+w > maxX {
			w = maxX - x
		}
		if y+h > maxY {
			h = maxY - y
		}

		if w < 1 {
			w = 1
		}
		if h < 1 {
			h = 1
		}

		item.field.SetProperties(x, y, w, h)
	}
}

func (epv *EvenPopupView) updateTextFields(g *gocui.Gui) error {
	for i := range epv.TextFields {
		if err := epv.TextFields[i].Update(g); err != nil {
			return err
		}
	}

	return nil
}
