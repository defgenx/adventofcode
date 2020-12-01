package common

import (
	"bufio"
	"io/ioutil"
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

func ReadFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("File reading error", err)
	}
	return data
}
