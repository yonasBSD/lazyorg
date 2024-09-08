package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	c := NewCalendar()

	days := []*Day{
		newDay("d1", nil),
		newDay("d2", nil),
		newDay("d3", nil),
		newDay("d4", nil),
		newDay("d5", nil),
		newDay("d6", nil),
		newDay("d7", nil),
	}

	w := newWeek("week", days, "Semaine 1", c)
	g.SetManager(w)

	if err := initKeybindings(g, w); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func initKeybindings(g *gocui.Gui, w *Week) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'H', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return w.prevWeek(g)
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'L', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return w.nextWeek(g)
		}); err != nil {
		return err
	}
	return nil

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
