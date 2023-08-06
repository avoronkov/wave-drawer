package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	label := widget.NewLabel("Hello World!")

	canvas := NewCanvas(label)

	button := widget.NewButton("Click me!", func() {
		label.SetText(fmt.Sprintf("Welcome: %v", canvas.Size()))
	})

	content := container.New(layout.NewGridLayout(1), label, canvas, button)

	w.SetContent(content)
	w.ShowAndRun()
}
