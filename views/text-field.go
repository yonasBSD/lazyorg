package views

import (
	"github.com/jroimartin/gocui"
)

type TextField struct {
	*BaseView
	Title string
	Body  string

	IsSelected bool
}

func NewTextField(name, title, body string) *TextField {
	tf := &TextField{
		BaseView:   NewBaseView(name),
		Title:      title,
		Body:       "",
		IsSelected: false,
	}

	return tf
}

func (tf *TextField) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		tf.Name,
		tf.X,
		tf.Y,
		tf.X+tf.W,
		tf.Y+tf.H,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = tf.Title
		v.Editable = true
	}

	if tf.IsSelected {
		g.SetCurrentView(tf.Name)
	}

	// fmt.Fprintln(v, tf.Body)

	return nil
}
