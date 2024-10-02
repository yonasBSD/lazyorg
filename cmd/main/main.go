package main

import (
	"log"

	"github.com/HubertBel/go-organizer/cmd/types"
	_ "github.com/HubertBel/go-organizer/cmd/types"
	"github.com/HubertBel/go-organizer/views"
	"github.com/jroimartin/gocui"
)

func main() {
	path := "../../database.db"

	database := &types.Database{}
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

	av := views.NewAppView(g, database)
	g.SetManager(av)

	if err := views.InitKeybindings(g, av); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
