package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	grid := NewGrid(3200, 3200)

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{0, 0, 128, 255}
	red := color.RGBA{128, 0, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

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

	draggableRaster := NewInteractiveRaster(raster)
	draggableRaster.OnDragged = func(obj fyne.CanvasObject, e *fyne.DragEvent) {
		x, y := draggableRaster.LocationForPosition(e.Position)
		// log.Printf("OnDragged %v -> %v, %v", e.Position, x, y)
		grid.Set(y, x, true)
		obj.Refresh()
	}

	var r fyne.CanvasObject = raster
	useDraggable := true
	if useDraggable {
		r = draggableRaster
	}

	content := container.NewBorder(toolbar, nil, nil, nil, r)

	log.Printf("r.Position: %v", r.Position())

	w.SetContent(content)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}
