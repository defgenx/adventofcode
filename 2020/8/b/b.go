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
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	posBuffer := make([]int, 0)
	valid := true
BreakEnd:
	for {
		// Iterate and change nos already tested and changed jmp / nop
		for index := 0; index < len(parsedContent); index++ {
			_, jmpExists := parsedContent[index]["jmp"]
			_, nopExists := parsedContent[index]["nop"]
			if !containsPos(index, posBuffer) && (jmpExists || nopExists) {
				editMap(index, parsedContent)
				posBuffer = append(posBuffer, index)
				break
			}
		}
		// Reinit tmp variables
		counter = 0
		valid = true
		posLastBuffer := make([]int, 0)
		for index := 0; index < len(parsedContent); index++ {
			if containsPos(index, posLastBuffer) {
				valid = false
				break
			}
			posLastBuffer = append(posLastBuffer, index)
			if jmp, exists := parsedContent[index]["jmp"]; exists {
				index = (index + jmp) - 1
			} else if acc, exists := parsedContent[index]["acc"]; exists {
				counter += acc
			}
		}
		if valid {
			break BreakEnd
		}
		// Rollback content
		editMap(posBuffer[len(posBuffer) - 1], parsedContent)
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



func editMap(index int, sliceMap []map[string]int) {
	if val, exists := sliceMap[index]["nop"]; exists && index > 0 {
		delete(sliceMap[index], "nop")
		sliceMap[index]["jmp"] = val
	} else if val, exists := sliceMap[index]["jmp"]; exists {
		delete(sliceMap[index], "jmp")
		sliceMap[index]["nop"] = val
	}
}