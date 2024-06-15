package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type Piece struct {
	pieceID        int
	icon           int
	color          string
	spacesCanMove  int16
	canBeEnPassant bool
	teamID         int
}

var G = "\u001b[38;5;243m"
var W = "\u001b[38;5;15m"
var B = "\u001b[38;5;232m"
var bgRed = "\u001b[;45m"
var bgCyan = "\u001b[;46m"
var bgBlack = "\u001b[;40m"
var Br = "\u001b[38;5;94;m"
var blackKing = Piece{'K', '\u2654', B, 1, false, 1}
var blackQueen = Piece{'Q', '\u2655', B, 7, false, 1}
var blackRook = Piece{'R', '\u2656', B, 7, false, 1}
var blackKnight = Piece{'N', '\u2658', B, 3, false, 1}
var blackBishop = Piece{'B', '\u2657', B, 7, false, 1}
var blackPawn = Piece{'P', '\u2659', B, 2, false, 1}
var whiteKing = Piece{'K', '\u2654', W, 1, false, 0}
var whiteQueen = Piece{'Q', '\u2655', W, 7, false, 0}
var whiteRook = Piece{'R', '\u2656', W, 7, false, 0}
var whiteKnight = Piece{'N', '\u2658', W, 3, false, 0}
var whiteBishop = Piece{'B', '\u2657', W, 7, false, 0}
var whitePawn = Piece{'P', '\u2659', W, 2, false, 0}
var emptySquare = Piece{'0', ' ', B, 0, false, 0}

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
	for i := 0; i < 8; i++ {
		fmt.Printf(Br + "|")
		for j := 0; j < 8; j++ {
			if colorBool {
				bgColor = bgRed
			} else {
				bgColor = bgCyan
			}
			fmt.Printf(bgColor+board[i][j].color+" %c "+Br+bgBlack+"|", board[i][j].icon)
			colorBool = !colorBool
		}
		fmt.Printf(W+" %d", 8-i)
		fmt.Printf(Br + "\n -------------------------------\n" + W)
		colorBool = !colorBool
	}
}

