package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

type TimeView struct {
	Name string
	X    int
	Y    int
	W    int
	H    int
	Body string
}

func NewTimeView(name string, x, y, w, h int, body string) *TimeView {
	return &TimeView{Name: name, X: x, Y: y, W: w, H: h, Body: body}
}

func (tv *TimeView) SetProperties(x, y, w, h int) {
	tv.X = x
	tv.Y = y
	tv.W = w
	tv.H = h
}

func (tv *TimeView) Update(g *gocui.Gui) error {
	v, err := g.SetView(tv.Name, tv.X, tv.Y, tv.X+tv.W, tv.Y+tv.H)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

    v.Clear()
	tv.updateBody(v)

	return nil
}

func (tv *TimeView) TimeToPosition(t string) int {
 
	lines := strings.Split(tv.Body, "\n")

	for i, v := range lines {
		if strings.Contains(v, t) {
			return i + tv.Y + 1
		}
	}

	return 0
}

func (tv *TimeView) DurationToHeight(d float64) int {
	return int(d * 2)
}

func (tv *TimeView) updateBody(v *gocui.View) {
	initialTime := 12 - tv.H/4
	halfTime := 0
    tv.Body = ""

	for range tv.H {
		if halfTime == 0 {
			tv.Body += fmt.Sprintf("%02dh%02d - \n", initialTime, halfTime)
			halfTime = 30
		} else {
			tv.Body += fmt.Sprintf("%02dh%02d\n", initialTime, halfTime)
			initialTime++
			halfTime = 0
		}
	}

	fmt.Fprintln(v, tv.Body)
}

