package main

import (
	"fmt"
)

type Square struct {
	squareRow  int
	squareFile int
}


var (
	G           = "\u001b[38;5;243m"
	W           = "\u001b[38;5;15m"
	B           = "\u001b[38;5;232m"
	bgRed       = "\u001b[;45m"
	bgCyan      = "\u001b[;46m"
	bgBlack     = "\u001b[;40m"
	Br          = "\u001b[38;5;94;m"
)

func initializeBoard() [8][8]Piece {

	board := [8][8]Piece{
		{blackRook, blackKnight, blackBishop, blackQueen, blackKing, blackBishop, blackKnight, blackRook},
		{blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn},
		{emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare},
		{emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare},
		{emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare},
		{emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare, emptySquare},
		{whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn},
		{whiteRook, whiteKnight, whiteBishop, whiteQueen, whiteKing, whiteBishop, whiteKnight, whiteRook},
	}
	return board
}

func printBoard(board [8][8]Piece) {
	fmt.Printf(W + "  a   b   c   d   e   f   g   h\n")
	fmt.Printf(Br + " -------------------------------\n")
	bgColor := bgRed
	colorBool := true
	pieceColor := B
	for i := 0; i < 8; i++ {
		fmt.Printf(Br + "|")
		for j := 0; j < 8; j++ {
			if board[i][j].teamID == 1 {
				pieceColor = B
			} else if board[i][j].teamID == 0 {
				pieceColor = W
			}
			if colorBool {
				bgColor = bgRed
			} else {
				bgColor = bgCyan
			}
			fmt.Printf(bgColor+pieceColor+" %c "+Br+bgBlack+"|", board[i][j].icon)
			colorBool = !colorBool
		}
	fmt.Printf(W+" %d", 8-i)
		fmt.Printf(Br + "\n -------------------------------\n" + W)
		colorBool = !colorBool
	}
}

func checkPieceInWay(board [8][8]Piece, pieceRow int, pieceFile int, destRow int, destFile int) bool {
	rowIterator := 1
	fileIterator := 1
	if pieceRow > destRow {
		rowIterator = -1
	} else if pieceRow == destRow {
		rowIterator = 0
	}
	if pieceFile > destFile {
		fileIterator = -1
	} else if pieceFile == destFile {
		fileIterator = 0
	}

	k, l := pieceRow, pieceFile
	k += rowIterator
	l += fileIterator
	for !(k == destRow && l == destFile) && k > -1 && l > -1 && k < 8 && l < 8 {
		if board[k][l].pieceID != '0' {
			return true
		}
		k += rowIterator
		l += fileIterator
	}
	return false
}

func getSpacesBetween(board [8][8]Piece, kingRow int, kingFile int, pieceRow int, pieceFile int) []Square {
	var spacesBetween []Square
	rowIterator := 1
	fileIterator := 1
	if kingRow > pieceRow {
		rowIterator = -1
	} else if kingRow == pieceRow {
		rowIterator = 0
	}
	if kingFile > pieceRow {
		fileIterator = -1
	} else if kingFile == pieceRow {
		fileIterator = 0
	}

	k, l := kingRow, kingFile
	k += rowIterator
	l += fileIterator
	for !(k == pieceRow && l == pieceFile) && k > 0 && l > 0 && k < 8 && l < 8 {
		spacesBetween = append(spacesBetween, Square{k, l})
		k += rowIterator
		l += fileIterator
	}
	return spacesBetween
}
