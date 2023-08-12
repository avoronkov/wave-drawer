package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type interactiveRaster struct {
	widget.BaseWidget

	min fyne.Size
	img *canvas.Raster

	OnDragged func(obj fyne.CanvasObject, e *fyne.DragEvent)
}

var _ fyne.Draggable = (*interactiveRaster)(nil)

func NewInteractiveRaster(raster *canvas.Raster) *interactiveRaster {
	r := &interactiveRaster{
		img: raster,
	}
	r.ExtendBaseWidget(r)
	return r
}

func (r *interactiveRaster) SetMinSize(size fyne.Size) {
	log.Printf("SetMinSize %v", size)
	pixWidth, _ := r.LocationForPosition(fyne.NewPos(size.Width, size.Height))
	scale := float32(1.0)
	c := fyne.CurrentApp().Driver().CanvasForObject(r.img)
	if c != nil {
		scale = c.Scale()
	}

	texScale := float32(pixWidth) / size.Width / scale
	size = fyne.NewSize(size.Width/texScale, size.Height/texScale)
	r.min = size
	r.Resize(size)
}

func (r *interactiveRaster) MinSize() fyne.Size {
	return r.min
}

func (r *interactiveRaster) CreateRenderer() fyne.WidgetRenderer {
	return &rasterWidgetRender{raster: r, bg: canvas.NewRasterWithPixels(bgPattern)}
}

func (r *interactiveRaster) Tapped(ev *fyne.PointEvent) {
}

func (r *interactiveRaster) TappedSecondary(*fyne.PointEvent) {
}

func (r *interactiveRaster) LocationForPosition(pos fyne.Position) (int, int) {
	c := fyne.CurrentApp().Driver().CanvasForObject(r.img)
	x, y := int(pos.X), int(pos.Y)
	if c != nil {
		x, y = c.PixelCoordinateForPosition(pos)
	}

	return x, y
}

type rasterWidgetRender struct {
	raster *interactiveRaster
	bg     *canvas.Raster
}

func bgPattern(x, y, _, _ int) color.Color {
	const boxSize = 25

	if (x/boxSize)%2 == (y/boxSize)%2 {
		return color.Gray{Y: 58}
	}

	return color.Gray{Y: 84}
}

func (r *rasterWidgetRender) Layout(size fyne.Size) {
	r.bg.Resize(size)
	r.raster.img.Resize(size)
}

func (r *rasterWidgetRender) MinSize() fyne.Size {
	return r.MinSize()
}

func (r *rasterWidgetRender) Refresh() {
	canvas.Refresh(r.raster)
}

func (r *rasterWidgetRender) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *rasterWidgetRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.raster.img}
}

func (r *rasterWidgetRender) Destroy() {
}

func (r *interactiveRaster) Dragged(e *fyne.DragEvent) {
	if r.OnDragged != nil {
		r.OnDragged(r, e)
	}
}

func (r *interactiveRaster) DragEnd() {

}
