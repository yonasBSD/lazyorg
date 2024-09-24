package views

type UiProperties struct {
    Name string
    X, Y, W, H int
}

func (p *UiProperties) SetProperties(x, y, w, h int) {
    p.X = x
    p.Y = y
    p.W = w
    p.H = h
}
