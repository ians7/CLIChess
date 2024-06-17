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

var (
	G           = "\u001b[38;5;243m"
	W           = "\u001b[38;5;15m"
	B           = "\u001b[38;5;232m"
	bgRed       = "\u001b[;45m"
	bgCyan      = "\u001b[;46m"
	bgBlack     = "\u001b[;40m"
	Br          = "\u001b[38;5;94;m"
	blackKing   = Piece{'K', '\u2654', B, 1, false, 1}
	blackQueen  = Piece{'Q', '\u2655', B, 7, false, 1}
	blackRook   = Piece{'R', '\u2656', B, 7, false, 1}
	blackKnight = Piece{'N', '\u2658', B, 3, false, 1}
	blackBishop = Piece{'B', '\u2657', B, 7, false, 1}
	blackPawn   = Piece{'P', '\u2659', B, 2, false, 1}
	whiteKing   = Piece{'K', '\u2654', W, 1, false, 0}
	whiteQueen  = Piece{'Q', '\u2655', W, 7, false, 0}
	whiteRook   = Piece{'R', '\u2656', W, 7, false, 0}
	whiteKnight = Piece{'N', '\u2658', W, 3, false, 0}
	whiteBishop = Piece{'B', '\u2657', W, 7, false, 0}
	whitePawn   = Piece{'P', '\u2659', W, 2, false, 0}
	emptySquare = Piece{'0', ' ', B, 0, false, 0}
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
	if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?x[a-h][1-8]\n$`, input); err == nil && match {
		*capture = true
		*capturingPieceFile = int(input[2] - 'a')
		if input[0] >= 97 && input[0] <= 104 {
			*pieceType = 'P'
			*capturingPieceFile = int(input[0] - 'a')
		} else {
			*pieceType = int(input[0])
		}
		return true
	} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?[a-h][1-8]\n$`, input); err == nil && match {
		if input[0] >= 97 && input[0] <= 104 {
			*pieceType = 'P'
			*capturingPieceFile = int(input[0] - 'a')
		} else {
			*pieceType = int(input[0])
		}
		return true
	} else if match, err := regexp.MatchString(`^[NKQBR][a-h][1-8]\n$`, input); err == nil && match {
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

func pawnPromote() string {
	for {
		fmt.Println("What will you promote to?(Q, N, B, R)")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read user input. Aborting.")
		}
		if match, err := regexp.MatchString(`^[QNBR]\n$`, input); err == nil && match {
			return input
		} else {
			fmt.Println("Improper input.")
		}
	}

}

func checkPieceInWay(board [8][8]Piece, pieceRow int16, pieceFile int16, destRow int16, destFile int16) bool {
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
	k += int16(rowIterator)
	l += int16(fileIterator)
	for !(k == destRow && l == destFile) && k < 8 && l < 8 {
		if board[k][l].pieceID != '0' {
			return true
		}
		k += int16(rowIterator)
		l += int16(fileIterator)
	}
	return false
}

