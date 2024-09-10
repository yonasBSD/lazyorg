package views

import (
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

func NewWeekView(name string, x, y, w, h int, dayViews []DayView) *WeekView {
	wv := &WeekView{Name: name, X: x, Y: y, W: w, H: h, DayViews: dayViews}
	wv.InitDayViews()
	return wv
}

func (wv *WeekView) SetPropreties(x, y, w, h int) {
	wv.X = x
	wv.Y = y
	wv.W = w
	wv.H = h
}

func (wv *WeekView) InitDayViews() {
	wv.DayViews[0] = *NewDayView("sunday", 0, 0, 0, 0, "", nil)
	wv.DayViews[1] = *NewDayView("monday", 0, 0, 0, 0, "", nil)
	wv.DayViews[2] = *NewDayView("tuesday", 0, 0, 0, 0, "", nil)
	wv.DayViews[3] = *NewDayView("wednesday", 0, 0, 0, 0, "", nil)
	wv.DayViews[4] = *NewDayView("thursday", 0, 0, 0, 0, "", nil)
	wv.DayViews[5] = *NewDayView("friday", 0, 0, 0, 0, "", nil)
	wv.DayViews[6] = *NewDayView("saturday", 0, 0, 0, 0, "", nil)
}

func (wv *WeekView) calculateDayViewWidth() int {
	width := math.Round(float64((wv.W-7)/7)) - border
	return int(width)
}

func (wv *WeekView) updateDayViews(g *gocui.Gui, x0 int, y0 int) error {
	width := wv.calculateDayViewWidth()
	height := wv.H - y0 - 1
	x := x0
	for _, v := range wv.DayViews {
		v.SetPropreties(x, y0, width, height)
		v.Layout(g)
		x += width + border
	}
	return nil
}

func (wv *WeekView) Layout(g *gocui.Gui) error {
	v, err := g.SetView(wv.Name, wv.X, wv.Y, wv.X+wv.W, wv.Y+wv.H)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Autoscroll = true
	}

    y0 := wv.Y + 2
    x0 := wv.X + 7
	err = wv.updateDayViews(g, x0, y0)
	if err != nil {
		return err
	}
	wv.writeBody(v, y0)

	return nil
}

func (wv *WeekView) writeBody(v *gocui.View, y0 int) {
	height := wv.H - y0 - 2

    v.Title = wv.Body

	v.Clear()

	fmt.Fprintln(v)
	fmt.Fprintln(v)
	fmt.Fprintln(v)

	wv.writeTimes(v, height)
}

func (wv *WeekView) writeTimes(v *gocui.View, height int) {
	initialTime := 12 - height/4
	halfTime := 0

	for i := 0; i < int(height); i++ {
		fmt.Fprintf(v, "%02dh%02d\n", initialTime, halfTime)

		if halfTime == 0 {
			halfTime = 30
		} else {
			initialTime++
			halfTime = 0
		}
	}
}
