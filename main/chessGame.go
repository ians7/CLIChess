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
	canBeEnPassant bool
	teamID         int
}

type Square struct {
	squareRow int16
	squareFile int16
}

var (
	G           = "\u001b[38;5;243m"
	W           = "\u001b[38;5;15m"
	B           = "\u001b[38;5;232m"
	bgRed       = "\u001b[;45m"
	bgCyan      = "\u001b[;46m"
	bgBlack     = "\u001b[;40m"
	Br          = "\u001b[38;5;94;m"
	blackKing   = Piece{'K', '\u2654', false, 1}
	blackQueen  = Piece{'Q', '\u2655', false, 1}
	blackRook   = Piece{'R', '\u2656', false, 1}
	blackKnight = Piece{'N', '\u2658', false, 1}
	blackBishop = Piece{'B', '\u2657', false, 1}
	blackPawn   = Piece{'P', '\u2659', false, 1}
	whiteKing   = Piece{'K', '\u2654', false, 0}
	whiteQueen  = Piece{'Q', '\u2655', false, 0}
	whiteRook   = Piece{'R', '\u2656', false, 0}
	whiteKnight = Piece{'N', '\u2658', false, 0}
	whiteBishop = Piece{'B', '\u2657', false, 0}
	whitePawn   = Piece{'P', '\u2659', false, 0}
	emptySquare = Piece{'0', ' ', false, 0}
)

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
		tempBoard, success := whiteMovement(board, row, file, capture, capturingPieceFile, pieceType, input)
		if wkInCheck, bkInCheck, isMate := detectCheckOnKing(tempBoard); wkInCheck || bkInCheck {
			if wkInCheck {
				fmt.Println("white king is in check")
				return prevBoard, false
			} else if bkInCheck {
				fmt.Println("black king is in check")
				if isMate {
					fmt.Println("White wins!")
				}
			}
		}
		tempBoard = removeWhiteEnPassant(tempBoard)
		return tempBoard, success
	} else {
		tempBoard, success := blackMovement(board, row, file, capture, capturingPieceFile, pieceType, input)
		fmt.Println(success)
		if wkInCheck, bkInCheck, isMate := detectCheckOnKing(tempBoard); wkInCheck || bkInCheck {
			if bkInCheck {
				fmt.Println("black king is in check")
				return prevBoard, false
			} else if wkInCheck {
				fmt.Println("white king is in check")
				if isMate {
					fmt.Println("Black wins!")
				}
			}
		}
		tempBoard = removeBlackEnPassant(tempBoard)
		return tempBoard, success
	}

	return board, false
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
	for !(k == destRow && l == destFile) && k > 0 && l > 0 && k < 8 && l < 8 {
		if board[k][l].pieceID != '0' && board[k][l] != board[pieceRow][pieceFile] {
			return true
		}
		k += int16(rowIterator)
		l += int16(fileIterator)
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
		spacesBetween = append(spacesBetween, Square{int16(k), int16(l)})
		k += rowIterator
		l += fileIterator
	}
	return spacesBetween

}

