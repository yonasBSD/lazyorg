package views

import (
	"github.com/jroimartin/gocui"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type View interface {
	Update(g *gocui.Gui) error
	SetProperties(x, y, w, h int)
	GetName() string
	GetProperties() (int, int, int, int)
	AddChild(name string, child View) (View, bool)
	GetChild(name string) (View, bool)
	ClearChildren(g *gocui.Gui) error
	Children() *orderedmap.OrderedMap[string, View]
}

type BaseView struct {
	Name       string
	X, Y, W, H int
	children   *orderedmap.OrderedMap[string, View]
}

func NewBaseView(name string) *BaseView {
	return &BaseView{
		Name:     name,
		children: orderedmap.New[string, View](),
	}
}

func (bv *BaseView) SetProperties(x, y, w, h int) {
	bv.X, bv.Y, bv.W, bv.H = x, y, w, h
}

func (bv *BaseView) GetProperties() (int, int, int, int) {
	return bv.X, bv.Y, bv.W, bv.H
}

func (bv *BaseView) GetName() string {
	return bv.Name
}

func (bv *BaseView) ClearChildren(g *gocui.Gui) error {
	for pair := bv.children.Oldest(); pair != nil; pair = pair.Next() {
		if err := g.DeleteView(pair.Value.GetName()); err != nil {
			return err
		}
		bv.children.Delete(pair.Key)
	}

	return nil
}

func (bv *BaseView) AddChild(name string, child View) (View, bool) {
	child, ok := bv.children.Set(name, child)
	return child, ok
}

func (bv *BaseView) GetChild(name string) (View, bool) {
	child, ok := bv.children.Get(name)
	return child, ok
}

func (bv *BaseView) Children() *orderedmap.OrderedMap[string, View] {
	return bv.children
}

func (bv *BaseView) UpdateChildren(g *gocui.Gui) error {
	for pair := bv.children.Oldest(); pair != nil; pair = pair.Next() {
		if err := pair.Value.Update(g); err != nil {
			return err
		}
	}

	return nil
}
