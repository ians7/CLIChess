package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	"regexp"
)

type Piece struct {
	pieceID        int
	icon           int
	canBeEnPassant bool
	teamID         int
}

var (
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
	emptySquare = Piece{'0', ' ', false, -1}
)

func pawnPromote() string {
	for {
		fmt.Println("What will you promote to?(Q, N, B, R)")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed. Aborting.")
		}
		if match, err := regexp.MatchString(`^[QNBR]\n$`, input); err == nil && match {
			return input
		} else {
			fmt.Println("Improper input.")
		}
	}
}

func getSpacesCanMove(pieceRow int, pieceFile int, board [8][8]Piece) []Square {
	var spacesCanMove []Square
	currPiece := board[pieceRow][pieceFile]
	switch currPiece.pieceID {
	case 'P':
		if currPiece.teamID == 0 {
			if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow-2, pieceFile) && pieceRow == 1 {
				spacesCanMove = append(spacesCanMove, Square{pieceRow - 2, pieceFile})
				spacesCanMove = append(spacesCanMove, Square{pieceRow - 1, pieceFile})
			} else if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow-1, pieceFile) {
				spacesCanMove = append(spacesCanMove, Square{pieceRow - 1, pieceFile})
			}
		} else {
			if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow+2, pieceFile) && pieceRow == 6 {
				spacesCanMove = append(spacesCanMove, Square{pieceRow + 2, pieceFile})
				spacesCanMove = append(spacesCanMove, Square{pieceRow + 1, pieceFile})
			} else if !checkPieceInWay(board, pieceRow, pieceFile, pieceRow+1, pieceFile) {
				spacesCanMove = append(spacesCanMove, Square{pieceRow + 1, pieceFile})
			}
		}

	case 'Q':
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				fileDist := math.Abs(float64(pieceFile - j))
				rowDist := math.Abs(float64(pieceRow - i))
				if (fileDist == rowDist || fileDist == 0 || rowDist == 0) && !checkPieceInWay(board, pieceRow, pieceFile, i, j) {
					spacesCanMove = append(spacesCanMove, Square{i, j})
				}
			}
		}

	case 'R':
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				fileDist := math.Abs(float64(pieceFile - j))
				rowDist := math.Abs(float64(pieceRow - i))
				if (fileDist == 0 || rowDist == 0) && !checkPieceInWay(board, pieceRow, pieceFile, i, j) {
					spacesCanMove = append(spacesCanMove, Square{i, j})
				}
			}
		}

	case 'N':
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				fileDist := math.Abs(float64(pieceFile - j))
				rowDist := math.Abs(float64(pieceRow - i))
				if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
					spacesCanMove = append(spacesCanMove, Square{i, j})
				}
			}
		}

	case 'B':
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				fileDist := math.Abs(float64(pieceFile - j))
				rowDist := math.Abs(float64(pieceRow - i))
				if (fileDist == rowDist) && !checkPieceInWay(board, pieceRow, pieceFile, i, j) {
					spacesCanMove = append(spacesCanMove, Square{i, j})
				}
			}
		}

	case 'K':
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if currPiece.teamID == 0 {
					board, success, isMate := whiteMovement(board, i, j, 'K', "")
					if isMate {
						fmt.Println("failure")
					}
					board[0][0] = emptySquare // do not remove
					if success {
						spacesCanMove = append(spacesCanMove, Square{i, j})
					}
				} else {
					board, success, isMate := blackMovement(board, i, j, 'K', "")
					if isMate {
						fmt.Println("failure")
					}
					board[0][0] = emptySquare // do not remove
					if success {
						spacesCanMove = append(spacesCanMove, Square{i, j})
					}
				}
			}
		}
	}
	return spacesCanMove
}


