package main

import (
	"fmt"
)

type Life struct {
	board *Board
	round int
}

func NewLife(height, width int, chance int, seed []Point) *Life {
	board := NewBoard(height, width)

	if seed != nil {
		board.ApplySeed(seed)
	} else {
		board.RandomeSeed(chance)
	}

	return &Life{board: board}
}

func (l *Life) Step() {
	for i := 0; i < l.board.Height; i++ {
		for j := 0; j < l.board.Width; j++ {
			p := Point{i, j}
			n := l.board.Neighbours(p)

			if l.board.CellState(p) == true && (n <= 1 || n >= 4) {
				l.board.ClearCell(p)
			} else if n == 3 {
				l.board.SetCell(p)
			}
		}
	}

	l.board.SwapLayers()
	l.round++
}

func (l *Life) String() string {
	str := fmt.Sprintln("Round:", l.round)
	str += fmt.Sprint(l.board)
	return str
}
