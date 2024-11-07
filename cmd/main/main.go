package main

import (
	"log"

	"github.com/HubertBel/lazyorg/internal/database"
	"github.com/HubertBel/lazyorg/internal/ui"
	"github.com/HubertBel/lazyorg/pkg/views"
	"github.com/jroimartin/gocui"
)

func main() {
	path := "../../tmp/database.db"

	database := &database.Database{}
	err := database.InitDatabase(path)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Db.Close()

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	av := views.NewAppView(g, database)
	g.SetManager(av)

	if err := ui.InitKeybindings(g, av); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
