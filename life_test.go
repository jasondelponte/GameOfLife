package main

import (
	"testing"
)

func TestStep(t *testing.T) {
	seed := make([]Point, 5)
	seed[0] = Point{3, 3}
	seed[1] = Point{4, 6}
	seed[2] = Point{5, 6}
	seed[3] = Point{6, 6}
	seed[4] = Point{10, 14}

	life := NewLife(15, 15, 0, seed)

	life.Step()

	if life.board.CellState(Point{4, 6}) != false {
		t.Error("Invlid state for 4,6 after step")
	}
	if life.board.CellState(Point{5, 6}) != true {
		t.Error("Invlid state for 5,6 after step")
	}
	if life.board.CellState(Point{6, 6}) != false {
		t.Error("Invlid state for 6,6 after step")
	}

	if life.board.CellState(Point{5, 5}) != true {
		t.Error("Invlid state for 5,5 after step")
	}
	if life.board.CellState(Point{5, 7}) != true {
		t.Error("Invlid state for 5,7 after step")
	}

}
