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
	fileContent := make([]int, 0)
	for buffer.Scan() {
		line := buffer.Text()
		intVal, _ := strconv.Atoi(line)
		fileContent = append(fileContent, intVal)
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}

	bufferPreamble := make([]int, 0)
	for _, intVal := range fileContent {
		found := false
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
	begin := 0
	tmpContiguous := make([]int, 0)
	sum := 0
	for i := begin; i < len(fileContent); i++ {
		if i == len(fileContent) - 1 {
			begin++
			i = begin
			tmpContiguous = make([]int, 0)
			sum = 0
			continue
		}
		tmpContiguous = append(tmpContiguous, fileContent[i])
		sum += fileContent[i]
		if sum == val {
			break
		}
	}
	min, max := contiguousMinAndMax(tmpContiguous)
	log.Printf("Contigous list: %v", tmpContiguous)
	log.Printf("Contigous min %d and max %d", min, max)
	log.Printf("Encryption weakness %d", min + max)
}

func contiguousMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}