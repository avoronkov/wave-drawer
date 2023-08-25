package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Bezzier")

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

	grid.OnUpdate = raster.Refresh

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			grid.Clear()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			grid.Bezier()
		}),
		widget.NewToolbarAction(theme.ViewRestoreIcon(), func() {
			grid.Lagrange()
		}),
	)
	status := widget.NewLabel("OK")

	draggableRaster := NewInteractiveRaster(raster)
	draggableRaster.OnDown = func(e fyne.PointEvent) {
		pos := e.Position
		x1, y1 := draggableRaster.LocationForPosition(pos)
		grid.Set(y1, x1, true)
	}
	draggableRaster.OnLayout = func(size fyne.Size) {
		var w, h int
		defer func() {
			if r := recover(); r != nil {
				status.SetText(fmt.Sprintf("FATAL size=%v Resize(%v, %v): %v", size, h, w, r))
			}
		}()
		pos := fyne.NewPos(size.Width, size.Height)
		w, h = draggableRaster.LocationForPosition(pos)
		if h < 0 || w < 0 {
			return
		}
		grid.Resize(h, w)
		status.SetText(fmt.Sprintf("Resized: w=%v, h=%v", w, h))
	}

	var r fyne.CanvasObject = raster
	useDraggable := true
	if useDraggable {
		r = draggableRaster
	}

	content := container.NewBorder(toolbar, status, nil, nil, r)

	w.SetContent(content)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}
