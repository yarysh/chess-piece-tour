## Chess piece tour

A program that makes sequence of moves of a chess piece on a chessboard such that the piece visits every square exactly once.
You can create your own piece with specified move rules. It's possible to start from any position.
Based on Warnsdorff's rule.

## Examples

**Pawn's tour with special started position**
```go
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
```

**Knight's tour**
```go
package main

import (
	"fmt"
	"github.com/yarysh/chess-piece-tour/chess"
)

func main() {
	board := chess.NewBoard(10)
	knight := chess.NewPiece(board, [][2]int{
		{2, 1}, {1, 2}, {2, -1}, {1, -2},
		{-2, 1}, {-1, 2}, {-2, -1}, {-1, -2},
	})
	err := knight.Tour()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(board)
}
```
