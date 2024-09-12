package views

import (
	"bytes"
	"fmt"
	"math"

	"github.com/jroimartin/gocui"
)

const (
	border = 1
)

type WeekView struct {
	Name     string
	X, Y     int
	W, H     int
	Body     string
	DayViews []DayView
}

func NewWeekView(name string, x, y, w, h int) *WeekView {
	wv := &WeekView{Name: name, X: x, Y: y, W: w, H: h, DayViews: make([]DayView, 7)}
	wv.initDayViews()
	return wv
}

func (wv *WeekView) Update(x, y, w, h int) {
	wv.X = x
	wv.Y = y
	wv.W = w
	wv.H = h

	wv.updateDayViews(x+9, y+3)
}

func (wv *WeekView) Layout(g *gocui.Gui) error {
	v, err := g.SetView(wv.Name, wv.X, wv.Y, wv.X+wv.W, wv.Y+wv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Wrap = true
		v.Autoscroll = true
	}

	for _, v := range wv.DayViews {
		err := v.Layout(g)
		if err != nil {
			return err
		}
	}

    v.Clear()
    fmt.Fprintln(v, wv.Body)

	return nil
}

func (wv *WeekView) initDayViews() {
	wv.DayViews[0] = *NewDayView("sunday", 0, 0, 0, 0, "", nil)
	wv.DayViews[1] = *NewDayView("monday", 0, 0, 0, 0, "", nil)
	wv.DayViews[2] = *NewDayView("tuesday", 0, 0, 0, 0, "", nil)
	wv.DayViews[3] = *NewDayView("wednesday", 0, 0, 0, 0, "", nil)
	wv.DayViews[4] = *NewDayView("thursday", 0, 0, 0, 0, "", nil)
	wv.DayViews[5] = *NewDayView("friday", 0, 0, 0, 0, "", nil)
	wv.DayViews[6] = *NewDayView("saturday", 0, 0, 0, 0, "", nil)
}

func (wv *WeekView) updateDayViews(x, y int) {
	w := calculateDayViewWidth(x, wv.W)
	h := wv.H - wv.Y - 4
	for i := range wv.DayViews {
		wv.DayViews[i].Update(x, y, w, h)
		x += w + border
	}
}

func calculateDayViewWidth(x, w int) int {
	width := math.Round(float64((w-x)/7)) - border
	return int(width)
}

func (wv *WeekView) WriteTimes(b *bytes.Buffer, height int) {
	initialTime := 12 - height/4
	halfTime := 0

	for i := 0; i < int(height); i++ {

		if halfTime == 0 {
            s := fmt.Sprintf("%02dh%02d -\n", initialTime, halfTime)
            b.WriteString(s)

			halfTime = 30
		} else {
            s := fmt.Sprintf("%02dh%02d\n", initialTime, halfTime)
            b.WriteString(s)

			initialTime++
			halfTime = 0
		}
	}
}
