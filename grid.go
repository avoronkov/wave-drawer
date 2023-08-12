package main

type Grid struct {
	data   [][]bool
	w, h   int
	middle int
}

func NewGrid(h, w int) *Grid {
	g := &Grid{}
	g.Resize(h, w)
	return g
}

func (g *Grid) Resize(h, w int) {
	newData := make([][]bool, h)
	for y := 0; y < h; y++ {
		newData[y] = make([]bool, w)
	}

	for y := 0; y < h && y < g.h; y++ {
		for x := 0; x < w && x < g.w; x++ {
			newData[y][x] = g.data[y][x]
		}
	}
	g.data = newData
	g.h = h
	g.w = w
	g.middle = g.h / 2
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

func (g *Grid) GetState(y, x int) PState {
	if g.Get(y, x) {
		return PPoint
	}
	if y == g.middle {
		return PMainLine
	}

	if x%100 == 0 || (y-g.middle)%100 == 0 {
		return PLine
	}
	return PEmpty
}

func (g *Grid) Normalize() {
	for x := 0; x < g.w; x++ {
		// upper point
		upperPoint := false
		for y := 0; y <= g.middle; y++ {
			if g.data[y][x] {
				g.cleanColumn(x)
				g.data[y][x] = true
				upperPoint = true
				break
			}
		}
		if upperPoint {
			continue
		}
		// lower point
		for y := g.h - 1; y >= g.middle; y-- {
			if g.data[y][x] {
				g.cleanColumn(x)
				g.data[y][x] = true
				break
			}
		}
	}
}

func (g *Grid) cleanColumn(x int) {
	for y := 0; y < g.h; y++ {
		g.data[y][x] = false
	}
}
