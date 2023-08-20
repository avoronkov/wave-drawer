package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestGridSetRange(t *testing.T) {
	t.Run("Forward x movement", func(t *testing.T) {
		is := is.New(t)

		g := NewGrid(4, 4)
		g.SetRange(0, 0, 3, 3, true)

		exp := [][]bool{
			{true, false, false, false},
			{false, true, false, false},
			{false, false, true, false},
			{false, false, false, true},
		}

		is.Equal(g.data, exp)
	})

	t.Run("Idle x", func(t *testing.T) {
		is := is.New(t)

		g := NewGrid(4, 4)
		g.SetRange(2, 0, 3, 0, true)

		exp := [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{true, false, false, false},
		}

		is.Equal(g.data, exp)
	})

	t.Run("Backward x movement", func(t *testing.T) {
		is := is.New(t)

		g := NewGrid(4, 4)
		g.SetRange(3, 3, 0, 0, true)

		exp := [][]bool{
			{true, false, false, false},
			{false, true, false, false},
			{false, false, true, false},
			{false, false, false, true},
		}

		is.Equal(g.data, exp)
	})
}
