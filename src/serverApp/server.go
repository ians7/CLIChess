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
	player2Conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Failed to accept connection")
	}
	fmt.Println("Two players detected, creating game...\n")
	for {
		game := createGame()
		gameLoop(player1Conn, player2Conn, game)
		return
	}
}


func createGame() Game {
	return Game{initializeBoard(), true}
}

func gameLoop(p1Conn net.Conn, p2Conn net.Conn, game Game) error {
	isMate := false
	success := false
	whiteMsg := "White turn: "
	blackMsg := "Black turn: " 
	whiteBoard := createWhiteBoardMsg(game.board)
	blackBoard := createBlackBoardMsg(game.board)
	p1Conn.Write([]byte("1" + whiteBoard + whiteMsg))
	p2Conn.Write([]byte("0" + blackBoard + "Waiting for opponent...\n"))
	for {
		buf := make([]byte, 64)

		input, err := handleInput(game, buf, p1Conn, p2Conn)	
		if err != nil {
			return fmt.Errorf("Error reading user input")
		}
		game.board, success, isMate = executeTurn(input, game.turn, game.board)
		whiteBoard := createWhiteBoardMsg(game.board)
		blackBoard := createBlackBoardMsg(game.board)
		if success {
			game.turn = !game.turn
			if isMate {
				handleMate(game, p1Conn, p2Conn)
				break
			}
			if game.turn {
				p1Conn.Write([]byte("1" + whiteBoard + whiteMsg))
				p2Conn.Write([]byte("0" + blackBoard + "Waiting for opponent...\n"))
			} else {
				p1Conn.Write([]byte("0" + whiteBoard + "Waiting for opponent...\n"))
				p2Conn.Write([]byte("1" + blackBoard + blackMsg))
			}
		} else {
			fmt.Println("input = ", input)
			if game.turn {
				p1Conn.Write([]byte("1" + "\rInvalid input.\n"))
			} else {
				p2Conn.Write([]byte("1" + "\rInvalid input.\n"))
			}
		}
	}
	return nil
}

func handleInput(game Game, buf []byte, p1Conn net.Conn, p2Conn net.Conn) (string, error) {
	n := -1
	var err error
	if game.turn {
		n, err = p1Conn.Read(buf)
		if err != nil {
			return "", fmt.Errorf("Error reading user input") 
		}
	} else {
		n, err = p2Conn.Read(buf)
		if err != nil {
			return "", fmt.Errorf("Error reading user input") 
		}
	}
	return string(buf[:n]), nil
}

func handleMate(game Game, p1Conn net.Conn, p2Conn net.Conn) {
	whiteBoard := createWhiteBoardMsg(game.board)
	blackBoard := createBlackBoardMsg(game.board)
	if !game.turn {
		p2Conn.Write([]byte(whiteBoard + "Black wins!\n"))
		p1Conn.Write([]byte(blackBoard + "Black wins!\n"))
	} else {
		p1Conn.Write([]byte(whiteBoard + "White wins!\n"))
		p2Conn.Write([]byte(blackBoard + "White wins!\n"))
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
		row := 8 - i
		msg = msg + W + " " + strconv.Itoa(row)
		msg = msg + Br + "\n -------------------------------\n" + W
		colorBool = !colorBool
	}
	return msg
}

func UNUSED(x ...interface{}) {}
