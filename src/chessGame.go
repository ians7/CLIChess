package main

import (
    "fmt"
    "os"
    "bufio"
//    "strings"
)

type Piece struct {
    identity int
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
		    fmt.Printf(board[i][j].color+" %s"+Br+" |", string(board[i][j].icon))
		}
		fmt.Printf(W + " %d", 8 - i)
   		fmt.Printf(Br + "\n -------------------------------\n")
    }
}

func verifyInput(board [8][8]Piece, input string) bool {
    move := input[:len(input) - 1]
    if move[0] >= 97 && move[0] <= 104 {
		fmt.Println(move + " is a pawn move")
		return true
    }	


    return false
}

func movePiece(board [8][8]Piece, input string) {
	if !verifyInput(board, input) {
		fmt.Println("This move was invalid")
	}

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
		movePiece(board, input)
    }

}
