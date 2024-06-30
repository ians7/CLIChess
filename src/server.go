package main

import (
	"fmt"
	"net"
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
	fmt.Println("Two players detected, creating game...")
	gameLoop(player1Conn, player2Conn)

}


func createGame() Game {
	return Game{initializeBoard(), true}
}

func gameLoop(p1Conn net.Conn, p2Conn net.Conn) {
	game := createGame()
	success := false
	isMate := false
	whiteMsg := "White turn: "
	blackMsg := "Black turn: " 
	buf := make([]byte, 4096)
	for {
		boardMsg := []byte(createBoardMsg(game.board))
		p1Conn.Write(boardMsg)
		p2Conn.Write(boardMsg)
		var err error	
		n := -1
		if game.turn == true {
			p1Conn.Write([]byte(whiteMsg))
			p2Conn.Write([]byte("Waiting for opponent..."))
			n, err = p1Conn.Read(buf)
			if err != nil {
				fmt.Println("invalid input")
				return
			}
		} else {
			p2Conn.Write([]byte(blackMsg))
			p1Conn.Write([]byte("Waiting for opponent..."))
			n, err = p2Conn.Read(buf)
			if err != nil {
				fmt.Println("invalid input")
				return
			}
		}
		input := string(buf[:n])
		game.board, success, isMate = executeTurn(input, game.turn, game.board)
		if (!success) {
			if !game.turn {
				p1Conn.Write([]byte("improper input"))
			} else {
				p2Conn.Write([]byte("improper input"))
			}
		} else {
			if isMate {
				fmt.Println("Game Over!")
				return
			}
			game.turn = !game.turn
		}
	}
}
func createBoardMsg(board [8][8]Piece) string {
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
		msg = msg + W + string(row)
		msg = msg + Br + "\n -------------------------------\n" + W
		colorBool = !colorBool
	}
	return msg
}
