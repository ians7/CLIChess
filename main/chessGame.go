package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Piece struct {
	ID             int
	icon           int
	color          string
	spacesCanMove  int16
	canBeEnPassant bool
}

var G = "\u001b[38;5;243m"
var W = "\u001b[38;5;15m"
var B = "\u001b[38;5;232m"
var Br = "\u001b[38;5;94m"
var blackKing = Piece{'K', '\u2654', G, 1, false}
var blackQueen = Piece{'Q', '\u2655', G, 7, false}
var blackRook = Piece{'R', '\u2656', G, 7, false}
var blackKnight = Piece{'N', '\u2658', G, 3, false}
var blackBishop = Piece{'B', '\u2657', G, 7, false}
var blackPawn = Piece{'P', '\u2659', G, 2, false}
var whiteKing = Piece{'K', '\u2654', W, 1, false}
var whiteQueen = Piece{'Q', '\u2658', W, 7, false}
var whiteRook = Piece{'R', '\u2656', W, 7, false}
var whiteKnight = Piece{'N', '\u2658', W, 3, false}
var whiteBishop = Piece{'B', '\u2657', W, 7, false}
var whitePawn = Piece{'P', '\u2659', W, 2, false}
var emptySquare = Piece{'0', ' ', B, 0, false}

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
	for i := 0; i < 8; i++ {
		fmt.Printf(Br + "|")
		for j := 0; j < 8; j++ {
			fmt.Printf(board[i][j].color+" %c"+Br+" |", board[i][j].icon)
		}
		fmt.Printf(W+" %d", 8-i)
		fmt.Printf(Br + "\n -------------------------------\n")
	}
}

func parseInput(input string, pieceType *int, capture *bool, capturingPieceFile *int) bool {
	if match, err := regexp.MatchString(`([a-hNKQBR])?([a-h1-8])x[a-h][1-8]\n`, input); err == nil && match {
		*capture = true
		*capturingPieceFile = int(input[2] - '0' - 49)
		if input[0] >= 97 && input[0] <= 104 {
			*pieceType = 'P'
			*capturingPieceFile = int(input[0] - '0' - 49)
		} else {
			*pieceType = int(input[0])
		}
		return true
	} else if match, err := regexp.MatchString(`[a-h][1-8]\n`, input); err == nil && match {
		*pieceType = 'P'
		*capture = false
		return true
	} else {
		fmt.Println("Please input valid chess notation", input)
		fmt.Println(match, err)
		return false
	}
}

func whiteMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int) ([8][8]Piece, bool) {
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			if board[i][j].color == W && int16(board[i][j].ID) == int16(pieceType) {
				if pieceType == 'P' {
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i-1) && board[row][file].color == G {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].color == G && board[file][row].canBeEnPassant {
							board[row-1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						} else {
							return board, false
						}
					}
					if i - row > 0 && i - row <= board[i][j].spacesCanMove && file == j && board[row][file].color == B && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						if i-row == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						board[row][file].spacesCanMove = 1
						return board, true
					}
				}
			if pieceType == 'K' {

			}
			if pieceType == 'Q' {

			}
			if pieceType == 'R' {

			}
			if pieceType == 'N' {

			}
			if pieceType == 'B' {

			}
			}
		}
	}
	return board, false
}

func blackMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int) ([8][8]Piece, bool) {
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			if board[i][j].color == G && int16(board[i][j].ID) == int16(pieceType) {
				if pieceType == 'P' {
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i+1) && board[row][file].color == W {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].color == W && board[file][row].canBeEnPassant {
							board[row+1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						} else {
							return board, false
						}
					}
					if row - i > 0 && row-i <= board[i][j].spacesCanMove && file == j && board[row][file].color == B && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						if row-i == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						board[row][file].spacesCanMove = 1
						return board, true
					}
				}
			if pieceType == 'K' {

			}
			if pieceType == 'Q' {

			}
			if pieceType == 'R' {

			}
			if pieceType == 'N' {

			}
			if pieceType == 'B' {

			}
			}
		}
	}
	return board, false;
}

func executeTurn(input string, whiteTurn bool, board [8][8]Piece) ([8][8]Piece, bool) {
	pieceType := 0
	capture := false
	capturingPieceFile := 0
	if !parseInput(input, &pieceType, &capture, &capturingPieceFile) {
		return board, false
	}
	row := int16(8 - (input[len(input)-2] - '0'))
	file := int16((input[len(input)-3] - '0') - 49)

	if whiteTurn {
		board, success := whiteMovement(board, row, file, capture, capturingPieceFile, pieceType)
		return board, success
	} else {
		board, success := blackMovement(board, row, file, capture, capturingPieceFile, pieceType)
		return board, success
	}

	return board, false
}

func main() {
	board := initializeBoard()
	fmt.Println("Enter proper chess notation to make a move.")
	whiteTurn := true
	for {
		printBoard(board)

		if whiteTurn {
			fmt.Printf("White move: ")
		} else {
			fmt.Printf("Black move: ")
		}

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read user input. Aborting.")
		}
		success := false
		board, success = executeTurn(input, whiteTurn, board)
		if success {
			whiteTurn = !whiteTurn
		} else {
			fmt.Println("Please input valid chess notation.")
		}
	}

}
