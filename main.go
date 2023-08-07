package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	grid := NewGrid(3200, 3200)

	scale := float32(1000.0 / 840.0)
	log.Printf("Scale=%v", scale)

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{0, 0, 128, 255}
	red := color.RGBA{128, 0, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			v := white
			if x%50 == 0 || y%50 == 0 {
				v = blue
			}
			if x%1000 == 0 || y%1000 == 0 {
				v = red
			}
			if grid.Get(y, x) {
				v = black
			}
			return v
		})

	log.Printf("ScaleMode: %v", raster.ScaleMode)

	draggableRaster := MakeDraggableCanvasObject(raster, func(obj fyne.CanvasObject, e *fyne.DragEvent) {
		log.Printf("OnDragged: pos=%v, abs=%v, size=%v, (%v)", e.Position, e.AbsolutePosition, obj.Size(), e.Dragged)
		p := e.Position
		grid.Set(int(p.Y*scale), int(p.X*scale), true)
		obj.Refresh()
	})
	useDraggable := true

	if useDraggable {
		w.SetContent(draggableRaster)
	} else {
		w.SetContent(raster)
	}
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}
