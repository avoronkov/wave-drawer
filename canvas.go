package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type Canvas struct {
	raster *canvas.Raster

	texter SetTexter
}

var _ fyne.CanvasObject = (*Canvas)(nil)
var _ desktop.Hoverable = (*Canvas)(nil)
var _ fyne.Draggable = (*Canvas)(nil)

func NewCanvas(texter SetTexter) *Canvas {
	c := &Canvas{
		texter: texter,
	}
	c.raster = canvas.NewRasterWithPixels(c.render)
	return c
}

func (c *Canvas) render(x, y, w, h int) color.Color {
	log.Printf("Render: %v, %v, %v, %v", x, y, w, h)
	// blue := color.RGBA{B: 128, A: 0xff}
	white := color.RGBA{R: 128, G: 128, B: 128, A: 255}
	// center := int(h / 2)
	col := white
	/*
		if y == center {
			col = blue
		}
	*/
	return col
}

func (c *Canvas) MouseDown(e *desktop.MouseEvent) {
	c.texter.SetText(fmt.Sprintf("MouseDown: %v", e.Position))
}

func (c *Canvas) MouseIn(e *desktop.MouseEvent) {
	c.texter.SetText(fmt.Sprintf("MouseIn: %v", e.Position))
}

func (c *Canvas) MouseMoved(e *desktop.MouseEvent) {
	c.texter.SetText(fmt.Sprintf("MouseMoved: %v", e.Position))
}

func (c *Canvas) MouseOut() {

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

// fyne.CanvasObject implementation
func (c *Canvas) Hide() {
	c.raster.Hide()
}

func (c *Canvas) MinSize() fyne.Size {
	return c.raster.MinSize()
}

func (c *Canvas) Move(pos fyne.Position) {
	c.raster.Move(pos)
}

func (c *Canvas) Position() fyne.Position {
	return c.raster.Position()
}

func (c *Canvas) Refresh() {
	c.raster.Refresh()
}

func (c *Canvas) Resize(size fyne.Size) {
	c.raster.Resize(size)
}

func (c *Canvas) Show() {
	c.raster.Show()
}

func (c *Canvas) Size() fyne.Size {
	return c.raster.Size()
}

func (c *Canvas) Visible() bool {
	return c.raster.Visible()
}

func (c *Canvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.raster)
}

type canvasRenderer struct {
	canvas *Canvas
}

var _ fyne.WidgetRenderer = (*canvasRenderer)(nil)

func (r *canvasRenderer) Layout(size fyne.Size) {
	r.canvas.raster.Resize(size)
}

func (r *canvasRenderer) Destroy() {
}

func (r *canvasRenderer) Objects() []fyne.CanvasObject {
	rect := canvas.NewRectangle(color.White)
	return []fyne.CanvasObject{rect}
}

func (r *canvasRenderer) MinSize() fyne.Size {
	return r.canvas.MinSize()
}

func (r *canvasRenderer) Refresh() {
	r.canvas.raster.Refresh()
}
