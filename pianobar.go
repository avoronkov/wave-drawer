package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type octaveWrapper struct {
	N int
}

func (w *octaveWrapper) Inc() {
	if w.N < 7 {
		w.N++
	}
}

func (w *octaveWrapper) Dec() {
	if w.N > 2 {
		w.N--
	}
}

func initPianoButtons(piano *Piano, octave *octaveWrapper) (buttons []fyne.CanvasObject) {
	for _, n := range []string{"C", "D", "E", "F", "G", "A", "B"} {
		title := fmt.Sprintf("%v", n)
		nt := n
		buttons = append(buttons, widget.NewButton(title, func() {
			go piano.Play(octave.N, nt)
		}))
	}
	return
}

func NewPianoBar(piano *Piano) *fyne.Container {
	octave := &octaveWrapper{N: 2}
	buttons := initPianoButtons(piano, octave)

	label := widget.NewLabel(fmt.Sprintf("%v", octave.N))
	allButtons := append([]fyne.CanvasObject{
		widget.NewButton("-1", func() {
			octave.Dec()
			label.SetText(fmt.Sprintf("%v", octave.N))
		}),
		label,
		widget.NewButton("+1", func() {
			octave.Inc()
			label.SetText(fmt.Sprintf("%v", octave.N))
		}),
	}, buttons...)
	return container.NewVBox(
		allButtons...,
	)
}
