package views

import (
	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

var weekdayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

const (
	padding = 1
)

type WeekView struct {
	*BaseView

	Calendar *types.Calendar

	TimeView *TimeView

	// DayViews []DayView
	// EventViews []EventView

	// EventPopupView EventPopupView

	// Database *types.Database
}

func NewWeekView(c *types.Calendar, tv *TimeView) *WeekView {
	wv := &WeekView{
        BaseView: NewBaseView("week"),
        Calendar: c,
		TimeView: tv,
	}

    // TODO
	wv.AddChild(weekdayNames[0], NewDayView(weekdayNames[0], c.CurrentWeek.Days[0], tv))
	wv.AddChild(weekdayNames[1], NewDayView(weekdayNames[1], c.CurrentWeek.Days[1], tv))
	wv.AddChild(weekdayNames[2], NewDayView(weekdayNames[2], c.CurrentWeek.Days[2], tv))
	wv.AddChild(weekdayNames[3], NewDayView(weekdayNames[3], c.CurrentWeek.Days[3], tv))
	wv.AddChild(weekdayNames[4], NewDayView(weekdayNames[4], c.CurrentWeek.Days[4], tv))
	wv.AddChild(weekdayNames[5], NewDayView(weekdayNames[5], c.CurrentWeek.Days[5], tv))
	wv.AddChild(weekdayNames[6], NewDayView(weekdayNames[6], c.CurrentWeek.Days[6], tv))

	// wv.TimeView = *NewTimeView(UiProperties{Name: "time"}, "")
	// wv.EventViews = make([]EventView, 0)
	// wv.EventPopupView = *NewEvenPopup(UiProperties{Name: "popup"})

	return wv
}

func (wv *WeekView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		wv.Name,
		wv.X,
		wv.Y,
		wv.X+wv.W,
		wv.Y+wv.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	wv.updateChildViewProperties()

	if err = wv.UpdateChildren(g); err != nil {
		return err
	}

	// if err = wv.updateDayViews(g); err != nil {
	// 	return err
	// }

	// err = wv.updateEvents(g)
	// if err != nil {
	// 	return err
	// }

	// err = wv.updatePopupView(g)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (wv *WeekView) updateChildViewProperties() {
	x := wv.X
	w := wv.W/7 - padding

	for _, weekday := range weekdayNames {
		if dayView, ok := wv.GetChild(weekday); ok {

			dayView.SetProperties(
				x,
				wv.Y+1,
				w,
				wv.H-2,
			)
		}

		x += w + padding
	}
}

// func (wv *WeekView) updateEvents(g *gocui.Gui) error {
// 	err := wv.Calendar.UpdateEventsFromDatabase(wv.Database)
// 	if err != nil {
// 		return err
// 	}
//
// 	wv.clearEventViews(g)
// 	wv.createEventViews()
//
// 	err = wv.updateEventViews(g)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// func (wv *WeekView) updatePopupView(g *gocui.Gui) error {
// 	if !wv.EventPopupView.IsVisible {
// 		return nil
// 	}
//
// 	w := wv.Properties.W / 2
// 	h := wv.Properties.H / 3
//
// 	x := wv.Properties.X + (wv.Properties.W/2 - w/2)
// 	y := wv.Properties.Y + (wv.Properties.H/2 - h/2)
//
// 	wv.EventPopupView.Properties.SetProperties(x, y, w, h)
//
// 	if err := wv.EventPopupView.Update(g); err != nil {
// 		return err
// 	}
//
// 	if err := wv.EventPopupView.UpdateViewOnTop(g); err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// func (wv *WeekView) updateDayViews(g *gocui.Gui) error {
// 	x := wv.Properties.X
// 	y := wv.Properties.Y + 1
// 	w := wv.Properties.W/7 - padding
//     h := wv.Properties.H - 2
//
// 	for i := range wv.DayViews {
// 		wv.DayViews[i].Properties.SetProperties(x, y, w, h)
// 		wv.DayViews[i].Body = wv.Calendar.CurrentWeek.Days[i].FormatDayBody()
//
// 		if err := wv.DayViews[i].Update(g); err != nil {
// 			return err
// 		}
//
// 		x += w + padding
// 	}
// 	return nil
// }

// func (wv *WeekView) clearEventViews(g *gocui.Gui) error {
//
// 	for _, v := range wv.EventViews {
// 		err := g.DeleteView(v.Properties.Name)
// 		if err != nil && err != gocui.ErrUnknownView {
// 			return err
// 		}
// 	}
// 	wv.EventViews = nil
//
// 	return nil
// }

// func (wv *WeekView) createEventViews() {
// 	var eventViews []EventView
//
// 	for i, day := range wv.Calendar.CurrentWeek.Days {
// 		for _, event := range day.Events {
// 			x := wv.DayViews[i].Properties.X + 1
// 			y := wv.TimeView.TimeToPosition(event.FormatHour())
// 			w := wv.DayViews[i].Properties.W - 2
// 			h := wv.TimeView.DurationToHeight(event.DurationHour)
//
// 			name := fmt.Sprint(event.Name, event.Time.Day(), event.FormatHour())
// 			body := fmt.Sprint(event.Name, " | ", event.Location)
//
// 			if y > wv.DayViews[i].Properties.Y && y+h < wv.DayViews[i].Properties.Y+wv.DayViews[i].Properties.H {
// 				eventViews = append(eventViews, *NewEvenView(UiProperties{Name: name, X: x, Y: y, W: w, H: h}, body))
// 			}
// 		}
// 	}
//
// 	wv.EventViews = eventViews
// }

// func (wv *WeekView) updateEventViews(g *gocui.Gui) error {
//
// 	for i := range wv.EventViews {
// 		err := wv.EventViews[i].Update(g)
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }

// func (wv *WeekView) AddTestEvents() error {
// 	t := time.Date(2024, time.September, 16, 10, 30, 0, 0, time.Now().Location())
// 	e := types.NewEvent("Archi", "Architecture des microprocesseurs", "PLT-2510", t, 2.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	t = time.Date(2024, time.September, 19, 13, 30, 0, 0, time.Now().Location())
// 	e = types.NewEvent("Archi", "Architecture des microprocesseurs", "PLT-2700", t, 2.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	t = time.Date(2024, time.September, 16, 13, 30, 0, 0, time.Now().Location())
// 	e = types.NewEvent("Astrophysique", "Introduction a l'astrophysique", "VCH-2840", t, 1.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	t = time.Date(2024, time.September, 18, 13, 30, 0, 0, time.Now().Location())
// 	e = types.NewEvent("Astrophysique", "Introduction a l'astrophysique", "VCH-3850", t, 2.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	t = time.Date(2024, time.September, 18, 15, 30, 0, 0, time.Now().Location())
// 	e = types.NewEvent("Russe", "Russe elementaire 1", "En ligne", t, 3.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	t = time.Date(2024, time.September, 19, 9, 30, 0, 0, time.Now().Location())
// 	e = types.NewEvent("Robotique", "Introduction a la robotique mobile", "PLT-2765", t, 3.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	t = time.Date(2024, time.September, 20, 10, 30, 0, 0, time.Now().Location())
// 	e = types.NewEvent("Robotique", "Introduction a la robotique mobile", "PLT-3928", t, 2.0, 7, 13)
// 	wv.Database.AddRecurringEvents(e)
//
// 	return nil
// }