func whiteMovement(board [8][8]Piece, row int, file int, pieceType int, input string) ([8][8]Piece, bool, bool) {
	prevBoard := board
	success := false
	if match, err := regexp.MatchString(`^O-O\n$`, input); err == nil && match && whiteShortCastle {
		if !checkPieceInWay(board, 7, 4, 7, 7) {
			board[7][6] = board[7][4]
			board[7][4] = emptySquare
			board[7][5] = board[7][7]
			board[7][7] = emptySquare
			whiteLongCastle = false
			whiteShortCastle = false
			success = true
		}
	} else if match, err := regexp.MatchString(`^O-O-O\n$`, input); err == nil && match && whiteLongCastle {
		if !checkPieceInWay(board, 7, 4, 7, 0) {
			board[7][2] = board[7][4]
			board[7][4] = emptySquare
			board[7][3] = board[7][0]
			board[7][0] = emptySquare
			whiteLongCastle = false
			whiteShortCastle = false
			success = true
		}
	}
	if board[row][file].teamID == 0 {
		return board, false, false
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fileDisamb, rowDisamb := false, false
			disambRow, disambFile := -1, -1
			parseDisamb(&fileDisamb, &rowDisamb, &disambRow, &disambFile, input)

			if rowDisamb && fileDisamb && i != int(8-(input[2]-'0')) {
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
					if fileDisamb && j == disambFile {
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
							success = true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].canBeEnPassant {
							board[row-1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							success = true
						}
					} else if i-row > 0 && i-row <= spacesCanMove && file == j && board[row][file].pieceID == '0' {
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
						success = true
					} else if !fileDisamb && file == j && (i-row > 2 || i-row < 1) {
						return board, false, false
					}

				case 'K':
					if checkPieceInWay(board, i, j, row, file) || rowDist > 1 || fileDist > 1 {
						return board, false, false
					}

					if (rowDist == fileDist || file-j == 0 || row-i == 0) && board[row][file].teamID != 0 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						success = true
						whiteLongCastle = false
						whiteShortCastle = false
					}

				case 'Q':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (!fileDisamb && !rowDisamb) {
						if rowDist == fileDist || file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) {
								return board, false, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							success = true
						}
					}

				case 'R':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if file-j == 0 || row-i == 0 {
							if j == 0 {
								whiteLongCastle = false
							} else if j == 7 {
								whiteShortCastle = false
							}
							if checkPieceInWay(board, i, j, row, file) {
								return board, false, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							success = true
						}
					}

				case 'N':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							success = true
						}
					}
				case 'B':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if (i%2 == j%2 && row%2 == file%2) || (i%2 != j%2 && row%2 != file%2) {
							if checkPieceInWay(board, i, j, row, file) {
								return board, false, false
							}

							if rowDist == fileDist {
								board[row][file] = board[i][j]
								board[i][j] = emptySquare
								success = true
							}
						}
					}
				}
			}
			if success {
				break
			}		
		}
		if success {
			break
		}		
	}
	if !success {
		return board, false, false
	}
	isCheck, kingSquare, pieceSquare := detectCheckOnKing(board)
	if isCheck {
		if board[kingSquare.squareRow][kingSquare.squareFile].teamID == 0 {
			return prevBoard, false, false
		} else {
			if isMate := isMate(int(kingSquare.squareRow), int(kingSquare.squareFile), int(pieceSquare.squareRow), int(pieceSquare.squareFile), board); isMate {
				return board, true, true
			}
			return board, true, false
		}
	}
	return board, true, false
}

