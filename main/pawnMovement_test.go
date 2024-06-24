package main

import "testing"
import "fmt"

func TestBasicPawnMovement(t *testing.T) {
	board := initializeBoard()
	isMate := false
	success := false
	whiteTurn := true
	passedAll := true
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e5\n", whiteTurn, board)
	if success == true {
		t.Errorf("\u001b[31m" + "should not be able to move to e5")
	}
	board, success, isMate = executeTurn("a3\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a3")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("f5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed f5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a4")
		passedAll = false
	}
	printBoard(board)
	if passedAll {
		fmt.Println("\u001b[32m" + "PASSED basic pawn movement.")
	}
	if isMate {
		fmt.Println("Mate!")
	}
}

func TestMultiplePossiblePawnCaptures(t *testing.T) {
	
	// f takes

	board := initializeBoard()
	isMate := false
	success := false
	passedAll := true
	whiteTurn := true
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a6\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a6")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("f4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed f4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("fxe5\n", whiteTurn, board)
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
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a6\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a6")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("f4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed f4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("dxe5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed dxe5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	printBoard(board)
	if passedAll {
		fmt.Println("\u001b[32m" + "PASSED selecting capturing piece." + "\u001b[37m")
	}
	if isMate {
		fmt.Println("Mate!")
	}
}

func TestEnPassantWhite(t *testing.T) {
	board := initializeBoard()
	isMate := false
	success := false
	whiteTurn := true
	passedAll := true
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a6\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a6")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("exd5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed exd5")
		passedAll = false
	}
	printBoard(board)

	board = initializeBoard()
	success = false
	whiteTurn = true
	passedAll = true
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d6\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d6")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("exd5\n", whiteTurn, board)
	if success == true {
		t.Errorf("\u001b[31m" + "exd5 should not have succeeded")
		passedAll = false
	}
	printBoard(board)
	if passedAll {
		fmt.Println("\u001b[32m" + "Passed all WHITE en passant tests" + "\u001b[37m")
	}
	if isMate {
		fmt.Println("Mate!")
	}
}

func TestEnPassantBlack(t *testing.T) {
	board := initializeBoard()
	success := false
	isMate := false
	whiteTurn := true
	passedAll := true
	board, success, isMate = executeTurn("a3\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a3")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("dxe4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed dxe4")
		passedAll = false
	}
	printBoard(board)

	board = initializeBoard()
	success = false
	whiteTurn = true

	board, success, isMate = executeTurn("e3\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e3")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d5\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d5")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("a3\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed a3")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("d4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed d4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("e4\n", whiteTurn, board)
	if success == false {
		t.Errorf("\u001b[31m" + "Failed e4")
		passedAll = false
	}
	whiteTurn = !whiteTurn
	board, success, isMate = executeTurn("dxe4\n", whiteTurn, board)
	if success == true {
		t.Errorf("\u001b[31m" + "exd4 should not have succeeded")
		passedAll = false
	}
	printBoard(board)
	if passedAll {
		fmt.Println("\u001b[32m" + "Passed all BLACK en passant tests" + "\u001b[37m")
	}
	if isMate {
		fmt.Println("Mate!")
	}
}
