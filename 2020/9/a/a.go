package main

import (
	"2020/common"
	"log"
	"strconv"
)

func main() {
	file := common.OpenInputFile("9/input.txt")
	buffer := common.ReadBuffer(file)
	val := 0
	bufferPreamble := make([]int, 0)
	for buffer.Scan() {
		found := false
		line := buffer.Text()
		intVal, _ := strconv.Atoi(line)
		if len(bufferPreamble) < 25 {
			bufferPreamble = append(bufferPreamble, intVal)
		} else {
		breakFind:
			for _, i := range bufferPreamble {
				for _, j := range bufferPreamble {
					if j == i {
						continue
					}
					if intVal == (i + j) {
						found = true
						break breakFind
					}
				}
			}
			if found == false {
				val = intVal
				break
			}
			bufferPreamble = bufferPreamble[1:]
			bufferPreamble = append(bufferPreamble, intVal)
		}
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("First number: %v", val)
}