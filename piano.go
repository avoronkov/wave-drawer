package main

import (
	"bytes"
	"log"
	"runtime"
	"time"

	"github.com/avoronkov/waver/lib/midisynth/filters"
	"github.com/avoronkov/waver/lib/midisynth/instruments"
	"github.com/avoronkov/waver/lib/midisynth/player"
	"github.com/avoronkov/waver/lib/midisynth/wav"
	"github.com/avoronkov/waver/lib/midisynth/waves"
	"github.com/avoronkov/waver/lib/notes"

	oto "github.com/hajimehoshi/oto/v2"
)

type Piano struct {
	player *player.Player
	grid   *Grid

	scale notes.Scale

	context *oto.Context
}

func NewPiano(grid *Grid) (*Piano, error) {
	settings := wav.Default
	player := player.New(settings)
	// Init oto.Context
	c, ready, err := oto.NewContext(
		settings.SampleRate,
		settings.ChannelNum,
		settings.BitDepthInBytes,
	)
	if err != nil {
		return nil, err
	}
	<-ready

	return &Piano{
		player:  player,
		grid:    grid,
		scale:   notes.NewStandard(),
		context: c,
	}, nil
}

func (p *Piano) Play(octave int, note string) {
	n, ok := p.scale.Note(octave, note)
	if !ok {
		log.Printf("Unknown note: %v%v", note, octave)
		return
	}
	wave, err := p.Wave()
	if err != nil {
		log.Printf("Failed to create Wave: %v", err)
		return
	}
	data, done := p.player.PlayContext(wave, waves.NewNoteCtx(n.Freq, 0.33, 4.0, 0.0))
	pl := p.context.NewPlayer(data)
	pl.Play()
	<-done
	time.Sleep(1 * time.Second)
	runtime.KeepAlive(pl)
}

func (p *Piano) Wave() (waves.Wave, error) {
	buffer := &bytes.Buffer{}
	p.grid.Normalize()
	p.grid.Save(buffer)

	w, err := waves.ParseFormInput(buffer)
	if err != nil {
		return nil, err
	}
	var Adsr filters.AdsrFilter
	in := instruments.NewInstrument(w, Adsr.New())
	return in.Wave(), nil
}
