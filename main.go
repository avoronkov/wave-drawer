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

	grid := NewGrid(0, 0)

	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{0, 0, 128, 255}
	red := color.RGBA{128, 0, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	colorMapping := map[PState]color.Color{
		PEmpty:    white,
		PPoint:    black,
		PLine:     blue,
		PMainLine: red,
	}

	raster := canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			state := grid.GetState(y, x)
			return colorMapping[state]
		},
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			grid.Normalize()
			raster.Refresh()
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	draggableRaster := NewInteractiveRaster(raster)
	draggableRaster.OnDragged = func(obj fyne.CanvasObject, e *fyne.DragEvent) {
		x, y := draggableRaster.LocationForPosition(e.Position)
		// log.Printf("OnDragged %v -> %v, %v", e.Position, x, y)
		grid.Set(y, x, true)
		obj.Refresh()
	}
	draggableRaster.OnLayout = func(size fyne.Size) {
		pos := fyne.NewPos(size.Width, size.Height)
		w, h := draggableRaster.LocationForPosition(pos)
		grid.Resize(h, w)
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
