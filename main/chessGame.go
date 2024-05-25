package main

import (
    "fmt"
    "os"
    "bufio"
//    "strings"
)

type Piece struct {
    ID int
    icon int
    color string
    spacesCanMove int
}

var G = "\u001b[38;5;243m"
var W = "\u001b[38;5;15m"
var B = "\u001b[38;5;232m"
var Br = "\u001b[38;5;94m"

var blackKing = Piece{'K', '\u2654', G, 1}
var blackQueen = Piece{'Q', '\u2655', G, 8}
var blackRook = Piece{'R', '\u2656', G, 8}
var blackKnight = Piece{'N', '\u2658', G, 3}
var blackBishop = Piece{'B', '\u2657', G, 8}
var blackPawn = Piece{'P', '\u2659', G, 2}
var whiteKing = Piece{'K', '\u2654', W, 1}
var whiteQueen = Piece{'Q', '\u2658', W, 8}
var whiteRook = Piece{'R', '\u2656', W, 8}
var whiteKnight = Piece{'N', '\u2658', W, 3}
var whiteBishop = Piece{'B', '\u2657', W, 8}
var whitePawn = Piece{'P', '\u2659', W, 2}
var emptySquare = Piece{'0', ' ', B, 0}

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
		    fmt.Printf(board[i][j].color+" U+%X"+Br+" |", board[i][j].icon)
		}
		fmt.Printf(W + " %d", 8 - i)
   		fmt.Printf(Br + "\n -------------------------------\n")
    }
}

func parseLocationToMove(input string) (x byte, y byte) {
	location := input
	if input[0] >= 97 && input[0] <= 104 {
		location = input[0:]
	}

	return location[0] - 97, location[1] - 49
}

func movePiece(input string, whiteTurn bool, board [8][8]Piece) [8][8]Piece {
	var pieceType byte 
	if input[0] >= 97 && input[0] <= 104 {
		pieceType = 'P'
	} else if input[0] == 'N' || input[0] == 'K' || input[0] == 'Q' || input[0] == 'B' || input[0] == 'R' {
		pieceType = input[0]
	} else {
		fmt.Println("Please input valid chess notation")
		return board
	}
	fmt.Println("Piece Type", pieceType)

	xCor, yCor := parseLocationToMove(input)

	if whiteTurn {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if board[i][j].color == W && byte(board[i][j].ID) == pieceType {
					
					if pieceType == 'P' {
						if yCor - byte(board[i][j].spacesCanMove) >= 0 {
							fmt.Printf("Pawn Move\n")
							board[xCor][yCor] = board[i][j]
							board[i][j] = emptySquare
							return board
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
	}

	return board
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
		board = movePiece(input, whiteTurn, board)
		whiteTurn = !whiteTurn
    }

}
