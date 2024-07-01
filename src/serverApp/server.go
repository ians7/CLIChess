package main

import (
	"fmt"
	"net"
	"strconv"
)


type Game struct {
	board [8][8]Piece
	turn bool
}

func main() {
	// Establish the server on port 20000
	ln, err := net.Listen("tcp", ":20000")
	if err != nil {
		fmt.Println("Failed to create server")
		return
	}
	defer ln.Close()
	fmt.Println("Creating server on localhost:20000")
	player1Conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection")
	}
	player1Conn.Write([]byte("0"))
	player2Conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection")
	}
	player2Conn.Write([]byte("1"))
	fmt.Println("Two players detected, creating game...\n")
	for {
		game := createGame()
		gameLoop(player1Conn, player2Conn, game)
		if err == nil {
			break
		}
	}
}


func createGame() Game {
	return Game{initializeBoard(), true}
}

func gameLoop(p1Conn net.Conn, p2Conn net.Conn, game Game) error {
	success := false
	isMate := false
	whiteMsg := "White turn: "
	blackMsg := "Black turn: " 
	for {
		buf := make([]byte, 64)
		whiteBoard := []byte(createWhiteBoardMsg(game.board))
		blackBoard := []byte(createBlackBoardMsg(game.board))
		p1Conn.Write(whiteBoard)
		p2Conn.Write(blackBoard)
		input, err := handleInput(game, whiteMsg, blackMsg, buf, p1Conn, p2Conn)	
		if err != nil {
			return fmt.Errorf("Error reading user input")
		}
		game.board, success, isMate = executeTurn(input, game.turn, game.board)
		if (!success) {
			handleImproperInput(game, whiteBoard, p1Conn, p2Conn)
		} else {
			if isMate {
				handleMate(game, p1Conn, p2Conn)
				break
			}
			game.turn = !game.turn
		}
	}
	return nil
}

func handleInput(game Game, whiteMsg string, blackMsg string, buf []byte, p1Conn net.Conn, p2Conn net.Conn) (string, error) {
	n := -1
	var err error
	if game.turn {
		p1Conn.Write([]byte(whiteMsg))
		p2Conn.Write([]byte("Waiting for opponent...\n"))
		n, err = p1Conn.Read(buf)
		if err != nil {
			return "", fmt.Errorf("Error reading user input") 
		}
	} else {
		p2Conn.Write([]byte(blackMsg))
		p1Conn.Write([]byte("Waiting for opponent...\n"))
		n, err = p2Conn.Read(buf)
		if err != nil {
			return "", fmt.Errorf("Error reading user input") 
		}
	}
	return string(buf[:n]), nil
}

func handleImproperInput(game Game, board []byte, p1Conn net.Conn, p2Conn net.Conn) {
	if game.turn {
		p1Conn.Write([]byte("invalid input\n"))
	} else {
		p2Conn.Write([]byte("invalid input\n"))
	}
}

func handleMate(game Game, p1Conn net.Conn, p2Conn net.Conn) {
	whiteBoard := []byte(createWhiteBoardMsg(game.board))
	blackBoard := []byte(createBlackBoardMsg(game.board))
	p1Conn.Write(whiteBoard)
	p2Conn.Write(blackBoard)
	if !game.turn {
		p2Conn.Write([]byte("Black wins!\n"))
		p1Conn.Write([]byte("Black wins!\n"))
	} else {
		p1Conn.Write([]byte("White wins!\n"))
		p2Conn.Write([]byte("White wins!\n"))
	}
}

func createWhiteBoardMsg(board [8][8]Piece) string {
	msg := ""
	msg = msg + W + "  a   b   c   d   e   f   g   h\n"
	msg = msg + Br + " -------------------------------\n"
	bgColor := bgRed
	colorBool := true
	pieceColor := B
	for i := 0; i < 8; i++ {
		msg = msg + Br + "|" 
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
			msg = msg + bgColor + pieceColor + " " + string(board[i][j].icon) + " " +  bgBlack + "|" + Br + bgBlack
			colorBool = !colorBool
		}
		row := 8-i
		msg = msg + W + " " + strconv.Itoa(row)
		msg = msg + Br + "\n -------------------------------\n" + W
		colorBool = !colorBool
	}
	return msg
}
func createBlackBoardMsg(board [8][8]Piece) string {
	msg := ""
	msg = msg + W + "  h   g   f   e   d   c   b   a\n"
	msg = msg + Br + " -------------------------------\n"
	bgColor := bgRed
	colorBool := true
	pieceColor := B
	for i := 7; i >= 0; i-- {
		msg = msg + Br + "|" 
		for j := 7; j >= 0; j-- {
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
			msg = msg + bgColor + pieceColor + " " + string(board[i][j].icon) + " " +  bgBlack + "|" + Br + bgBlack
			colorBool = !colorBool
		}
		row := i
		msg = msg + W + " " + strconv.Itoa(row)
		msg = msg + Br + "\n -------------------------------\n" + W
		colorBool = !colorBool
	}
	return msg
}

func UNUSED(x ...interface{}) {}
