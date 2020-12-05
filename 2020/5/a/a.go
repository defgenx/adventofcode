package main

import (
	"2020/common"
	"log"
	"strings"
)

func main() {
	file := common.OpenInputFile("5/input.txt")
	buffer := common.ReadBuffer(file)
	highestSeatID := 0
	for buffer.Scan() {
		line := buffer.Text()
		log.Print(line)
		splittedLine := strings.Split(line, "")
		fbPart := splittedLine[:7]
		rowNb := findRow(fbPart)
		log.Printf("Row: %d", rowNb)
		lrPart := splittedLine[7:]
		lineNb := findLine(lrPart)
		log.Printf("Line: %d", lineNb)
		newHighestSeatID := rowNb * 8 + lineNb
		log.Printf("Seat ID: %d", newHighestSeatID)
		if highestSeatID < newHighestSeatID {
			highestSeatID = newHighestSeatID
		}
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Highest seat ID: %d", highestSeatID)
}

func findRow(parts []string) int {
	var lower = 0
	var upper = 127
	for _, part := range parts {
		if part == "F" {
			upper = computeUpperRight(upper, lower)
		} else {
			lower = computeLowerLeft(lower, upper)
		}
	}
	return upper
}

func findLine(parts []string) int {
	var left = 0
	var right = 7
	for _, part := range parts {
		if part == "L" {
			right = computeUpperRight(right, left)
		} else {
			left = computeLowerLeft(left, right)
		}
	}
	return right
}

func computeUpperRight(a, b int) int {
	return a - ((a - b) + 1) / 2
}

func computeLowerLeft(a, b int) int {
	return a + ((b - a) + 1) / 2
}