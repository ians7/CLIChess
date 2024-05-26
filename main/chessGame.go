package main

import (
    "fmt"
    "os"
    "bufio"
	"os/exec"
	"strings"
)

type Piece struct {
    ID int
    icon int
    color string
    spacesCanMove int16
}

var G = "\u001b[38;5;243m"
var W = "\u001b[38;5;15m"
var B = "\u001b[38;5;232m"
var Br = "\u001b[38;5;94m"

var blackKing = Piece{'K', '\u2654', G, 1}
var blackQueen = Piece{'Q', '\u2655', G, 7}
var blackRook = Piece{'R', '\u2656', G, 7}
var blackKnight = Piece{'N', '\u2658', G, 3}
var blackBishop = Piece{'B', '\u2657', G, 7}
var blackPawn = Piece{'P', '\u2659', G, 2}
var whiteKing = Piece{'K', '\u2654', W, 1}
var whiteQueen = Piece{'Q', '\u2658', W, 7}
var whiteRook = Piece{'R', '\u2656', W, 7}
var whiteKnight = Piece{'N', '\u2658', W, 3}
var whiteBishop = Piece{'B', '\u2657', W, 7}
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
		    fmt.Printf(board[i][j].color+" %c"+Br+" |", board[i][j].icon)
		}
		fmt.Printf(W + " %d", 8 - i)
   		fmt.Printf(Br + "\n -------------------------------\n")
    }
}

func movePiece(input string, whiteTurn bool, board [8][8]Piece) ([8][8]Piece, bool) {
	var pieceType byte 
	capture := false
	if strings.Contains(input, "x") {
		capture = true
	}
	var capturingPieceFile int16 = 0
	if input[0] >= 97 && input[0] <= 104 {
		if capture {
			capturingPieceFile = int16(input[0] - '0' - 48)
		}
		pieceType = 'P'
	} else if input[0] == 'N' || input[0] == 'K' || input[0] == 'Q' || input[0] == 'B' || input[0] == 'R' {
		pieceType = input[0]
	} else {
		fmt.Println("Please input valid chess notation")
		return board, false;
	}

	row := int16(8 - (input[len(input) - 2] - '0'))
	file := int16((input[len(input) - 3] - '0') - 49)
	fmt.Println(file, row)
	color := "" 
	if whiteTurn {
		color = W
	} else {
		color = G
	}

	for i := int16(0) ; i < 8; i++ {
		for j := int16(0) ; j < 8; j++ {
			if board[i][j].color == color && int16(board[i][j].ID) == int16(pieceType) {
				if pieceType == 'P' {
					if whiteTurn {
						if capture && j == capturingPieceFile { 
							if ((file == j + 1 && row == i - 1) || (file == j - 1 && row == i - 1)) && board[row][file].color == G {
								fmt.Println(j, capturingPieceFile)
								board[row][file] = board[i][j]
								board[i][j] = emptySquare
								return board, true
							} else {
								return board, false
							}
						} 
						if i - row <= board[i][j].spacesCanMove && file == j && board[row][file].color == B && capturingPieceFile == 0 {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file].spacesCanMove = 1
							return board, true
						}
					} else {
						if capture {
							if (file == j + 1 && row == i + 1) || (file == j - 1 && row == i + 1) && board[row][file].color == W {
								board[row][file] = board[i][j]
								board[i][j] = emptySquare
								return board, true
							} else {
								return board, false
							}
						} 
						if row - i <= board[i][j].spacesCanMove && file == j {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file].spacesCanMove = 1
							return board, true
						}
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
		board, success = movePiece(input, whiteTurn, board)
		if success {
			whiteTurn = !whiteTurn
		}
		exec.Command("clear")
    }

}
