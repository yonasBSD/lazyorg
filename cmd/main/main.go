package main

import (
	"log"

	_ "github.com/HubertBel/go-organizer/cmd/types"
	"github.com/HubertBel/go-organizer/views"
	"github.com/jroimartin/gocui"
)

func main() {
	// path := "../../database.db"

	// database := types.Database{}
	// err := database.InitDatabase(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Db.Close()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true

	// wv := views.NewWeekView(&database)
	av := views.NewAppView(g)
	g.SetManager(av)

	if err := views.InitKeybindings(g, av); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