func detectCheckOnKing(board [8][8]Piece) (bool, bool) {
	kingCount := 0
	whiteKingRow := 0
	whiteKingFile := 0
	blackKingRow := 0
	blackKingFile := 0
	wkInCheck := false
	bkInCheck := false
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].pieceID == 'K' {
				if board[i][j].teamID == 1 {
					blackKingRow = i
					blackKingFile = j
				}
				if board[i][j].teamID == 0 {
					whiteKingRow = i
					whiteKingFile = j
				}
				kingCount++
			}
			if kingCount == 2 {
				break
			}
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			currPiece := board[i][j]
			wkFileDist := math.Abs(float64(whiteKingFile - j))
			wkRowDist := math.Abs(float64(whiteKingRow - i))
			bkFileDist := math.Abs(float64(blackKingFile - j))
			bkRowDist := math.Abs(float64(blackKingRow - i))
			if currPiece.pieceID == 'P' {
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist && wkFileDist == 1 && wkRowDist == 1) && !wkInCheck {
					fmt.Println("White king is in check by a pawn")
					wkInCheck = true
				} else if currPiece.teamID == 0 && (bkFileDist == bkRowDist && bkFileDist == 1 && bkRowDist == 1) && !bkInCheck {
					fmt.Println("Black king is in check by a pawn")
					bkInCheck = true
				}
			} else if currPiece.pieceID == 'Q' {
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist || wkFileDist == 0 || wkRowDist == 0) && !wkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(whiteKingRow), int16(whiteKingFile)) {
						wkInCheck = false
					} else {
						fmt.Println("White king is in check by a queen")
						wkInCheck = true
					}
				} else if currPiece.teamID == 0 && (math.Abs(bkFileDist) == math.Abs(bkRowDist) || bkFileDist == 0 || bkRowDist == 0) && !bkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(blackKingRow), int16(blackKingFile)) && !bkInCheck {
						fmt.Println("Black king is NOT in check by a queen")
						bkInCheck = false
					} else {
						fmt.Println("Black king is in check by a queen")
						bkInCheck = true
					}
				}
			} else if currPiece.pieceID == 'R' {
				if currPiece.teamID == 1 && (wkFileDist == 0 || wkRowDist == 0) && !wkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(whiteKingRow), int16(whiteKingFile)) {
						wkInCheck = false
					} else {
						fmt.Println("White king is in check by a rook")
						wkInCheck = true
					}
				} else if currPiece.teamID == 0 && (bkFileDist == 0 || bkRowDist == 0) && !bkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(blackKingRow), int16(blackKingFile)) {
						bkInCheck = false
					} else {
						bkInCheck = true
					}
				}
			} else if currPiece.pieceID == 'N' {
				if currPiece.teamID == 1 && ((wkRowDist == 2 && wkFileDist == 1) || (wkRowDist == 1 && wkFileDist == 2)) && !wkInCheck {
					fmt.Println("White king is in check by a knight")
					wkInCheck = true
				} else if currPiece.teamID == 0 && ((bkRowDist == 2 && bkFileDist == 1) || (bkRowDist == 1 && bkFileDist == 2)) && !bkInCheck {
					fmt.Println("Black king is in check by a knight")
					bkInCheck = true
				}
			} else if currPiece.pieceID == 'B' {
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist) && !wkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(whiteKingRow), int16(whiteKingFile)) {
						wkInCheck = false
					} else {
						fmt.Println("White king is in check by a bishop")
						wkInCheck = true
					}
				} else if currPiece.teamID == 0 && (bkFileDist == bkRowDist) && !bkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(blackKingRow), int16(blackKingFile)) {
						bkInCheck = false
					} else {
						fmt.Println("Black king is in check by a bishop")
						bkInCheck = true
					}
				}
			}
		}
	}
	fmt.Println(wkInCheck, bkInCheck)
	return wkInCheck, bkInCheck
}
func whiteMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int, input string) ([8][8]Piece, bool) {
	if capture && board[row][file].teamID == 0 {
		return board, false
	}
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			fileDisamb := false
			rowDisamb := false

			// Checking for disambiguation
			if match, err := regexp.MatchString(`^[a-hNKQBR][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				rowDisamb = true
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]x?[a-h][1-8]\n$`, input); err == nil && match {
				fileDisamb = true
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				rowDisamb = true
				fileDisamb = true
			}
			if rowDisamb && fileDisamb && i != int16(8-(input[2]-'0')) {
				fmt.Println("i =", i, "disambiguated row =", 8-(input[2]-'0'))
				break
			}
			rowDist := math.Abs(float64(row - i))
			fileDist := math.Abs(float64(file - j))
			if board[i][j].teamID == 0 && board[i][j].pieceID == pieceType {
				if pieceType == 'P' {
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i-1) && board[row][file].teamID == 1 {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							if row == 0 {
								promoteTo := pawnPromote()
								if promoteTo == "Q\n" {
									board[row][file] = whiteQueen
								} else if promoteTo == "B\n" {
									board[row][file] = whiteBishop
								} else if promoteTo == "N\n" {
									board[row][file] = whiteKnight
								} else if promoteTo == "R\n" {
									board[row][file] = whiteRook
								}
							}
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].pieceID == 'P' && board[row][file].canBeEnPassant {
							board[row-1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						}
					}
					if i-row > 0 && i-row <= board[i][j].spacesCanMove && file == j && board[row][file].pieceID == '0' && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						if rowDist == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						board[row][file].spacesCanMove = 1
						if row == 0 {
							fmt.Println("Row is equal to 0")
							promoteTo := pawnPromote()
							if promoteTo == "Q\n" {
								board[row][file] = whiteQueen
							} else if promoteTo == "B\n" {
								board[row][file] = whiteBishop
							} else if promoteTo == "N\n" {
								board[row][file] = whiteKnight
							} else if promoteTo == "R\n" {
								board[row][file] = whiteRook
							}
						}
						return board, true
					}
				} else if pieceType == 'K' {
					if checkPieceInWay(board, i, j, row, file) == true || rowDist > 1 || fileDist > 1 {
						return board, false
					}

					if rowDist == fileDist || file-j == 0 || row-i == 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}

				} else if pieceType == 'Q' {

					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (!fileDisamb && !rowDisamb) {
						if rowDist == fileDist || file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				} else if pieceType == 'R' {

					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				} else if pieceType == 'N' {
					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				} else if pieceType == 'B' {

					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if (i%2 == j%2 && row%2 == file%2) || (i%2 != j%2 && row%2 != file%2) {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}

							if rowDist == fileDist {
								board[row][file] = board[i][j]
								board[i][j] = emptySquare
								return board, true
							}
						}
					}
				}
			}
		}
	}
	return board, false
}

func blackMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int, input string) ([8][8]Piece, bool) {
	if capture && board[row][file].teamID == 1 {
		return board, false
	}
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			fileDisamb := false
			rowDisamb := false

			// Checking for disambiguation
			if match, err := regexp.MatchString(`^[a-hNKQBR][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				rowDisamb = true
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]x?[a-h][1-8]\n$`, input); err == nil && match {
				fileDisamb = true
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				if board[input[1]-'a'][input[2]-'0'].pieceID != pieceType {
					break
				}
			}
			rowDist := math.Abs(float64(row - i))
			fileDist := math.Abs(float64(file - j))
			if board[i][j].teamID == 1 && int16(board[i][j].pieceID) == int16(pieceType) {
				if pieceType == 'P' {
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i+1) && board[row][file].teamID == 0 {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							if row == 7 {
								promoteTo := pawnPromote()
								if promoteTo == "Q\n" {
									board[row][file] = blackQueen
								} else if promoteTo == "B\n" {
									board[row][file] = blackBishop
								} else if promoteTo == "N\n" {
									board[row][file] = blackKnight
								} else if promoteTo == "R\n" {
									board[row][file] = blackRook
								}
							}
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
						if rowDist == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						board[row][file].spacesCanMove = 1
						if row == 7 {
							promoteTo := pawnPromote()
							if promoteTo == "Q\n" {
								board[row][file] = blackQueen
							} else if promoteTo == "B\n" {
								board[row][file] = blackBishop
							} else if promoteTo == "N\n" {
								board[row][file] = blackKnight
							} else if promoteTo == "R\n" {
								board[row][file] = blackRook
							}
						}
						return board, true
					}
				} else if pieceType == 'K' {
					if checkPieceInWay(board, i, j, row, file) == true || rowDist > 1 || fileDist > 1 {
						return board, false
					}

					if rowDist == fileDist || file-j == 0 || row-i == 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}
				} else if pieceType == 'Q' {

					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if rowDist == fileDist || file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}

							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				} else if pieceType == 'R' {
					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				} else if pieceType == 'N' {

					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				} else if pieceType == 'B' {

					if (rowDisamb && i == int16(input[1]-'0')) || (fileDisamb && j == int16(input[1]-'a')) || (fileDisamb == rowDisamb) {
						if (i%2 == j%2 && row%2 == file%2) || (i%2 != j%2 && row%2 != file%2) {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}

							if rowDist == fileDist {
								board[row][file] = board[i][j]
								board[i][j] = emptySquare
								return board, true
							}
						}
					}
				}
			}
		}
	}
	return board, false
}

func executeTurn(input string, whiteTurn bool, board [8][8]Piece) ([8][8]Piece, bool) {
	prevBoard := board
	pieceType := 0
	capture := false
	capturingPieceFile := 0
	if !parseInput(input, &pieceType, &capture, &capturingPieceFile) {
		return board, false
	}
	row := int16(8 - (input[len(input)-2] - '0'))
	file := int16(input[len(input)-3] - 'a')

	if row > 8 || row < 0 || file > 8 || file < 0 {
		return board, false
	}

	if capture == false && board[row][file].pieceID != '0' {
		return board, false
	}

	if whiteTurn {
		board, success := whiteMovement(board, row, file, capture, capturingPieceFile, pieceType, input)
		if wkInCheck, bkInCheck := detectCheckOnKing(board); wkInCheck || bkInCheck {
			if wkInCheck {
				fmt.Println("white king is in check")
				return prevBoard, false
			} else if bkInCheck {
				fmt.Println("black king is in check")
				return board, success
			}
		}
		return board, success
	} else {
		board, success := blackMovement(board, row, file, capture, capturingPieceFile, pieceType, input)
		if wkInCheck, bkInCheck := detectCheckOnKing(board); wkInCheck || bkInCheck {
			if bkInCheck {
				fmt.Println("black king is in check")
				return prevBoard, false
			} else if wkInCheck {
				fmt.Println("white king is in check")
				return board, success
			}
		}
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
