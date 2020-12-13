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
	bufferJolter := make([]int, 0)
	bufferJolter = append(bufferJolter, 0)
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
	bufferTmp := make([]int, len(bufferJolter))
	bufferTmp[0] = 1
	for i := 0; i < len(bufferJolter) ; i++ {
		for j := 0; j < i; j++ {
			if bufferJolter[i] - bufferJolter[j] <= 3 {
				bufferTmp[i] += bufferTmp[j]
			}
		}
	}
	log.Printf("Jolter max combination: %v", bufferTmp[len(bufferTmp) - 1])
}