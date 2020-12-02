package main

import (
	"2020/common"
	"log"
	"strconv"
)

func main() {
	file := common.OpenInputFile("1/input.txt")
	buffer := common.ReadBuffer(file)
	var list []int
	for buffer.Scan() {
		val, _ := strconv.Atoi(buffer.Text())
		list = append(list, val)
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[i] + list[j] == 2020 {
				log.Printf("Correct value is: %d", list[i] * list[j])
				return
			}
		}
	}
}
