package main

import (
	"fmt"
	"math"
)


var (
	whiteLongCastle = true
	blackLongCastle = true
	whiteShortCastle = true
	blackShortCastle = true
)

func executeTurn(input string, whiteTurn bool, board [8][8]Piece) ([8][8]Piece, bool, bool) {
	isMate := false
	row := 0
	file := 0
	prevBoard := board
	pieceType := parsePieceType(input)

	if pieceType != -1 && pieceType != 'O' {
		row = int(8 - (input[len(input)-1] - '0'))
		file = int(input[len(input)-2] - 'a')
	} 

	if (row > 8 || row < 0 || file > 8 || file < 0) && pieceType != -1 && pieceType != 'O' {
		return board, false, isMate
	}

	if whiteTurn {
		tempBoard, success, isMate := whiteMovement(board, row, file, pieceType, input)
		if isMate {
			isMate = true
			fmt.Println("White wins!")
		}
		if success {
			tempBoard = removeBlackEnPassant(tempBoard)
			return tempBoard, success, isMate
		} else {
			return prevBoard, success, isMate
		}
	} else {
		tempBoard, success, isMate := blackMovement(board, row, file, pieceType, input)
		if isMate {
			isMate = true
			fmt.Println("Black wins!")
		} else if success {
			tempBoard = removeWhiteEnPassant(tempBoard)
			return tempBoard, success, isMate
		} else {
			return prevBoard, success, isMate
		}
	}

	return board, false, isMate
}

func parsePieceType(input string) int {
	pieceType := -1
	for i := 0 ; i < len(input) ; i++ {
		if input[i] >= 65 && input[i] <= 90 {
			pieceType = int(input[i])
		}
	}
	if pieceType == -1 && pieceType != 'O' {
		pieceType = 'P'
	} else if pieceType == 'O' {
		pieceType = 'O'
		fmt.Println("pieceType =", pieceType)
	}
	return int(pieceType)
}

func isMate(kingRow int, kingFile int, pieceRow int, pieceFile int, board [8][8]Piece) bool {
	var spacesBetween []Square = getSpacesBetween(board, kingRow, kingFile, pieceRow, pieceFile)
	var spacesCanMove []Square
	if len(getSpacesCanMove(kingRow, kingFile, board)) > 0 {
		return false
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].teamID == board[kingRow][kingFile].teamID {
				spacesCanMove = getSpacesCanMove(i, j, board)
				for _, squareBetween := range spacesBetween {
					for _, squareMove := range spacesCanMove {
						if squareMove.row == squareBetween.row && squareMove.file == squareBetween.file {
							return false
						}	
					}
				}
			}
		}
	}
	return true
}

func detectCheckOnKing(board [8][8]Piece) (bool, Square, Square) {
	kingCount := 0
	wkRow := -1 
	wkFile := -1 
	bkRow := -1
	bkFile := -1
	checkingPieceRow := -1
	checkingPieceFile := -1
	wkCheck := false
	bkCheck := false
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
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist && wkFileDist == 1 && wkRowDist == 1) && !wkCheck {
					wkCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				} else if currPiece.teamID == 0 && (bkFileDist == 1 && bkRowDist == 1) && !bkCheck {
					bkCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				}

			case 'Q':
				if currPiece.teamID == 1 && (math.Abs(wkFileDist) == math.Abs(wkRowDist) || wkFileDist == 0 || wkRowDist == 0) && !wkCheck {
					if checkPieceInWay(board, i, j, wkRow, wkFile) {
						wkCheck = false
					} else {
						wkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (math.Abs(bkFileDist) == math.Abs(bkRowDist) || bkFileDist == 0 || bkRowDist == 0) && !bkCheck {
					if checkPieceInWay(board, i, j, bkRow, bkFile) {
						bkCheck = false
					} else {
						bkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			case 'R':
				if currPiece.teamID == 1 && (wkFileDist == 0 || wkRowDist == 0) && !wkCheck {
					if checkPieceInWay(board, i, j, wkRow, wkFile) {
						wkCheck = false
					} else {
						wkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (bkFileDist == 0 || bkRowDist == 0) && !bkCheck {
					if checkPieceInWay(board, i, j, bkFile, bkFile) {
						bkCheck = false
					} else {
						bkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			case 'N':
				if currPiece.teamID == 1 && ((wkRowDist == 2 && wkFileDist == 1) || (wkRowDist == 1 && wkFileDist == 2)) && !wkCheck {
					wkCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				} else if currPiece.teamID == 0 && ((bkRowDist == 2 && bkFileDist == 1) || (bkRowDist == 1 && bkFileDist == 2)) && !bkCheck {
					bkCheck = true
					checkingPieceRow = i
					checkingPieceFile = j
				}
			case 'B':
				if currPiece.teamID == 1 && (wkFileDist == wkRowDist) && !wkCheck {
					if checkPieceInWay(board, i, j, wkRow, wkFile) {
						wkCheck = false
					} else {
						wkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (bkFileDist == bkRowDist) && !bkCheck {
					if checkPieceInWay(board, i, j, bkRow, bkFile) {
						bkCheck = false
					} else {
						bkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			}
		}
	}
	if wkCheck {
		return true, Square{wkRow, wkFile}, Square{checkingPieceRow, checkingPieceFile}
	}
	if bkCheck {
		return true, Square{bkRow, bkFile}, Square{checkingPieceRow, checkingPieceFile}
	}
	return false, Square{-1, -1}, Square{-1, -1}
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
