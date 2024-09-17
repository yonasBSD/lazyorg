package views

import (
	// "bytes"
	"fmt"
	"math"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"

	// "math"

	"github.com/jroimartin/gocui"
)

const (
	border = 1
)

type WeekView struct {
	Name string
	X, Y int
	W, H int
	Body string

	TimeView   TimeView
	DayViews   []DayView
	EventViews []EventView

	Calendar *types.Calendar
}

func NewWeekView() *WeekView {
	wv := &WeekView{Name: "week"}
    wv.Calendar = types.NewCalendar(types.NewDay(time.Now(), nil))
	wv.initDayViews()
	return wv
}

func (wv *WeekView) SetProperties(x, y, w, h int) {
	wv.X = x
	wv.Y = y
	wv.W = w
	wv.H = h

	// wv.updateDayViews(x+9, y+3)
}

func (wv *WeekView) Layout(g *gocui.Gui) error {
    maxX, maxY := g.Size()
    wv.SetProperties(0, 0, maxX-1, maxY-1)
    wv.Body = wv.Calendar.FormatWeekBody()

	v, err := g.SetView(wv.Name, wv.X, wv.Y, wv.X+wv.W, wv.Y+wv.H)
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

	err = wv.updateDayViews(g)
	if err != nil {
		return err
	}

	return nil
}

func (wv *WeekView) UpdateToNextWeek() error {
	wv.Calendar.UpdateToNextWeek()
	return nil // TODO error handling
}

func (wv *WeekView) UpdateToPrevWeek() error {
	wv.Calendar.UpdateToPrevWeek()
	return nil // TODO error handling
}


func (wv *WeekView) initDayViews() {
	wv.DayViews = make([]DayView, 7)
	wv.DayViews[0] = *NewDayView("sunday", 0, 0, 0, 0, "")
	wv.DayViews[1] = *NewDayView("monday", 0, 0, 0, 0, "")
	wv.DayViews[2] = *NewDayView("tuesday", 0, 0, 0, 0, "")
	wv.DayViews[3] = *NewDayView("wednesday", 0, 0, 0, 0, "")
	wv.DayViews[4] = *NewDayView("thursday", 0, 0, 0, 0, "")
	wv.DayViews[5] = *NewDayView("friday", 0, 0, 0, 0, "")
	wv.DayViews[6] = *NewDayView("saturday", 0, 0, 0, 0, "")
}

func (wv *WeekView) updateDayViews(g *gocui.Gui) error {
	x := wv.X + 1 // Add time size
	y := wv.Y + 3
	w := calculateDayViewWidth(x, wv.W)
	h := wv.H - wv.Y - 4

	for i := range wv.DayViews {
		wv.DayViews[i].SetProperties(x, y, w, h)
        wv.DayViews[i].Body = wv.Calendar.CurrentWeek.Days[i].FormatDayBody()

		err := wv.DayViews[i].Update(g)
		if err != nil {
			return err
		}

		x += w + border
	}
	return nil
}

func calculateDayViewWidth(x, w int) int {
	width := math.Round(float64((w-x)/7)) - border
	return int(width)
}

// func (wv *WeekView) WriteTimes(b *bytes.Buffer, height int) {
// 	initialTime := 12 - height/4
// 	halfTime := 0
//
// 	for i := 0; i < int(height); i++ {
//
// 		if halfTime == 0 {
//             s := fmt.Sprintf("%02dh%02d -\n", initialTime, halfTime)
//             b.WriteString(s)
//
// 			halfTime = 30
// 		} else {
//             s := fmt.Sprintf("%02dh%02d\n", initialTime, halfTime)
//             b.WriteString(s)
//
// 			initialTime++
// 			halfTime = 0
// 		}
// 	}
// }
