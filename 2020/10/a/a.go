package main

import (
	"2020/common"
	"log"
	"sort"
	"strconv"
)

func main() {
	file := common.OpenInputFile("10/input.txt")
	buffer := common.ReadBuffer(file)
	jolter := 0
	diff1 := 0
	diff3 := 0
	bufferJolter := make([]int, 0)
	for buffer.Scan() {
		line := buffer.Text()
		intVal, _ := strconv.Atoi(line)
		bufferJolter = append(bufferJolter, intVal)
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Ints(bufferJolter)
	bufferJolter = append(bufferJolter, bufferJolter[len(bufferJolter) - 1] + 3)
	for i := 0; i < len(bufferJolter); i++ {
		diff := bufferJolter[i] - jolter
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
		jolter = bufferJolter[i]
	}
	log.Printf("Jolter diff: %v", diff1 * diff3)
}