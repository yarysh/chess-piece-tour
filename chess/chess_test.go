package chess

import (
	"testing"
)

func TestBoardInitSquares(t *testing.T) {
	boardSize := 10
	board := NewBoard(boardSize)
	squareCnt := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board.squares[i][j].VisitedOnMove != 0 {
				t.Errorf("square x=%d, y=%d should contain 0, got %d", i, j, board.squares[i][j])
				return
			}
			squareCnt++
		}
	}
	if squareCnt != boardSize*boardSize {
		t.Errorf("board shoud have %d squares, got %d", boardSize*boardSize, squareCnt)
		return
	}
}

func TestIsValidCoord(t *testing.T) {
	boardSize := 10
	board := NewBoard(boardSize)
	coords := []struct {
		X, Y  int
		Valid bool
	}{
		{X: 0, Y: 0, Valid: true},
		{X: boardSize - 1, Y: boardSize - 1, Valid: true},
		{X: boardSize, Y: boardSize, Valid: false},
		{X: -1, Y: -1, Valid: false},
	}

	for _, coord := range coords {
		if coord.Valid != board.IsValidCoord(coord.X, coord.Y) {
			want := "valid"
			if !coord.Valid {
				want = "not " + want
			}
			t.Errorf("cords X=%d, Y=%d should be %s", coord.X, coord.Y, want)
			return
		}
	}
}

func TestXYtoRowCol(t *testing.T) {
	boardSize := 10
	board := NewBoard(boardSize)
	data := [][2][2]int{
		{{0, 0}, {9, 0}},
		{{9, 0}, {9, 9}},
		{{0, 9}, {0, 0}},
		{{9, 9}, {0, 9}},
	}
	for _, item := range data {
		row, col := board.XYtoRowCol(item[0][0], item[0][1])
		if row != item[1][0] || col != item[1][1] {
			t.Errorf("cords X=%d, Y=%d should be converted to row=%d, col=%d", item[0][0], item[0][1], item[1][0], item[1][1])
			return
		}
	}
}

func TestTour(t *testing.T) {
	boardSize := 10
	board := NewBoard(boardSize)
	pawn := NewPiece(board, [][2]int{
		{0, 3}, {2, 2}, {3, 0}, {2, -2},
		{0, -3}, {-2, -2}, {-3, 0}, {-2, 2},
	})
	pawn.SetPosition(9, 9)

	err := pawn.Tour()
	if err != nil {
		t.Errorf("got error: %s", err)
		return
	}

	moves := map[int][2]int{}
	gotSum := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			move, exists := moves[board.squares[i][j].VisitedOnMove]
			if exists {
				t.Errorf("move #%d is already existed at the square x=%d, y=%d", board.squares[i][j], move[0], move[1])
				return
			}
			moves[board.squares[i][j].VisitedOnMove] = [2]int{i, j}
			gotSum += board.squares[i][j].VisitedOnMove
		}
	}
	// arithmetic progression
	wantSum := int((float32(1+boardSize*boardSize) / 2.0) * float32(boardSize*boardSize))
	if wantSum != gotSum {
		t.Errorf("sum of all moves want %d, got %d", wantSum, gotSum)
		return
	}
}
