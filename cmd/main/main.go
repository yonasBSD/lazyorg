package main

import (
	"log"

	"github.com/HubertBel/go-organizer/cmd/database"
	"github.com/jroimartin/gocui"
)

func main() {
    path := "../database/database.db"

    database := database.Database{}
    err := database.InitDatabase(path)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Db.Close()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

    cc := NewCalendarController()
    cc.InitDatabase(path)

	g.SetManager(cc)

	if err := initKeybindings(g, cc); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func initKeybindings(g *gocui.Gui, cc *CalendarController) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'H', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return cc.UpdateToPrevWeek()
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'L', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return cc.UpdateToNextWeek()
		}); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
