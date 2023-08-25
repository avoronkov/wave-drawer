package main

import (
	"fmt"
	"io"
)

type Grid struct {
	data   [][]bool
	w, h   int
	middle int

	bezierDone bool

	OnUpdate func()
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

func (g *Grid) Clear() {
	defer g.onUpdate()
	g.bezierDone = false
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			g.data[y][x] = false
		}
	}
}

func (g *Grid) Get(y, x int) bool {
	if y >= 0 && y < g.h && x >= 0 && x < g.w {
		return g.data[y][x]
	}
	return false
}

func (g *Grid) Set(y, x int, val bool) {
	defer g.OnUpdate()
	if y >= 0 && y < g.h && x >= 0 && x < g.w {
		g.data[y][x] = val
	}
}

func (g *Grid) set(y, x int, val bool) {
	if y >= 0 && y < g.h && x >= 0 && x < g.w {
		g.data[y][x] = val
	}
}

// Simple Bresenham implementation.
func (g *Grid) SetRange(y1, x1, y2, x2 int, val bool) {
	defer g.onUpdate()
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

	slope_error_new := 0
	for x, y := xMin, yMin; x <= xMax; x++ {
		g.set(y, x, val)

		// Add slope to increment angle formed
		slope_error_new += m_new

		// Slope error reached limit, time to
		// increment y and update slope error.
		for slope_error_new > 0 {
			y += ydiff
			slope_error_new -= 2 * (xMax - xMin)
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

	if (x%100 == 0 && y%4 == 0) || ((y-g.middle)%100 == 0 && x%4 == 0) {
		return PLine
	}
	return PEmpty
}

func (g *Grid) Normalize() {
	defer g.onUpdate()
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

	// Trim left tail
	for x := 0; x < g.w; x++ {
		y, ok := g.getPoint(x)
		if !ok {
			continue
		}
		if y > g.middle {
			g.cleanColumn(x)
		} else {
			break
		}
	}

	// Trim right tail
	for x := g.w - 1; x >= 0; x-- {
		y, ok := g.getPoint(x)
		if !ok {
			continue
		}
		if y < g.middle {
			g.cleanColumn(x)
		} else {
			break
		}
	}
}

func (g *Grid) Bezier() {
	if g.bezierDone {
		return
	}
	defer g.onUpdate()
	points := g.points()

	for t := 0.0; t <= 1.0; t += 0.001 {
		p := Bezzier(points, t)
		g.set(int(p.Y+0.5), int(p.X+0.5), true)
	}
	g.bezierDone = true
}

func (g *Grid) Lagrange() {
	if g.bezierDone {
		return
	}
	defer g.onUpdate()
	points := g.points()
	lg := NewLagrange(points)
	for x := 0; x < g.w; x++ {
		y := lg.Value(float64(x))
		g.set(int(y+0.5), x, true)
	}
	g.bezierDone = true
}

func (g *Grid) points() (points []Point) {
	for x := 0; x < g.w; x++ {
		y, ok := g.getPoint(x)
		if !ok {
			continue
		}
		points = append(points, Point{float64(x), float64(y)})
	}
	return
}

func (g *Grid) Save(out io.Writer) {
	startX := 0
	startXSet := false
	for x := 0; x < g.w; x++ {
		y, ok := g.getPoint(x)
		if !ok {
			continue
		}
		if !startXSet {
			startXSet = true
			startX = x
		}

		vx := x - startX
		vy := g.middle - y
		fmt.Fprintf(out, "%v %v\n", vx, vy)
	}
}

func (g *Grid) getPoint(x int) (y int, ok bool) {
	for y := 0; y < g.h; y++ {
		if g.data[y][x] {
			return y, true
		}
	}
	return 0, false
}

func (g *Grid) cleanColumn(x int) {
	for y := 0; y < g.h; y++ {
		g.data[y][x] = false
	}
}

func (g *Grid) onUpdate() {
	if g.OnUpdate != nil {
		g.OnUpdate()
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