func getSpacesCanMove(pieceRow int16, pieceFile int16, board [8][8]Piece) []Square {
	var spacesCanMove []Square
	currPiece := board[pieceRow][pieceFile]
	switch currPiece.pieceID {
		case 'P':
			if currPiece.teamID == 0 {
				if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow + 2, pieceFile) && pieceRow == 1 {
					spacesCanMove = append(spacesCanMove, Square{pieceRow + 2, pieceFile})
					spacesCanMove = append(spacesCanMove, Square{pieceRow - 1, pieceFile})
				} else if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow + 1, pieceFile) {
					spacesCanMove = append(spacesCanMove, Square{pieceRow + 1, pieceFile})
				}
			} else {
				if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow - 2, pieceFile) && pieceRow == 6 {
					spacesCanMove = append(spacesCanMove, Square{pieceRow - 1, pieceFile})
					spacesCanMove = append(spacesCanMove, Square{pieceRow - 2, pieceFile})
				} else if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow - 1, pieceFile) {
					spacesCanMove = append(spacesCanMove, Square{pieceRow - 1, pieceFile})
				}
			}	

		case 'Q':
			for i := int16(0) ; i < 8 ; i++ {
				for j := int16(0) ; j < 8 ; j++ {
					fileDist := math.Abs(float64(pieceFile - j))
					rowDist := math.Abs(float64(pieceRow - i))
					if (fileDist == rowDist || fileDist == 0 || fileDist == 0) && !checkPieceInWay(board, pieceRow, pieceFile, i, j) {
						spacesCanMove = append(spacesCanMove, Square{i, j})
					}
				}
			}

		case 'R':
			for i := int16(0) ; i < 8 ; i++ {
				for j := int16(0) ; j < 8 ; j++ {
					fileDist := math.Abs(float64(pieceFile - j))
					rowDist := math.Abs(float64(pieceRow - i))
					if (fileDist == 0 || rowDist == 0) && !checkPieceInWay(board, pieceRow, pieceFile, i, j) {
						spacesCanMove = append(spacesCanMove, Square{i, j})
					}
				}
			}

		case 'N':
			for i := int16(0) ; i < 8 ; i++ {
				for j := int16(0) ; j < 8 ; j++ {
					fileDist := math.Abs(float64(pieceFile - j))
					rowDist := math.Abs(float64(pieceRow - i))
					if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
						spacesCanMove = append(spacesCanMove, Square{i, j})
					}
				}
			}
		
		case 'B':
			for i := int16(0) ; i < 8 ; i++ {
				for j := int16(0) ; j < 8 ; j++ {
					fileDist := math.Abs(float64(pieceFile - j))
					rowDist := math.Abs(float64(pieceRow - i))
					if (fileDist == rowDist) && !checkPieceInWay(board, pieceRow, pieceFile, i, j) {
						spacesCanMove = append(spacesCanMove, Square{i, j})
					}
				}
			}

		case 'K': 
			for i := int16(0) ; i < 8 ; i++ {
				for j := int16(0) ; j < 8 ; j++ {
					if currPiece.teamID == 0 {
						board, success := whiteMovement(board, i, j, false, -1, 'K', "")
						board[0][0] = emptySquare // do not remove
						if success {
							fmt.Println("row =", i, "file =", j)
							spacesCanMove = append(spacesCanMove, Square{i,j})
						}
					} else {
						board, success := blackMovement(board, i, j, false, -1, 'K', "")
						board[0][0] = emptySquare // do not remove
						if success {
							fmt.Println("row =", i, "file =", j)
							spacesCanMove = append(spacesCanMove, Square{i,j})
						}
					}
				}
			}
	}
	return spacesCanMove
}

func isMated(kingRow int, kingFile int, pieceFile int, pieceRow int, board [8][8]Piece) bool {
	var spacesBetween []Square = getSpacesBetween(board, kingRow, kingFile, pieceRow, pieceFile)
	var spacesCanMove []Square
	if len(getSpacesCanMove(int16(kingRow), int16(kingFile), board)) > 0 {
		fmt.Println("Spaces can move is non-empty")
		return false
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			spacesCanMove = getSpacesCanMove(int16(i), int16(j), board)
			for _, squareMove := range spacesBetween {
				for _, squareBetween := range spacesCanMove {
					if squareMove.squareRow == squareBetween.squareRow && squareMove.squareFile == squareBetween.squareFile {
						fmt.Println("piece can block")
						return false
					}
				}
			}
		}
	}
	return true
}

