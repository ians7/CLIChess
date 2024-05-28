package main

import "testing"
import "fmt"

func TestBasicPawnMovement(t *testing.T) {
	board := initializeBoard()
	success := false
	whiteTurn := true
	passedAll := true
	board, success = movePiece("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("e5\n", whiteTurn, board)
	if success == true {
		t.Errorf("\u001b[31m" + "should not be able to move to e5")
	}
	board, success = movePiece("a3\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a3")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("f5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed f5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("a4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a4")
		passedAll = false
	}
	printBoard(board)
	if passedAll {
		fmt.Println("\u001b[32m" + "PASSED basic pawn movement.")
	}
}

func TestMultiplePossiblePawnCaptures(t *testing.T) {
	
	// f takes

	board := initializeBoard()
	success := false
	passedAll := true
	whiteTurn := true
	board, success = movePiece("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("d4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("a6\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a6")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("f4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed f4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("a5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("fxe5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed fxe5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	printBoard(board)
		
	// d takes
	board = initializeBoard()
	success = false
	whiteTurn = true
	board, success = movePiece("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("d4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("a6\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a6")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("f4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed f4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("a5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success = movePiece("dxe5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed dxe5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	printBoard(board)
	if passedAll {
		fmt.Println("\u001b[32m" + "PASSED selecting capturing piece." + "\u001b[37m")
	}
}
