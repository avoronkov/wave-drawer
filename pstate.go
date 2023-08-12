package main

type PState int

const (
	PEmpty PState = iota
	PPoint
	PMainLine
	PLine
)