func detectCheckOnKing(board [8][8]Piece) (bool, bool, bool) {
	kingCount := 0
	wkRow := 0
	wkFile := 0
	bkRow := 0
	bkFile := 0
	checkingPieceRow := -1
	checkingPieceFile := -1
	wkInCheck := false
	bkInCheck := false
	isMate := false
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].pieceID == 'K' {
				if board[i][j].teamID == 1 {
					bkRow = i
					bkFile = j
				}
				if board[i][j].teamID == 0 {
					wkRow = i
					wkFile = j
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
			wkFileDist := math.Abs(float64(wkFile - j))
			wkRowDist := math.Abs(float64(wkRow - i))
			bkFileDist := math.Abs(float64(bkFile - j))
			bkRowDist := math.Abs(float64(bkRow - i))
			switch currPiece.pieceID {
			case 'P':
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist && wkFileDist == 1 && wkRowDist == 1) && !wkInCheck {
					fmt.Println("White king is in check by a pawn")
					wkInCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				} else if currPiece.teamID == 0 && (bkFileDist == 1 && bkRowDist == 1) && !bkInCheck {
					fmt.Println("Black king is in check by a pawn")
					fmt.Println(i, j)
					bkInCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				}

			case 'Q':
				if currPiece.teamID == 1 && (math.Abs(wkFileDist) == math.Abs(wkRowDist) || wkFileDist == 0 || wkRowDist == 0) && !wkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(wkRow), int16(wkFile)) {
						wkInCheck = false
					} else {
						fmt.Println("White king is in check by a queen")
						wkInCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (math.Abs(bkFileDist) == math.Abs(bkRowDist) || bkFileDist == 0 || bkRowDist == 0) && !bkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(bkRow), int16(bkFile)) {
						fmt.Println("Black king is NOT in check by a queen")
						bkInCheck = false
					} else {
						fmt.Println("Black king is in check by a queen")
						bkInCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			case 'R':
				if currPiece.teamID == 1 && (wkFileDist == 0 || wkRowDist == 0) && !wkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(wkRow), int16(wkFile)) {
						wkInCheck = false
					} else {
						fmt.Println("White king is in check by a rook")
						wkInCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (bkFileDist == 0 || bkRowDist == 0) && !bkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(bkFile), int16(bkFile)) {
						bkInCheck = false
					} else {
						bkInCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			case 'N':
				if currPiece.teamID == 1 && ((wkRowDist == 2 && wkFileDist == 1) || (wkRowDist == 1 && wkFileDist == 2)) && !wkInCheck {
					fmt.Println("White king is in check by a knight")
					wkInCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				} else if currPiece.teamID == 0 && ((bkRowDist == 2 && bkFileDist == 1) || (bkRowDist == 1 && bkFileDist == 2)) && !bkInCheck {
					fmt.Println("Black king is in check by a knight")
					bkInCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				}
			case 'B':
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist) && !wkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(wkRow), int16(wkFile)) {
						wkInCheck = false
					} else {
						fmt.Println("White king is in check by a bishop")
						wkInCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (bkFileDist == bkRowDist) && !bkInCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(bkFile), int16(bkFile)) {
						bkInCheck = false
					} else {
						fmt.Println("Black king is in check by a bishop")
						bkInCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			}
		}
	}
	if wkInCheck && isMated(wkRow, wkFile, checkingPieceRow, checkingPieceFile, board) {
		fmt.Println("Game over!")
		isMate = true
	} else if bkInCheck && isMated(bkRow, bkFile, checkingPieceRow, checkingPieceFile, board) {
		fmt.Println("Game over!")
		isMate = true
	}
	fmt.Println(wkInCheck, bkInCheck, isMate)
	return wkInCheck, bkInCheck, isMate
}

func removeWhiteEnPassant(board [8][8]Piece) [8][8]Piece {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].teamID == 0 && board[i][j].pieceID == 'P' {
				board[i][j].canBeEnPassant = false
			}
		}
	}
	return board
}

func removeBlackEnPassant(board [8][8]Piece) [8][8]Piece {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].teamID == 1 && board[i][j].pieceID == 'P' {
				board[i][j].canBeEnPassant = false
			}
		}
	}
	return board
}

