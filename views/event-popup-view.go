package views

import (
	"strconv"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/j-04/gocui-component"
	"github.com/jroimartin/gocui"
)

const (
	LabelWidth = 12
	FieldWidth = 20

	TimeFormat = "2006-01-02 15:04"
)

type EventPopupView struct {
	*BaseView
	Form     *component.Form
	Calendar *types.Calendar
	Database *types.Database

	IsVisible bool
}

func NewEvenPopup(g *gocui.Gui, c *types.Calendar, db *types.Database) *EventPopupView {

	epv := &EventPopupView{
		BaseView:  NewBaseView("popup"),
		Form:      nil,
		Calendar:  c,
		Database:  db,
		IsVisible: false,
	}

	return epv
}

func (epv *EventPopupView) Update(g *gocui.Gui) error {
	return nil
}

func (epv *EventPopupView) Show(g *gocui.Gui) error {
	if epv.IsVisible {
		return nil
	}

	form := component.NewForm(g, "New Event", epv.X, epv.Y, epv.W, epv.H)

	form.AddInputField("Name", LabelWidth, FieldWidth)
	form.AddInputField("Time", LabelWidth, FieldWidth).SetText(epv.Calendar.CurrentDay.Date.Format(TimeFormat))
	form.AddInputField("Location", LabelWidth, FieldWidth)
	form.AddInputField("Duration", LabelWidth, FieldWidth).SetText("1.0")
	form.AddInputField("Frequency", LabelWidth, FieldWidth).SetText("7")
	form.AddInputField("Occurence", LabelWidth, FieldWidth).SetText("1")
	form.AddInputField("Description", LabelWidth, FieldWidth)

	form.AddButton("Add", epv.AddEvent)
	form.AddButton("Cancel", epv.Close)

	form.SetCurrentItem(0)

	form.Draw()

	epv.Form = form
	epv.IsVisible = true

	return nil
}

func (epv *EventPopupView) AddEvent(g *gocui.Gui, v *gocui.View) error {

	name := epv.Form.GetFieldText("Name")
	time, _ := time.Parse(TimeFormat, epv.Form.GetFieldText("Time"))
	location := epv.Form.GetFieldText("Location")

	duration, _ := strconv.ParseFloat(epv.Form.GetFieldText("Duration"), 64)
	frequency, _ := strconv.Atoi(epv.Form.GetFieldText("Frequency"))
	occurence, _ := strconv.Atoi(epv.Form.GetFieldText("Occurence"))

	description := epv.Form.GetFieldText("Description")

	event := types.NewEvent(name, description, location, time, duration, frequency, occurence)

    if _, err := epv.Database.AddRecurringEvents(event); err != nil {
        return err
    }

    if err := epv.Calendar.UpdateEventsFromDatabase(epv.Database); err != nil {
        return err
    }

	return epv.Close(g, v)
}

func (epv *EventPopupView) Close(g *gocui.Gui, v *gocui.View) error {
	epv.IsVisible = false
	return epv.Form.Close(g, v)
}

//
//
// func (epv *EventPopupView) RemovePopup(g *gocui.Gui) error {
//
// 	for _, v := range epv.Children() {
// 		err := g.DeleteView(v.GetName())
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	err := g.DeleteView(epv.Name)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func (epv *EventPopupView) NextField() {
// 	currentField := 0
//
// 	for i, v := range epv.TextFields {
// 		if v.IsSelected {
// 			currentField = i
// 		}
// 	}
//
// 	epv.TextFields[currentField].IsSelected = false
// 	if currentField == len(epv.TextFields)-1 {
// 		epv.TextFields[0].IsSelected = true
// 	} else {
// 		epv.TextFields[currentField+1].IsSelected = true
// 	}
// }
//
// func (epv *EventPopupView) PrevField() {
// 	currentField := 0
//
// 	for i, v := range epv.TextFields {
// 		if v.IsSelected {
// 			currentField = i
// 		}
// 	}
//
// 	epv.TextFields[currentField].IsSelected = false
// 	if currentField == 0 {
// 		epv.TextFields[len(epv.TextFields)-1].IsSelected = true
// 	} else {
// 		epv.TextFields[currentField-1].IsSelected = true
// 	}
// }
//
// func (epv *EventPopupView) UpdateViewOnTop(g *gocui.Gui) error {
// 	if _, err := g.SetViewOnTop(epv.Properties.Name); err != nil {
// 		return err
// 	}
//
// 	for i := range epv.TextFields {
// 		if _, err := g.SetViewOnTop(epv.TextFields[i].Properties.Name); err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
