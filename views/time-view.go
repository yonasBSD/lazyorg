package views

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

type TimeView struct {
	Properties UiProperties
	Body       string
}

func NewTimeView(properties UiProperties, body string) *TimeView {
	return &TimeView{Properties: properties, Body: body}
}

func (tv *TimeView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		tv.Properties.Name,
		tv.Properties.X,
		tv.Properties.Y,
		tv.Properties.X+tv.Properties.W,
		tv.Properties.Y+tv.Properties.H,
	)

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
			return i + tv.Properties.Y + 1
		}
	}

	return 0
}

func (tv *TimeView) DurationToHeight(d float64) int {
	return int(d * 2)
}

func (tv *TimeView) updateBody(v *gocui.View) {
	initialTime := 12 - tv.Properties.H/4
	halfTime := 0
	tv.Body = ""

	for range tv.Properties.H {
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
