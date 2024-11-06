package views

import (
    "fmt"
    "github.com/jroimartin/gocui"
)

const (
    KeybindsWidth = 50
    KeybindsHeight = 20 
)

type KeybindsView struct {
    *BaseView
    IsVisible bool
}

func NewKeybindsView() *KeybindsView {
    return &KeybindsView{
        BaseView:  NewBaseView("keybinds"),
        IsVisible: false,
    }
}

func (kbv *KeybindsView) Update(g *gocui.Gui) error {
    if !kbv.IsVisible {
        return nil
    }

    v, err := g.SetView(
        kbv.Name,
        kbv.X,
        kbv.Y,
        kbv.X+kbv.W,
        kbv.Y+kbv.H,
    )
    if err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        v.Title = " Keybindings "
    }

    v.Clear()
    
    fmt.Fprintln(v, " Navigation:")
    fmt.Fprintln(v, " h/l        - Previous/Next day")
    fmt.Fprintln(v, " H/L        - Previous/Next week")
    fmt.Fprintln(v, " j/k        - Move time cursor down/up")
    fmt.Fprintln(v, "")
    fmt.Fprintln(v, " Events:")
    fmt.Fprintln(v, " a          - Add new event")
    fmt.Fprintln(v, " d          - Delete current event")
    fmt.Fprintln(v, " D          - Delete all events with same name")
    fmt.Fprintln(v, "")
    fmt.Fprintln(v, " View Controls:")
    fmt.Fprintln(v, " Ctrl+d     - Hide side view")
    fmt.Fprintln(v, " Ctrl+s     - Show side view")
    fmt.Fprintln(v, " N          - Open notepad")
    fmt.Fprintln(v, " ?          - Toggle this help")
    fmt.Fprintln(v, "")
    fmt.Fprintln(v, " Notepad:")
    fmt.Fprintln(v, " Ctrl+r     - Clear content")
    fmt.Fprintln(v, " Ctrl+q     - Return to main view")
    fmt.Fprintln(v, "")
    fmt.Fprintln(v, " Global:")
    fmt.Fprintln(v, " Ctrl+c     - Quit")

    return nil
}

func (kbv *KeybindsView) Show(g *gocui.Gui) error {
    kbv.IsVisible = true
    maxX, maxY := g.Size()
    
    kbv.SetProperties(
        (maxX-KeybindsWidth)/2,
        (maxY-KeybindsHeight)/2,
        KeybindsWidth,
        KeybindsHeight,
    )
    
    return kbv.Update(g)
}

func (kbv *KeybindsView) Hide(g *gocui.Gui) error {
    kbv.IsVisible = false
    return g.DeleteView(kbv.Name)
}
