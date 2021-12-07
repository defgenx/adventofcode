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

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Range(start, end int) []int {
	s := make([]int, 0, 1+(end-start))
	for start <= end {
		s = append(s, start)
		start += 1
	}
	return s
}
