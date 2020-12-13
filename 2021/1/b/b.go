package main

import (
	"2021/common"
	"log"
	"strconv"
)

func main() {
	file := common.OpenInputFile("1/input.txt")
	buffer := common.ReadBuffer(file)
	var rawList []int
	var countIncr = 0
	for buffer.Scan() {
		val, _ := strconv.Atoi(buffer.Text())
		rawList = append(rawList, val)
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(rawList); i++ {
		j := i+1
		if j+3 > len(rawList) || len(rawList[i:i+3]) < 3 || len(rawList[j:j+3]) < 3 {
			break
		}
		//log.Print(rawList[i:i+3])
		//log.Print(rawList[j:j+3])
		if common.Sum(rawList[i:i+3]) < common.Sum(rawList[j:j+3]) {
			countIncr++
		}
	}
	log.Printf("Number of increase: %d", countIncr)
}
