package chess

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"
)

// square - one square on a chessboard
type square struct {
	VisitedOnMove int
}

// IsVisited returns true if square has been visited
func (s *square) IsVisited() bool {
	return s.VisitedOnMove != 0
}

// Board - chessboard
type Board struct {
	InUse   bool
	squares [][]*square
}

// GetSize returns the number of squares on a side of the chessboard
func (b *Board) GetSize() int {
	return len(b.squares)
}

// XYtoRowCol converts X, Y coordinates to the row and column numbers for internal square storage
func (b *Board) XYtoRowCol(x, y int) (int, int) {
	return (len(b.squares) - 1) - y, x
}

// IsValidCoord checks that square with specified coordinates exists
func (b *Board) IsValidCoord(x, y int) bool {
	row, col := b.XYtoRowCol(x, y)
	return row >= 0 && len(b.squares)-1 >= row && col >= 0 && len(b.squares[row])-1 >= col
}

// GetSquareByXY returns square by provided X, Y coordinates
func (b *Board) GetSquareByXY(x, y int) (*square, error) {
	if !b.IsValidCoord(x, y) {
		return nil, fmt.Errorf("invalid coordinate: X=%d, Y=%d", x, y)
	}
	row, col := b.XYtoRowCol(x, y)
	return b.squares[row][col], nil
}

func (b *Board) String() string {
	buffer := new(bytes.Buffer)
	maxMove := b.GetSize() * b.GetSize()
	writer := tabwriter.NewWriter(buffer, len(strconv.Itoa(maxMove)), 0, 0, ' ', tabwriter.Debug)
	for _, row := range b.squares {
		var line []string
		for _, square := range row {
			line = append(line, strconv.Itoa(square.VisitedOnMove))
		}
		fmt.Fprintf(writer, "|%s\t\n", strings.Join(line, "\t"))
	}
	writer.Flush()
	return strings.Trim(buffer.String(), "\n")
}

// NewBoard returns new chessboard
func NewBoard(size int) *Board {
	squares := make([][]*square, size)
	for row := range squares {
		squares[row] = make([]*square, size)
		for col := 0; col < size; col++ {
			squares[row][col] = &square{}
		}
	}
	return &Board{
		squares: squares,
	}
}
