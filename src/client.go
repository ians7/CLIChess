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
	conn, conErr := net.Dial("tcp", "172.19.142.219:20000")
	defer conn.Close()
	if conErr != nil {
		fmt.Println("Failed to connect to the server")
		return
	}
	buf := make([]byte, 4096)
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

	var err error
	for {
		n, err = conn.Read(buf)
		if err != nil {
			return
		}
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("%s", string(buf[:n]))
		if myTurn{
			scanner.Scan()
			conn.Write([]byte(scanner.Text()))
		}
		myTurn = !myTurn
	}
}