func parseInput(input string, pieceType *int, capture *bool, capturingPieceFile *int) bool {
	if match, err := regexp.MatchString(`^([a-hNKQBR])[a-h]x?[a-h][1-8]\n$`, input); err == nil && match {
		fmt.Println("have not implemeted disambiguation")
		return false
	} else if match, err := regexp.MatchString(`^([a-hNKQBR])[1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
		fmt.Println("have not implemeted disambiguation")
		return false
	} else if match, err := regexp.MatchString(`^([a-hNKQBR])[1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
		fmt.Println("have not implemeted disambiguation")
		return false
	} else if match, err := regexp.MatchString(`^([a-hNKQBR])[a-h][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
		fmt.Println("have not implemeted double disambiguation")
		return false
	}
	if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?x[a-h][1-8]\n$`, input); err == nil && match {
		*capture = true
		*capturingPieceFile = int(input[2] - '0' - 49)
		if input[0] >= 97 && input[0] <= 104 {
			*pieceType = 'P'
			*capturingPieceFile = int(input[0] - '0' - 49)
		} else {
			*pieceType = int(input[0])
		}
		return true
	} else if match, err := regexp.MatchString(`^[NKQBR][a-h][1-8]`, input); err == nil && match {
		*pieceType = int(input[0])
		return true
	} else if match, err := regexp.MatchString(`^[a-h][1-8]\n$`, input); err == nil && match {
		*pieceType = 'P'
		return true
	} else {
		fmt.Println("Please input valid chess notation (in parse input)", input)
		fmt.Println(match, err)
		return false
	}
}

func checkPieceInWay(board [8][8]Piece, pieceRow int16, pieceFile int16, destRow int16, destFile int16) bool {
	rowIterator := 1
	fileIterator := 1
	fmt.Println("pieceFile =", pieceFile, "destFile =", destFile)
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

	fmt.Println("rowIterator =", rowIterator, "fileIterator =", fileIterator)
	fmt.Println("row =", pieceRow, "file =", pieceFile)
	fmt.Println("destRow =", destRow, "destFile =", destFile)
	k, l := pieceRow, pieceFile
	for !(k == destRow && l == destFile) {
		if board[k][l].pieceID != '0' && board[k][l] != board[pieceRow][pieceFile] {
			fmt.Println("There is a piece in the way")
			return true
		}
		k += int16(rowIterator)
		l += int16(fileIterator)
	}
	return false
}

func detectCheckOnKing(board [8][8]Piece) Piece {
	fmt.Println("Have not implemented checks.")
	return emptySquare
}

func whiteMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int) ([8][8]Piece, bool) {

	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			if board[i][j].teamID == 0 && int16(board[i][j].pieceID) == int16(pieceType) {
				if pieceType == 'P' {
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i-1) && board[row][file].pieceID == 'P' {
							fmt.Println("normal capture conditions are true")
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].pieceID == 'P' && board[row][file].canBeEnPassant {
							fmt.Println("en passant conditions are true")
							board[row-1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						}
					}
					if i-row > 0 && i-row <= board[i][j].spacesCanMove && file == j && board[row][file].pieceID == '0' && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						if math.Abs(float64(i-row)) == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						board[row][file].spacesCanMove = 1
						return board, true
					}
				} else if pieceType == 'K' {
					if checkPieceInWay(board, i, j, row, file) == true || math.Abs(float64(row-i)) > 1 || math.Abs(float64(j-file)) > 1 {
						return board, false
					}

					if math.Abs(float64(row-i)) == math.Abs(float64(j-file)) || file-j == 0 || row-i == 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}

				} else if pieceType == 'Q' {
					if checkPieceInWay(board, i, j, row, file) == true {
						return board, false
					}

					if math.Abs(float64(row-i)) == math.Abs(float64(j-file)) || file-j == 0 || row-i == 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'R' {

					if file-j == 0 || row-i == 0 {
						if checkPieceInWay(board, i, j, row, file) == true {
							return board, false
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'N' {

					if (math.Abs(float64(row-i)) == 2 && math.Abs(float64(file-j)) == 1) || (math.Abs(float64(row-i)) == 1 && math.Abs(float64(file-j)) == 2) {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'B' {

					if (i%2 == j%2 && row%2 == file%2) || (i%2 != j%2 && row%2 != file%2) {
						if checkPieceInWay(board, i, j, row, file) == true {
							return board, false
						}

						if math.Abs(float64(row-i)) == math.Abs(float64(j-file)) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				}
			}
		}
	}
	return board, false
}

func blackMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int) ([8][8]Piece, bool) {
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			if board[i][j].teamID == 1 && int16(board[i][j].pieceID) == int16(pieceType) {
				if pieceType == 'P' {
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i+1) && board[row][file].pieceID == 'P' {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].pieceID == 'P' && board[row][file].canBeEnPassant {
							board[row+1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						}
					}
					if row-i > 0 && row-i <= board[i][j].spacesCanMove && file == j && board[row][file].pieceID == '0' && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						fmt.Println("canBeEnPassant ->", math.Abs(float64(row-i)))
						if math.Abs(float64(row-i)) == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						board[row][file].spacesCanMove = 1
						return board, true
					}
				} else if pieceType == 'K' {
					if checkPieceInWay(board, i, j, row, file) == true || math.Abs(float64(row-i)) > 1 || math.Abs(float64(j-file)) > 1 {
						return board, false
					}

					if math.Abs(float64(row-i)) == math.Abs(float64(j-file)) || file-j == 0 || row-i == 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'Q' {

					if checkPieceInWay(board, i, j, row, file) == true {
						return board, false
					}

					if math.Abs(float64(row-i)) == math.Abs(float64(j-file)) || file-j == 0 || row-i == 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'R' {
					if file-j == 0 || row-i == 0 {
						if checkPieceInWay(board, i, j, row, file) == true {
							return board, false
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'N' {

					if (math.Abs(float64(row-i)) == 2 && math.Abs(float64(file-j)) == 1) || (math.Abs(float64(row-i)) == 1 && math.Abs(float64(file-j)) == 2) {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'B' {

					if (i%2 == j%2 && row%2 == file%2) || (i%2 != j%2 && row%2 != file%2) {
						if checkPieceInWay(board, i, j, row, file) == true {
							return board, false
						}

						if math.Abs(float64(row-i)) == math.Abs(float64(j-file)) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				}
			}
		}
	}
	return board, false
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

	if row > 8 || row < 0 || file > 8 || file < 0 {
		return board, false
	}

	if capture == false && board[row][file].pieceID != '0' {
		return board, false
	}

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
			fmt.Println("Please input a valid move.")
		}
	}

}
