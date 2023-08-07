package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type draggedCanvasObject struct {
	fyne.CanvasObject
	onDragged func(obj fyne.CanvasObject, e *fyne.DragEvent)
}

func MakeDraggableCanvasObject(obj fyne.CanvasObject, fn func(obj fyne.CanvasObject, e *fyne.DragEvent)) *draggedCanvasObject {
	return &draggedCanvasObject{
		CanvasObject: obj,
		onDragged:    fn,
	}
}

var _ fyne.Draggable = (*draggedCanvasObject)(nil)

func (d *draggedCanvasObject) Dragged(e *fyne.DragEvent) {
	d.onDragged(d.CanvasObject, e)
}

func (d *draggedCanvasObject) DragEnd() {

}

func (d *draggedCanvasObject) CreateRenderer() fyne.WidgetRenderer {
	log.Printf("CreateRenderer(): %v", d.CanvasObject.Size())
	return widget.NewSimpleRenderer(d.CanvasObject)
}
