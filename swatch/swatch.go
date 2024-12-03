package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/hal-29/pixel/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, colcolor color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{
		Selected:     false,
		Color:        colcolor,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	s.Refresh()

	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  s,
	}

}
