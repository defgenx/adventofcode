package main

import (
	"2021/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("2/input.txt")
	depth := 0
	horizontal := 0
	aim := 0
	buffer := common.ReadBuffer(file)
	for buffer.Scan() {
		splittedVals := strings.Split(buffer.Text(), " ")
		direction := splittedVals[0]
		val, _ := strconv.Atoi(splittedVals[1])
		if direction == "forward" {
			horizontal += val
			depth += aim * val
		} else if direction == "up" {
			aim -= val
		} else {
			aim += val
		}
	}
	log.Printf("Number: %d", depth*horizontal)
}
