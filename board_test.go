package main

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	board := NewBoard(30, 40)

	if len(board.curr) != 30 {
		t.Error("Board curr column size not correct, expected 30, got,", len(board.curr))
	}
	if (len(board.curr[0])) != 40 {
		t.Error("Board curr row[0] size not correct, expected 30, got,", len(board.curr))
	}
	if (len(board.curr[29])) != 40 {
		t.Error("Board curr row[29] size not correct, expected 30, got,", len(board.curr))
	}
}

func TestSeedBoard(t *testing.T) {
	board := NewBoard(15, 15)
	seed := make([]Point, 5)
	seed[0] = Point{3, 3}
	seed[1] = Point{4, 6}
	seed[2] = Point{5, 6}
	seed[3] = Point{6, 6}
	seed[4] = Point{10, 14}
	board.ApplySeed(seed)

	if board.curr[3][3] != 1 || board.curr[4][6] != 1 || board.curr[5][6] != 1 || board.curr[6][6] != 1 || board.curr[10][14] != 1 {
		t.Error("Failed to seed board", board.curr)
	}

	if board.next[3][3] != 1 || board.next[4][6] != 1 || board.next[5][6] != 1 || board.next[6][6] != 1 || board.next[10][14] != 1 {
		t.Error("Failed to seed board", board.next)
	}
}

func TestNeighbours(t *testing.T) {
	board := NewBoard(15, 15)
	seed := make([]Point, 3)
	seed[0] = Point{4, 6}
	seed[1] = Point{5, 6}
	seed[2] = Point{6, 6}
	board.ApplySeed(seed)

	if count := board.Neighbours(Point{4, 6}); count != 1 {
		t.Error("Failed to count neighbours. Expected 1, got, ", count)
	}
	if count := board.Neighbours(Point{5, 6}); count != 2 {
		t.Error("Failed to count neighbours. Expected 2, got, ", count)
	}
	if count := board.Neighbours(Point{6, 6}); count != 1 {
		t.Error("Failed to count neighbours. Expected 1, got, ", count)
	}
	if count := board.Neighbours(Point{10, 14}); count != 0 {
		t.Error("Failed to count neighbours. Expected 0, got, ", count)
	}

}

func TestCellState(t *testing.T) {
	board := NewBoard(15, 15)
	seed := make([]Point, 3)
	seed[0] = Point{4, 6}
	seed[1] = Point{5, 6}
	seed[2] = Point{6, 6}
	board.ApplySeed(seed)

	if s := board.CellState(Point{4, 6}); s != true {
		t.Error("Failed to get cell state. Expected 1, got, ", s)
	}
	if s := board.CellState(Point{14, 0}); s != false {
		t.Error("Failed to get cell state. Expected 1, got, ", s)
	}
}

func TestSwapBoardAreas(t *testing.T) {
	board := NewBoard(15, 15)
	seed := make([]Point, 3)
	seed[0] = Point{4, 6}
	seed[1] = Point{5, 6}
	seed[2] = Point{6, 6}
	board.ApplySeed(seed)

	board.SetCell(Point{10, 10})
	if board.curr[10][10] != 0 || board.next[10][10] != 1 {
		t.Error("Setting cell effected the current layer not next")
	}
	board.ClearCell(Point{5, 6})
	if board.curr[5][6] != 1 || board.next[5][6] != 0 {
		t.Error("Clearing cell effected the current layer not next")
	}

	board.SwapLayers()
	if board.curr[10][10] != 1 || board.curr[10][10] != 1 {
		t.Error("Swap did not update set cells correctly")
	}
	board.ClearCell(Point{5, 6})
	if board.curr[5][6] != 0 || board.next[5][6] != 0 {
		t.Error("Swap did not update cleared cells correctly")
	}
}
