package main

type Lagrange struct {
	Points []Point
}

func NewLagrange(points []Point) *Lagrange {
	return &Lagrange{
		Points: points,
	}
}

func (l *Lagrange) ljx(j int, x float64) float64 {
	res := 1.0
	for m, p := range l.Points {
		if m == j {
			continue
		}
		res *= (x - p.X) / (l.Points[j].X - p.X)
	}
	return res
}

func (l *Lagrange) Value(x float64) (value float64) {
	for j, p := range l.Points {
		value += p.Y * l.ljx(j, x)
	}
	return
}
