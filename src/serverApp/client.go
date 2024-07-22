package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"os/exec"
	"regexp"
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

	buf := make([]byte, 4096)
	inputc := make(chan string)
	myTurn := -1
	go readString(inputc, conn, buf)
	str := <- inputc
	fmt.Printf("%s", str[1:len(str)])

	for {
		if str[0] - 48 == 0 {
			myTurn = 0
		} else {
			myTurn = 1
		}	
		if myTurn == 1 {
			for {
				scanner.Scan()
				if verifyInput(scanner.Text()) {
					conn.Write([]byte(scanner.Text()))
					break
				} else {
					fmt.Println("\rInvalid input.\n")
				}
			}	
		}
		go readString(inputc, conn, buf)
		str = <- inputc
		fmt.Printf("%s", str[1:len(str)])
	}
}

func readString(c chan string, conn net.Conn, buf []byte) {
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	c <- string(buf[:n])
}

func verifyInput(input string) bool {
	if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?x[a-h][1-8]$`, input); err == nil && match {
		return true
	} else if match, err := regexp.MatchString(`^[a-hNKQBR][a-h]?[1-8]?[a-h][1-8]$`, input); err == nil && match {
		return true
	} else if match, err := regexp.MatchString(`^[NKQBR][a-h][1-8]$`, input); err == nil && match {
		return true
	} else if match, err := regexp.MatchString(`^[a-h][1-8]$`, input); err == nil && match {
		return true
	} else if match, err := regexp.MatchString(`^O-O(-O)?$`, input); err == nil && match {
		return true
	} else {
		return false
	}
}
