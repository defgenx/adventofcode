package main

import (
	"2021/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("5/input.txt")
	var grid [1000][1000]int
	var count = 0
	buffer := common.ReadBuffer(file)
	for buffer.Scan() {
		rowText := buffer.Text()
		ranges := strings.Split(rowText, " -> ")
		leftVals := strings.Split(ranges[0], ",")
		rightVals := strings.Split(ranges[1], ",")
		x1, _ := strconv.Atoi(leftVals[0])
		y1, _ := strconv.Atoi(leftVals[1])
		x2, _ := strconv.Atoi(rightVals[0])
		y2, _ := strconv.Atoi(rightVals[1])
		if (x1 == x2) || (y1 == y2) {
			log.Print(rowText)
			if x1 > x2 {
				tmpX := x2
				x2 = x1
				x1 = tmpX
			}
			if y1 > y2 {
				tmpY := y2
				y2 = y1
				y1 = tmpY
			}
			for _, y := range common.Range(y1, y2) {
				for _, x := range common.Range(x1, x2) {
					grid[y][x]++
				}
			}
		}
	}
	for _, val := range grid {
		for _, i := range val {
			if i > 1 {
				count++
			}
		}
	}
	log.Printf("Zone count: %d", count)
}
