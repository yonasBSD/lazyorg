package main

import (
	"log"

	"github.com/HubertBel/go-organizer/cmd/types"
	"github.com/HubertBel/go-organizer/views"
	"github.com/jroimartin/gocui"
)

func main() {
    path := "../../database.db"

    database := types.Database{}
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

    wv := views.NewWeekView(&database)
	g.SetManager(wv)

	if err := initKeybindings(g, wv); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func initKeybindings(g *gocui.Gui, wv *views.WeekView) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'H', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			wv.UpdateToPrevWeek()
            return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'L', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			wv.UpdateToNextWeek()
            return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'a', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
            return wv.EvenPopupView.Update(g)
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding(wv.EvenPopupView.Name, 'b', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
            if _, err := g.View(wv.EvenPopupView.Name); err == nil {
                return wv.EvenPopupView.RemovePopup(g)
            }
            return nil
		}); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
