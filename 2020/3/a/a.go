package main

import (
	"2020/common"
	"log"
	"strings"
)

func main() {
	file := common.OpenInputFile("3/input.txt")
	buffer := common.ReadBuffer(file)
	var trees = 0
	var firstRow = true
	var right = 3
	for buffer.Scan() {
		// Ignore first row
		if firstRow {
			firstRow = false
			continue
		}
		// Split each char of the line
		slice := strings.Split(buffer.Text(), "")
		// If end of line is reached, return to the beginning
		if len(slice)-1 < right {
			right = right - len(slice)
		}
		if slice[right] == "#" {
			trees++
		}
		right += 3
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total tree encountered: %d", trees)
}
