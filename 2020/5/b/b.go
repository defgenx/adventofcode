package main

import (
	"2020/common"
	"fmt"
	"log"
	"strings"
)

const(
	MaxRow = 128
	MaxLine = 8
)

func main() {
	file := common.OpenInputFile("5/input.txt")
	buffer := common.ReadBuffer(file)
	mySeatID := 0
	var seatMap [MaxRow][MaxLine]int
	for buffer.Scan() {
		line := buffer.Text()
		splittedLine := strings.Split(line, "")
		fbPart := splittedLine[:7]
		rowNb := findRow(fbPart)
		lrPart := splittedLine[7:]
		lineNb := findLine(lrPart)
		newHighestSeatID := rowNb * 8 + lineNb
		seatMap[rowNb][lineNb] = newHighestSeatID
	}
	for i := 0; i < MaxRow; i ++ {
		for j := 0; j < MaxLine; j ++ {
			fmt.Print(seatMap[i][j], " | ")
			// We  must add the first test only because we wanted to print the grid
			if j > 0 && (j+1) < MaxRow &&
				seatMap[i][j] == 0 &&
				seatMap[i][j-1] != 0 &&
				seatMap[i][j+1] != 0 {
				mySeatID = i * 8 + j
			}
		}
		fmt.Println("")
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("My seat ID: %d", mySeatID)
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