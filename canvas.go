package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type Canvas struct {
	widget.Entry

	texter SetTexter
}

var _ fyne.CanvasObject = (*Canvas)(nil)
var _ desktop.Mouseable = (*Canvas)(nil)
var _ mobile.Touchable = (*Canvas)(nil)
var _ fyne.Draggable = (*Canvas)(nil)

func NewCanvas(texter SetTexter) *Canvas {
	c := &Canvas{
		texter: texter,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *Canvas) MouseDown(e *desktop.MouseEvent) {
	c.texter.SetText(fmt.Sprintf("MouseDown: %v", e.Position))
}

func (c *Canvas) TouchDown(e *mobile.TouchEvent) {

	c.texter.SetText(fmt.Sprintf("TouchDown: %v", e.Position))
}

func (c *Canvas) Dragged(e *fyne.DragEvent) {
	c.texter.SetText(fmt.Sprintf("Dragged: %v, %v", e.Position, e.Dragged))
}

func (c *Canvas) DragEnd() {
	c.texter.SetText("DragEnd")
}

type SetTexter interface {
	SetText(s string)
}
