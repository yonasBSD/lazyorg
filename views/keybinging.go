package views

import (
	"github.com/jroimartin/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func InitKeybindings(g *gocui.Gui, av *AppView) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'h', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			av.UpdateToPrevDay()
			return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'l', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			av.UpdateToNextDay()
			return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'j', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			av.UpdateToNextTime()
			return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'k', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			av.UpdateToPrevTime()
			return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'H', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			av.UpdateToPrevWeek()
			return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'L', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			av.UpdateToNextWeek()
			return nil
		}); err != nil {
		return err
	}

	// if err := g.SetKeybinding("", 'a', gocui.ModNone,
	// 	func(g *gocui.Gui, v *gocui.View) error {
	// 		wv.EventPopupView.Show()
	// 		return nil
	// 	}); err != nil {
	// 	return err
	// }

	// if err := g.SetKeybinding("", 'b', gocui.ModNone,
	// 	func(g *gocui.Gui, v *gocui.View) error {
	// 		return wv.EventPopupView.Hide(g)
	// 	}); err != nil {
	// 	return err
	// }

	// if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
	// 	func(g *gocui.Gui, v *gocui.View) error {
	// 		if wv.EventPopupView.IsVisible {
	//             wv.EventPopupView.PrevField()
	// 		}
	//         return nil
	// 	}); err != nil {
	// 	return err
	// }

	// if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
	// 	func(g *gocui.Gui, v *gocui.View) error {
	// 		if wv.EventPopupView.IsVisible {
	//             wv.EventPopupView.NextField()
	// 		}
	//         return nil
	// 	}); err != nil {
	// 	return err
	// }

	return nil
}