func blackMovement(board [8][8]Piece, row int, file int, pieceType int, input string) ([8][8]Piece, bool, bool) {
	success := false
	prevBoard := board
	if match, err := regexp.MatchString(`^O-O\n$`, input); err == nil && match && blackShortCastle {
		if !checkPieceInWay(board, 0, 4, 0, 7) {
			board[0][6] = board[0][4]
			board[0][4] = emptySquare
			board[0][5] = board[0][7]
			board[0][7] = emptySquare
			blackShortCastle = false
			blackLongCastle = false
			success = true
		}
	} else if match, err := regexp.MatchString(`^O-O-O\n$`, input); err == nil && match && blackLongCastle {
		if !checkPieceInWay(board, 0, 4, 0, 0) {
			board[0][2] = board[0][4]
			board[0][4] = emptySquare
			board[0][3] = board[0][0]
			board[0][0] = emptySquare
			blackShortCastle = false
			blackLongCastle = false
			success = true
		}
	}
	if !success && board[row][file].teamID == 1 {
		return board, false, false
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fileDisamb, rowDisamb := false, false
			disambRow, disambFile := -1, -1
			parseDisamb(&fileDisamb, &rowDisamb, &disambRow, &disambFile, input)

			rowDist := math.Abs(float64(row - i))
			fileDist := math.Abs(float64(file - j))
			if board[i][j].teamID == 1 && board[i][j].pieceID == pieceType {
				switch pieceType {
				case 'P':
					spacesCanMove := 1
					if i == 1 {
						spacesCanMove = 2
					}
					if fileDisamb && j == disambFile {
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
							success = true
						} else if ((file == j+1 || file == j-1) && row == i) && board[row][file].canBeEnPassant {
							board[row+1][file] = board[i][j]
							board[i][j] = emptySquare
							board[row][file] = emptySquare
							success = true
						}
					}
					if row-i > 0 && row-i <= spacesCanMove && file == j && board[row][file].pieceID == '0' {
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
						success = true
					} else if !fileDisamb && file == j && (row - i > 2 || row - i < 1) {
						return board, false, false
					}
	
				case 'K':
					if checkPieceInWay(board, i, j, row, file) || rowDist > 1 || fileDist > 1 {
						return board, false, false
					}
					if (rowDist == fileDist || file-j == 0 || row-i == 0) && board[row][file].teamID != 1 {
						board[row][file] = board[i][j]
						board[i][j] = emptySquare
						success = true
						blackShortCastle = false
						blackLongCastle = false
					}

				case 'Q':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if rowDist == fileDist || file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) {
								return board, false, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							success = true
						}
					}

				case 'R':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if file-j == 0 || row-i == 0 {
							if checkPieceInWay(board, i, j, row, file) {
								return board, false, false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							success = true
						}
					}

				case 'N':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if (rowDist == 2 && fileDist == 1) || (rowDist == 1 && fileDist == 2) {
							if j == 0 {
								whiteLongCastle = false
							} else if j == 7 {
								whiteShortCastle = false
							}
							board[row][file] = board[i][j]
							board[i][j] = emptySquare
							success = true
						}
					}

				case 'B':
					if (rowDisamb && i == disambRow) || (fileDisamb && j == disambFile) || (fileDisamb == rowDisamb) {
						if (i%2 == j%2 && row%2 == file%2) || (i%2 != j%2 && row%2 != file%2) {
							if checkPieceInWay(board, i, j, row, file) {
								return board, false, false
							}

							if rowDist == fileDist {
								board[row][file] = board[i][j]
								board[i][j] = emptySquare
								success = true
							}
						}
					}
				}
			}
			if success {
				break;
			}
		}
		if success {
			break;
		}
	}
	if !success {
		return board, false, false
	}
	isCheck, kingSquare, pieceSquare := detectCheckOnKing(board)
	if isCheck {
		if board[kingSquare.squareRow][kingSquare.squareFile].teamID == 1 {
			return prevBoard, false, false
		} else {
			if isMate := isMate(int(kingSquare.squareRow), int(kingSquare.squareFile), int(pieceSquare.squareRow), int(pieceSquare.squareFile), board); isMate {
				return board, true, true
			}
			return board, true, false
		}
	}
	return board, true, false
}

func parseDisamb(fileDisamb* bool, rowDisamb* bool, disambRow* int, disambFile* int, input string) {
	if match, err := regexp.MatchString(`^[a-h]x[a-h][1-8]\n$`, input); err == nil && match {
		*disambFile = int(input[0] - 'a')
		*fileDisamb = true
	} else if match, err := regexp.MatchString(`^[a-hNKQBR][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
		*rowDisamb = true
		*disambRow = int(input[1] - '0')
	} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]x?[a-h][1-8]\n$`, input); err == nil && match {
		*disambFile = int(input[1] - 'a')
		*fileDisamb = true
	} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h][1-8]x?[a-h][1-8]\n$`, input); err == nil && match {
		*rowDisamb = true
		*fileDisamb = true
		*disambRow = int(input[1] - '0')
		*disambFile = int(input[1] - 'a')
	}
}
