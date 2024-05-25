package main

import "testing"

func TestPawnMovement(t *testing.T) {
	
	board := initializeBoard()
	whiteTurn := true
	movePiece("e4", whiteTurn, board)
	if board['e' - 97][4].ID != 'P' {
		t.Errorf("Pawn did not move correctly.")
	}
	printBoard(board)

}