func whiteMovement(board [8][8]Piece, row int16, file int16, capture bool, capturingPieceFile int, pieceType int, input string) ([8][8]Piece, bool) {
	if capture && board[row][file].teamID == 0 {
		return board, false
	}
	for i := int16(0); i < 8; i++ {
		for j := int16(0); j < 8; j++ {
			fileDisamb := false
			rowDisamb := false
			disambRow := -1
			disambFile := -1

			// Checking for disambiguation
			if match, err := regexp.MatchString(`^[a-hNKQBR][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				rowDisamb = true
				disambRow = int(input[1] - '0')
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]x?[a-h][1-8]\n$`, input); err == nil && match {
				disambFile = int(input[1] - 'a')
				fileDisamb = true
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				rowDisamb = true
				fileDisamb = true
				disambRow = int(input[1] - '0')
				disambFile = int(input[1] - 'a')
			}
			if rowDisamb && fileDisamb && i != int16(8-(input[2]-'0')) {
				break
			}
			rowDist := math.Abs(float64(row - i))
			fileDist := math.Abs(float64(file - j))
			if board[i][j].teamID == 0 && board[i][j].pieceID == pieceType {
				switch pieceType {
				case 'P':
					spacesCanMove := 1
					if i == 6 {
						spacesCanMove = 2
					}
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i-1) && board[row][file].teamID == 1 {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							if row == 0 {
								switch pawnPromote() {
								case "Q\n":
									board[row][file] = whiteQueen
								case "B\n":
									board[row][file] = whiteBishop
								case "N\n":
									board[row][file] = whiteKnight
								case "R\n":
									board[row][file] = whiteRook
								}
							}
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].canBeEnPassant {
							board[row-1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						}
					}
					if i-row > 0 && i-row <= int16(spacesCanMove) && file == j && board[row][file].pieceID == '0' && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						if rowDist == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						if row == 0 {
							switch pawnPromote() {
							case "Q\n":
								board[row][file] = whiteQueen
							case "B\n":
								board[row][file] = whiteBishop
							case "N\n":
								board[row][file] = whiteKnight
							case "R\n":
								board[row][file] = whiteRook
							}
						}
						return board, true
					}

				case 'K':
					if checkPieceInWay(board, i, j, row, file) == true || rowDist > 1 || fileDist > 1 {
						return board, false
					}

					if (rowDist == fileDist || file-j == 0 || row-i == 0) && board[row][file].teamID != 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}

				case 'Q':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (!fileDisamb && !rowDisamb) {
						if rowDist == fileDist || file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}

				case 'R':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
						if file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}

				case 'N':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
						if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}
				case 'B':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
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
			disambRow := -1
			disambFile := -1
			// Checking for disambiguation
			if match, err := regexp.MatchString(`^[a-hNKQBR][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				rowDisamb = true
				disambRow = int(input[1] - '0')
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]x?[a-h][1-8]\n$`, input); err == nil && match {
				fileDisamb = true
				disambFile = int(input[1] - 'a')
			} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
				disambRow = int(input[1] - '0')
				disambFile = int(input[1] - 'a')
				if board[input[1]-'a'][input[2]-'0'].pieceID != pieceType {
					break
				}
			}
			rowDist := math.Abs(float64(row - i))
			fileDist := math.Abs(float64(file - j))
			if board[i][j].teamID == 1 && int16(board[i][j].pieceID) == int16(pieceType) {
				switch pieceType {
				case 'P':
					spacesCanMove := 1
					if i == 1 {
						spacesCanMove = 2
					}
					if capture && j == int16(capturingPieceFile) {
						if ((file == j+1 || file == j-1) && row == i+1) && board[row][file].teamID == 0 {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							if row == 7 {
								switch pawnPromote() {
								case "Q\n":
									board[row][file] = whiteQueen
								case "B\n":
									board[row][file] = whiteBishop
								case "N\n":
									board[row][file] = whiteKnight
								case "R\n":
									board[row][file] = whiteRook
								}
							}
							return board, true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].canBeEnPassant {
							board[row+1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							return board, true
						}
					}
					if row-i > 0 && row-i <= int16(spacesCanMove) && file == j && board[row][file].pieceID == '0' && capturingPieceFile == 0 {
						board[i][j].canBeEnPassant = false
						if rowDist == 2 {
							board[i][j].canBeEnPassant = true
						}
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						spacesCanMove = 1
						if row == 7 {
							switch pawnPromote() {
							case "Q\n":
								board[row][file] = whiteQueen
							case "B\n":
								board[row][file] = whiteBishop
							case "N\n":
								board[row][file] = whiteKnight
							case "R\n":
								board[row][file] = whiteRook
							}
						}
						fmt.Println("succesfully moved pawn")
						return board, true
					}

				case 'K':
					if checkPieceInWay(board, i, j, row, file) == true || rowDist > 1 || fileDist > 1 {
						return board, false
					}

					if (rowDist == fileDist || file-j == 0 || row-i == 0) && board[row][file].teamID != 1{
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						return board, true
					}

				case 'Q':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
						if rowDist == fileDist || file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}

							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}

				case 'R':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
						if file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) == true {
								return board, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}

				case 'N':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
						if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							return board, true
						}
					}

				case 'B':
					if (rowDisamb && i == int16(disambRow)) || (fileDisamb && j == int16(disambFile)) || (fileDisamb == rowDisamb) {
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
