package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)


type Square struct {
	squareRow  int16
	squareFile int16
}


func main() {
	board := initializeBoard()
	fmt.Println("Enter proper chess notation to make a move.")
	whiteTurn := true
	isMate := false
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
		board, success, isMate = executeTurn(input, whiteTurn, board)
		if success {
			if isMate {
				return
			}
			whiteTurn = !whiteTurn
		} else {
			fmt.Println("Please input a valid move.")
		}
	}

}


func executeTurn(input string, whiteTurn bool, board [8][8]Piece) ([8][8]Piece, bool, bool) {
	isMate := false
	prevBoard := board
	pieceType := 0
	if !parseInput(input, &pieceType) {
		return board, false, isMate
	}
	row := int16(8 - (input[len(input)-2] - '0'))
	file := int16(input[len(input)-3] - 'a')

	if row > 8 || row < 0 || file > 8 || file < 0 {
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
		}
		if success {
			tempBoard = removeWhiteEnPassant(tempBoard)
			return tempBoard, success, isMate
		} else {
			return prevBoard, success, isMate
		}
	}

	return board, false, isMate
}

func parseInput(input string, pieceType *int) bool {
	if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?x[a-h][1-8]\n$`, input); err == nil && match {
		if input[0] >= 97 && input[0] <= 104 {
			*pieceType = 'P'
		} else {
			*pieceType = int(input[0])
		}
		return true
	} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?[a-h][1-8]\n$`, input); err == nil && match {
		if input[0] >= 97 && input[0] <= 104 {
			*pieceType = 'P'
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

func isMate(kingRow int, kingFile int, pieceFile int, pieceRow int, board [8][8]Piece) bool {
	var spacesBetween []Square = getSpacesBetween(board, kingRow, kingFile, pieceRow, pieceFile)
	var spacesCanMove []Square
	if len(getSpacesCanMove(int16(kingRow), int16(kingFile), board)) > 0 {
		return false
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].teamID == board[kingRow][kingFile].teamID {
				spacesCanMove = getSpacesCanMove(int16(i), int16(j), board)
				for _, squareBetween := range spacesBetween {
					for _, squareMove := range spacesCanMove {
						if squareMove.squareRow == squareBetween.squareRow && squareMove.squareFile == squareBetween.squareFile {
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
					if checkPieceInWay(board, int16(i), int16(j), int16(wkRow), int16(wkFile)) {
						wkCheck = false
					} else {
						wkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (math.Abs(bkFileDist) == math.Abs(bkRowDist) || bkFileDist == 0 || bkRowDist == 0) && !bkCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(bkRow), int16(bkFile)) {
						bkCheck = false
					} else {
						bkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				}
			case 'R':
				if currPiece.teamID == 1 && (wkFileDist == 0 || wkRowDist == 0) && !wkCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(wkRow), int16(wkFile)) {
						wkCheck = false
					} else {
						wkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (bkFileDist == 0 || bkRowDist == 0) && !bkCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(bkFile), int16(bkFile)) {
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
					if checkPieceInWay(board, int16(i), int16(j), int16(wkRow), int16(wkFile)) {
						wkCheck = false
					} else {
						wkCheck = true
						checkingPieceRow = i
						checkingPieceFile = j
					}
				} else if currPiece.teamID == 0 && (bkFileDist == bkRowDist) && !bkCheck {
					if checkPieceInWay(board, int16(i), int16(j), int16(bkRow), int16(bkFile)) {
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
		return true, Square{int16(wkRow), int16(wkFile)}, Square{int16(checkingPieceRow), int16(checkingPieceFile)}
	}
	if bkCheck {
		return true, Square{int16(bkRow), int16(bkFile)}, Square{int16(checkingPieceRow), int16(checkingPieceFile)}
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
