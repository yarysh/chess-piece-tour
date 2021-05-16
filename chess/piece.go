package chess

import (
	"errors"
	"fmt"
)

type position struct {
	X, Y int
}

// Piece - chess piece
type Piece struct {
	Board     *Board
	MovesMade int
	MoveRules [][2]int
	position  position
}

// getPossibleMoves finds all possible moves for the piece based on that piece's move rules
func (p *Piece) getPossibleMoves(x, y int) []position {
	var moves []position
	for _, rule := range p.MoveRules {
		newX, newY := x+rule[0], y+rule[1]
		square, _ := p.Board.GetSquareByXY(newX, newY)
		if square != nil && !square.IsVisited() {
			moves = append(moves, position{X: newX, Y: newY})
		}
	}
	return moves
}

// SetPosition changes piece position on the chessboard
func (p *Piece) SetPosition(x, y int) error {
	if !p.Board.IsValidCoord(x, y) {
		return fmt.Errorf("position X=%d, Y=%d is out of board", x, y)
	}
	p.position.X = x
	p.position.Y = y
	return nil
}

// Move makes piece move and marks chessboard's square as visited
func (p *Piece) Move(x, y int) error {
	err := p.SetPosition(x, y)
	if err != nil {
		return err
	}
	square, err := p.Board.GetSquareByXY(x, y)
	if err != nil {
		return err
	}
	p.MovesMade++
	square.VisitedOnMove = p.MovesMade
	return nil
}

// Tour makes sequence of moves of a piece on a chessboard such that the piece visits every square exactly once
func (p *Piece) Tour() error {
	if p.Board.InUse {
		return errors.New("this board already in use")
	}
	if err := p.Move(p.position.X, p.position.Y); err != nil {
		return err
	}

	p.Board.InUse = true

	// Use Warnsdorff's rule
	for i := 1; i < p.Board.GetSize()*p.Board.GetSize(); i++ {
		possibleMoves := p.getPossibleMoves(p.position.X, p.position.Y)
		if len(possibleMoves) == 0 {
			return fmt.Errorf(
				"can't find next possible moves for position X=%d, Y=%d",
				p.position.X, p.position.Y,
			)
		}

		nextMove := possibleMoves[0]
		nextPossibleMoves := p.getPossibleMoves(nextMove.X, nextMove.Y)
		for j := 1; j < len(possibleMoves); j++ {
			move := possibleMoves[j]
			possibles := p.getPossibleMoves(move.X, move.Y)
			if len(possibles) > 0 && len(possibles) < len(nextPossibleMoves) {
				nextMove = move
				nextPossibleMoves = possibles
			}
		}

		if err := p.Move(nextMove.X, nextMove.Y); err != nil {
			return err
		}
	}

	return nil
}

// NewPiece returns new chess piece with specified chessboard and move rules
func NewPiece(board *Board, moveRules [][2]int) *Piece {
	return &Piece{
		Board:     board,
		MovesMade: 0,
		MoveRules: moveRules,
	}
}
