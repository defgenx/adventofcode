package main

import (
	"2020/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("8/input.txt")
	buffer := common.ReadBuffer(file)
	parsedContent := make([]map[string]int, 0)
	counter := 0
	for buffer.Scan() {
		line := buffer.Text()
		splitRow := strings.Split(line, " ")
		builtMap := make(map[string]int, 0)
		instruction, arg := splitRow[0], splitRow[1]
		intVal, _ := strconv.Atoi(arg)
		builtMap[instruction] = intVal
		parsedContent = append(parsedContent, builtMap)
	}
	log.Print(parsedContent)
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	posBuffer := make([]int, 0)

	for index := 0; index < len(parsedContent); index++ {
		if containsPos(index, posBuffer) {
			break
		}
		posBuffer = append(posBuffer, index)
		if jmp, exists := parsedContent[index]["jmp"]; exists {
			index = (index + jmp) - 1
		} else if acc, exists := parsedContent[index]["acc"]; exists {
			counter += acc
		}
	}
	log.Printf("Total count: %v", counter)
}

func containsPos(val int, positions []int) bool {
	for _, pos := range positions {
		if pos == val {
			return true
		}
	}
	return false
}