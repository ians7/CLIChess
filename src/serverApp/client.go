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
	conn, conErr := net.Dial("tcp", ":20000")
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
	if error1 != nil {
		fmt.Println("Failed to read team byte")
		return
	}
	myTurn := false
	if teamBuf[0] == '0' {
		myTurn = true
	} else if teamBuf[0] == '1' {
		myTurn = false
	}

	var err error = nil
	buf := make([]byte, 512)
	for {
		for {
			n, err = conn.Read(buf)
			fmt.Printf("%s", string(buf[:n]))
			if n < 512 || err != nil {
				break
			}
		}
		if myTurn{
			scanner.Scan()
			conn.Write([]byte(scanner.Text()))
		}
		myTurn = !myTurn
	}
}


