package main

import (
	"fmt"
	"github.com/yarysh/chess-piece-tour/chess"
)

func main() {
	board := chess.NewBoard(10)
	pawn := chess.NewPiece(board, [][2]int{
		{0, 3}, {2, 2}, {3, 0}, {2, -2},
		{0, -3}, {-2, -2}, {-3, 0}, {-2, 2},
	})
	pawn.SetPosition(7, 2)
	err := pawn.Tour()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(board)
}
