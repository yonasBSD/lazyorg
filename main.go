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

    events1 := []*Event{
        newEvent("e1", "Tennis"),
        newEvent("e2", "Cours"),
    }

    events2 := []*Event{
        newEvent("e3", "Epicerie"),
    }

    events3 := []*Event{
        newEvent("e4", "Ceci"),
        newEvent("e5", "est"),
        newEvent("e6", "un"),
        newEvent("e7", "test"),
    }

	days := []*Day{
		newDay("d1", []*Event{}, "Lundi"),
		newDay("d2", events1, "Mardi"),
		newDay("d3", events2, "Mercredi"),
		newDay("d4", []*Event{}, "Jeudi"),
        newDay("d5", events3, "Vendredi"),
		newDay("d6", []*Event{}, "Samedi"),
		newDay("d7", []*Event{}, "Dimanche"),
	}

	w := newWeek("w1", days, "Semaine 1")
	g.SetManager(w)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
