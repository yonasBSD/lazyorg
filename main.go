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

	d1Events := []*Event{}
	d2Events := []*Event{}
	d3Events := []*Event{}
	d4Events := []*Event{}
	d5Events := []*Event{}
	d6Events := []*Event{}
	d7Events := []*Event{}

	days := []*Day{
		newDay("d1", d1Events),
		newDay("d2", d2Events),
		newDay("d3", d3Events),
		newDay("d4", d4Events),
		newDay("d5", d5Events),
		newDay("d6", d6Events),
		newDay("d7", d7Events),
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

	if err := g.SetKeybinding("", gocui.KeyCtrlU, gocui.ModNone, // Change keybind
		func(g *gocui.Gui, v *gocui.View) error {
			return w.prevWeek(g)
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlO, gocui.ModNone, // Change keybind
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
