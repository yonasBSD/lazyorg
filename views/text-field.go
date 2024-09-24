package views

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type TextField struct {
	Properties UiProperties
	Body       string

	IsSelected bool
}

func NewTextField(properties UiProperties, body string, isSelected bool) *TextField {
	return &TextField{Properties: properties, Body: body, IsSelected: isSelected}
}

func (tf *TextField) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		tf.Properties.Name,
		tf.Properties.X,
		tf.Properties.Y,
		tf.Properties.X+tf.Properties.W,
		tf.Properties.Y+tf.Properties.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = tf.Properties.Name
		v.Editable = true
	}

	if tf.IsSelected {
		g.SetCurrentView(tf.Properties.Name)
	}

	fmt.Fprintln(v, tf.Body)

	return nil
}
