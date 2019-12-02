package common

import (
	"bufio"
	"log"
	"os"
)

func OpenInputFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func ReadBuffer(file *os.File) *bufio.Scanner {
	return bufio.NewScanner(file)
}
