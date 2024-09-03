package main

import (
	"os"
	"time"
	"log"
	"bufio"
	"io"
	"strconv"
	"fmt"
)

func newGameFile() os.File {
	err := os.Mkdir("gameLogs", 0755)
	if err != nil {
		fmt.Println(err)
	}
	fileName := time.Now().String()[:19]
	newFile, err := os.Create(string("gameLogs/" + fileName))
	if err != nil {
		log.Fatal(err)
	}
	return *newFile
}

func writeToFile(fp *os.File, isWhite bool, input string , isCheck bool , isMate bool ) {
	inputAsBytes := []byte(input)
	turnNum := 1
	br := bufio.NewReader(fp)
	for {
		b, err := br.ReadByte()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		if string(b) == "\n" {
			turnNum++
		}
	}
	if isWhite {
		str := strconv.Itoa(turnNum) + ". "
		fp.Write([]byte(str))
	}
	fp.Write(inputAsBytes)
	if isCheck && !isMate {
		fp.Write([]byte("+"))	
	} else if isMate {
		fp.Write([]byte("#"))
		return
	}
	if !isWhite {
		fp.Write([]byte("\n"))
	} else {
		fp.Write([]byte(" "))
	}
}

