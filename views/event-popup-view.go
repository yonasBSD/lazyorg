package views

import (
	"github.com/jroimartin/gocui"
)

type EventPopupView struct {
	Properties UiProperties
	TextFields []TextField
	IsVisible bool
}

func NewEvenPopup(properties UiProperties) *EventPopupView {
	epv := &EventPopupView{Properties: properties,
		TextFields: []TextField{
            *NewTextField(UiProperties{Name: "Name"}, "", true),
            *NewTextField(UiProperties{Name: "Time"}, "", false),
            *NewTextField(UiProperties{Name: "Location"}, "", false),
            *NewTextField(UiProperties{Name: "Duration"}, "1.0", false),
            *NewTextField(UiProperties{Name: "Frequency"}, "7", false),
            *NewTextField(UiProperties{Name: "Occurence"}, "1", false),
            *NewTextField(UiProperties{Name: "Description"}, "", false),
		},
		IsVisible: false,
	}

	return epv
}

func (epv *EventPopupView) Show() {
	epv.IsVisible = true
}

func (epv *EventPopupView) Hide(g *gocui.Gui) error {
	var err error

	if epv.IsVisible {
		err = epv.RemovePopup(g)
	}

	epv.IsVisible = false

	return err
}

func (epv *EventPopupView) Update(g *gocui.Gui) error {

	v, err := g.SetView(
		epv.Properties.Name,
		epv.Properties.X,
		epv.Properties.Y,
		epv.Properties.X+epv.Properties.W,
		epv.Properties.Y+epv.Properties.H,
	)
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

func (epv *EventPopupView) RemovePopup(g *gocui.Gui) error {

	for _, v := range epv.TextFields {
		err := g.DeleteView(v.Properties.Name)
		if err != nil {
			return err
		}
	}

	err := g.DeleteView(epv.Properties.Name)
	if err != nil {
		return err
	}

	return nil
}

func (epv *EventPopupView) NextField() {
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

func (epv *EventPopupView) PrevField() {
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

func (epv *EventPopupView) UpdateViewOnTop(g *gocui.Gui) error {
	if _, err := g.SetViewOnTop(epv.Properties.Name); err != nil {
		return err
	}

	for i := range epv.TextFields {
		if _, err := g.SetViewOnTop(epv.TextFields[i].Properties.Name); err != nil {
			return err
		}
	}

	return nil
}

func (epv *EventPopupView) setTextFieldsProperties() {
	const (
		margin       = 2
		padding      = 1
		columnsCount = 3
		rowsCount    = 3
	)

	innerWidth := epv.Properties.W - 2*margin
	innerHeight := epv.Properties.H - 2*margin

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
		x := epv.Properties.X + margin + (item.col * (int(fieldWidth) + padding))
		y := epv.Properties.Y + margin + (item.row * (fieldHeight + padding))
		w := fieldWidth
		h := fieldHeight * item.spanRows

		if item.field == &epv.TextFields[6] {
			w = innerWidth
			h = innerHeight - (2 * fieldHeight) - padding
		}

		maxX := epv.Properties.X + epv.Properties.W - 1
		maxY := epv.Properties.Y + epv.Properties.H - 1

		if x < epv.Properties.X {
			x = epv.Properties.X
		}
		if y < epv.Properties.Y {
			y = epv.Properties.Y
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

		item.field.Properties.SetProperties(x, y, w, h)
	}
}

func (epv *EventPopupView) updateTextFields(g *gocui.Gui) error {
	for i := range epv.TextFields {
		if err := epv.TextFields[i].Update(g); err != nil {
			return err
		}
	}

	return nil
}
