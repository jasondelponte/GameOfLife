package main

import (
	"math/rand"
	"time"
)

type Point struct {
	X, Y int
}

type Board struct {
	Height, Width int
	curr          [][]int
	next          [][]int
}

func NewBoard(height, width int) *Board {
	curr := make([][]int, height)
	next := make([][]int, height)
	for i := range curr {
		curr[i] = make([]int, width)
		next[i] = make([]int, width)
	}

	return &Board{
		Height: height,
		Width:  width,
		curr:   curr,
		next:   next,
	}
}

func (b *Board) ApplySeed(seed []Point) {
	for i := range seed {
		p := seed[i]
		b.curr[p.X][p.Y] = 1
		b.next[p.X][p.Y] = 1
	}
}

func (b *Board) RandomeSeed(chance int) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			v := rand.Intn(100)

			if v < chance {
				b.curr[i][j] = 1
				b.next[i][j] = 1
			}
		}
	}
}

func (b *Board) Neighbours(p Point) int {
	count := 0
	x := p.X - 1
	y := p.Y - 1
	for i := 0; i < 3; i++ {
		xo := b.offsetForEdge(x+i, b.Height)

		for j := 0; j < 3; j++ {
			yo := b.offsetForEdge(y+j, b.Width)

			if i == 1 && j == 1 { // Skip self
				continue
			}

			count += b.curr[xo][yo]
		}
	}

	return count
}

func (b *Board) offsetForEdge(i, max int) int {
	if i < 0 {
		i = max - 1
	} else if i == max {
		i = 0
	}
	return i
}

func (b *Board) CellState(p Point) bool {
	return b.curr[p.X][p.Y] == 1
}

func (b *Board) SetCell(p Point) {
	b.next[p.X][p.Y] = 1
}

func (b *Board) ClearCell(p Point) {
	b.next[p.X][p.Y] = 0
}

func (b *Board) SwapLayers() {
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			b.curr[i][j] = b.next[i][j]
		}
	}
}

func (b *Board) String() string {
	str := ""
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			if b.curr[i][j] == 1 {
				str += "X"
			} else {
				str += "."
			}
		}
		str += "\n"
	}

	return str
}
