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

	d1Events := []*Event{
		newEvent("e1", "Frisbee", "18h00", 1),
	}
	d2Events := []*Event{
		newEvent("e2", "Archi\n2510", "10h30", 2),
		newEvent("e3", "Astrophysique\n2840", "13h30", 1),
	}
	d3Events := []*Event{
		newEvent("e4", "Tennis\nPeps", "16h30", 2),
	}
	d4Events := []*Event{
		newEvent("e5", "Astrophysique\n3850", "13h30", 2),
		newEvent("e6", "Russe\nEn Ligne", "15h30", 3),
	}
	d5Events := []*Event{
		newEvent("e7", "Robotique\n2750", "9h30", 3),
		newEvent("e8", "Archi\n2700", "13h30", 2),
	}
	d6Events := []*Event{
		newEvent("e9", "Robotique\n3928", "10h30", 2),
	}
	d7Events := []*Event{}

	days := []*Day{
		newDay("d1", d1Events, "Dimanche"),
		newDay("d2", d2Events, "Lundi"),
		newDay("d3", d3Events, "Mardi"),
		newDay("d4", d4Events, "Mercredi"),
		newDay("d5", d5Events, "Jeudi"),
		newDay("d6", d6Events, "Vendredi"),
		newDay("d7", d7Events, "Samedi"),
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
