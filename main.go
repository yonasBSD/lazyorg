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

	days := []*Day{
        newDay("d1", "Lundi"),
        newDay("d2", "Mardi"),
        newDay("d3", "Mercredi"),
        newDay("d4", "Jeudi"),
        newDay("d5", "Vendredi"),
        newDay("d6", "Samedi"),
        newDay("d7", "Dimanche"),
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
