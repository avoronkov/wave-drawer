package main

type Grid struct {
	data [][]bool
	w, h int
}

func NewGrid(h, w int) *Grid {
	data := make([][]bool, h)
	for y := 0; y < h; y++ {
		data[y] = make([]bool, w)
	}
	return &Grid{
		data: data,
		w:    w,
		h:    h,
	}
}

func (g *Grid) Get(y, x int) bool {
	if y >= 0 && y < g.h && x >= 0 && x < g.w {
		return g.data[y][x]
	}
	return false
}

func (g *Grid) Set(y, x int, val bool) {
	if y >= 0 && y < g.h && x >= 0 && x < g.w {
		g.data[y][x] = val
	}
}
