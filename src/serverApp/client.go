package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conn, conErr := net.Dial("tcp", "127.0.0.1:20000")
	if conErr != nil {
		fmt.Println("Failed to connect to the server")
		return
	}
	defer conn.Close()
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	teamBuf := make([]byte, 1)
	n, error1 := conn.Read(teamBuf)
	if error1 != nil || n == 0 {
		fmt.Println("Failed to read team byte")
		return
	}
	myTurn := false
	if teamBuf[0] == '0' {
		myTurn = true
	} else if teamBuf[0] == '1' {
		myTurn = false
	}

	buf := make([]byte, 4096)
	c := make(chan string)
	for {
		go readString(c, conn, buf)
		str := <- c
		fmt.Printf("%s", str)
		if myTurn{
			scanner.Scan()
			conn.Write([]byte(scanner.Text()))
		}
		myTurn = !myTurn
	}
}

func readString(c chan string, conn net.Conn, buf []byte) {
	n, err := conn.Read(buf)
	if n < 512 {

	}
	if err != nil {
		return
	}
	c <- string(buf[:n])
}
