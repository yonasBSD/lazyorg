package views

import (
	"fmt"
	"math"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/jroimartin/gocui"
)

const (
	border = 1
)

type WeekView struct {
	Properties UiProperties
	Body       string

	TimeView   TimeView
	DayViews   []DayView
	EventViews []EventView

	EventPopupView EventPopupView

	Database *types.Database

	Calendar *types.Calendar
}

func NewWeekView(db *types.Database) *WeekView {
	wv := &WeekView{Properties: UiProperties{Name: "week"}, Database: db}
	wv.TimeView = *NewTimeView(UiProperties{Name: "time"}, "")
	wv.EventViews = make([]EventView, 0)
	wv.Calendar = types.NewCalendar(types.NewDay(time.Now(), nil))
	wv.EventPopupView = *NewEvenPopup(UiProperties{Name: "popup"})
	wv.initDayViews()
	return wv
}

func (wv *WeekView) Layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	wv.Properties.SetProperties(0, 0, maxX-1, maxY-1)
	wv.Body = wv.Calendar.FormatWeekBody()

	v, err := g.SetView(
		wv.Properties.Name,
		wv.Properties.X,
		wv.Properties.Y,
		wv.Properties.X+wv.Properties.W,
		wv.Properties.Y+wv.Properties.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Wrap = true
		v.Autoscroll = true
	}

	v.Clear()
	fmt.Fprintln(v, wv.Body)

	err = wv.updateTimeView(g)
	if err != nil {
		return err
	}

	err = wv.updateDayViews(g)
	if err != nil {
		return err
	}

	err = wv.updateEvents(g)
	if err != nil {
		return err
	}

	err = wv.updatePopupView(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) UpdateToNextWeek() {
	wv.Calendar.UpdateToNextWeek()
}

func (wv *WeekView) UpdateToPrevWeek() {
	wv.Calendar.UpdateToPrevWeek()
}

func (wv *WeekView) initDayViews() {
	wv.DayViews = make([]DayView, 7)
    wv.DayViews[0] = *NewDayView(UiProperties{Name: "sunday"}, "")
    wv.DayViews[1] = *NewDayView(UiProperties{Name: "monday"}, "")
    wv.DayViews[2] = *NewDayView(UiProperties{Name: "tuesday"}, "")
    wv.DayViews[3] = *NewDayView(UiProperties{Name: "wednesday"}, "")
    wv.DayViews[4] = *NewDayView(UiProperties{Name: "thursday"}, "")
    wv.DayViews[5] = *NewDayView(UiProperties{Name: "friday"}, "")
    wv.DayViews[6] = *NewDayView(UiProperties{Name: "saturday"}, "")
}

func (wv *WeekView) updateEvents(g *gocui.Gui) error {
	err := wv.Calendar.UpdateEventsFromDatabase(wv.Database)
	if err != nil {
		return err
	}

	wv.clearEventViews(g)
	wv.createEventViews()

	err = wv.updateEventViews(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) updatePopupView(g *gocui.Gui) error {
	if !wv.EventPopupView.IsVisible {
		return nil
	}

	w := wv.Properties.W / 2
	h := wv.Properties.H / 3

	x := wv.Properties.X + (wv.Properties.W/2 - w/2)
	y := wv.Properties.Y + (wv.Properties.H/2 - h/2)

	wv.EventPopupView.Properties.SetProperties(x, y, w, h)

	if err := wv.EventPopupView.Update(g); err != nil {
		return err
	}

	if err := wv.EventPopupView.UpdateViewOnTop(g); err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) updateTimeView(g *gocui.Gui) error {
	x := wv.Properties.X
	y := wv.Properties.Y + 3
	w := 8
	h := wv.Properties.H - wv.Properties.Y - 4

	wv.TimeView.Properties.SetProperties(x, y, w, h)
	err := wv.TimeView.Update(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) updateDayViews(g *gocui.Gui) error {
	x := wv.Properties.X + wv.TimeView.Properties.W + 1
	y := wv.Properties.Y + 3
	w := calculateDayViewWidth(x, wv.Properties.W)
	h := wv.Properties.H - wv.Properties.Y - 4

	for i := range wv.DayViews {
		wv.DayViews[i].Properties.SetProperties(x, y, w, h)
		wv.DayViews[i].Body = wv.Calendar.CurrentWeek.Days[i].FormatDayBody()

		err := wv.DayViews[i].Update(g)
		if err != nil {
			return err
		}

		x += w + border
	}
	return nil
}

func (wv *WeekView) clearEventViews(g *gocui.Gui) error {

	for _, v := range wv.EventViews {
		err := g.DeleteView(v.Properties.Name)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
	}
	wv.EventViews = nil

	return nil
}

func (wv *WeekView) createEventViews() {
	var eventViews []EventView

	for i, day := range wv.Calendar.CurrentWeek.Days {
		for _, event := range day.Events {
			x := wv.DayViews[i].Properties.X + 1
			y := wv.TimeView.TimeToPosition(event.FormatHour())
			w := wv.DayViews[i].Properties.W - 2
			h := wv.TimeView.DurationToHeight(event.DurationHour)

			name := fmt.Sprint(event.Name, event.Time.Day(), event.FormatHour())
			body := fmt.Sprint(event.Name, " | ", event.Location)

			if y > wv.DayViews[i].Properties.Y && y+h < wv.DayViews[i].Properties.Y+wv.DayViews[i].Properties.H {
                eventViews = append(eventViews, *NewEvenView(UiProperties{Name: name, X: x, Y: y, W: w, H: h}, body))
			}
		}
	}

	wv.EventViews = eventViews
}

func (wv *WeekView) updateEventViews(g *gocui.Gui) error {

	for i := range wv.EventViews {
		err := wv.EventViews[i].Update(g)
		if err != nil {
			return err
		}
	}

	return nil
}

func calculateDayViewWidth(x, w int) int {
	width := math.Round(float64((w-x)/7)) - border
	return int(width)
}

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
