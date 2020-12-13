package main

import (
	"2021/common"
	"log"
	"strconv"
)

func main() {
	file := common.OpenInputFile("1/input.txt")
	buffer := common.ReadBuffer(file)
	var countIncr = 0
	var buff = 0
	for buffer.Scan() {
		val, _ := strconv.Atoi(buffer.Text())
		if buff == 0 {
			buff = val
			continue
		}
		if buff < val {
			countIncr++
		}
		buff = val
	}
	log.Printf("Number of increase: %d", countIncr)
}
