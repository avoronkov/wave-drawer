package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

	grid.OnUpdate = raster.Refresh

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			grid.Clear()
			raster.Refresh()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			grid.Normalize()

			d := dialog.NewFileSave(func(out fyne.URIWriteCloser, err error) {
				if out == nil || err != nil {
					return
				}
				defer out.Close()
				grid.Save(out)
			}, w)
			d.Show()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			grid.Normalize()
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	status := widget.NewLabel("OK")

	draggableRaster := NewInteractiveRaster(raster)
	draggableRaster.OnDragged = func(obj fyne.CanvasObject, e *fyne.DragEvent) {
		fin := e.Position
		start := e.Position.Subtract(e.Dragged)
		x2, y2 := draggableRaster.LocationForPosition(fin)
		x1, y1 := draggableRaster.LocationForPosition(start)
		// log.Printf("OnDragged %v, delta=%v : %v, %v -> %v, %v", e.Position, e.Dragged, x1, y1, x2, y2)
		grid.SetRange(y1, x1, y2, x2, true)
		obj.Refresh()
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

	piano, err := NewPiano(grid)
	if err != nil {
		log.Fatal(err)
	}
	pianoBar := NewPianoBar(piano)

	content := container.NewBorder(toolbar, status, nil, pianoBar, r)

	log.Printf("r.Position: %v", r.Position())

	w.SetContent(content)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}
