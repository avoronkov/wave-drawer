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

// Simple Bresenham implementation.
func (g *Grid) SetRange(y1, x1, y2, x2 int, val bool) {
	xMin, yMin, xMax, yMax := x1, y1, x2, y2
	if xMax == xMin {
		if abs(yMax) > abs(yMin) {
			g.Set(yMax, xMax, val)
		} else {
			g.Set(yMin, xMin, val)
		}
		return
	}
	if xMax < xMin {
		xMin, yMin, xMax, yMax = x2, y2, x1, y1
	}

	m_new := 2 * (yMax - yMin)
	ydiff := 1
	if m_new < 0 {
		ydiff = -1
		m_new = -m_new
	}
	// int slope_error_new = m_new - (x2 - x1);
	slope_error_new := 0
	for x, y := xMin, yMin; x <= xMax; x++ {
		g.Set(y, x, val)

		// Add slope to increment angle formed
		slope_error_new += m_new

		// Slope error reached limit, time to
		// increment y and update slope error.
		for slope_error_new > 0 {
			y += ydiff
			slope_error_new -= 2 * (x2 - x1)
		}
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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
