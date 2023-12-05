package main

import (
	"2023/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	buffer := common.ReadFile("2/input.txt")
	lines := strings.Split(string(buffer), "\n")
	lineCount := 0
	for lineIdx, line := range lines {
		rows := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line, ":", ","), ";", ","), ",", "")
		splitRow := strings.Split(rows, " ")
		tooMuch := false
		for index, rest := range splitRow {
			if "red" == rest {
				val, _ := strconv.Atoi(splitRow[index-1])
				if val > 12 {
					tooMuch = true
					break
				}
			} else if "blue" == rest {
				val, _ := strconv.Atoi(splitRow[index-1])
				if val > 14 {
					tooMuch = true
					break
				}
			} else if "green" == rest {
				val, _ := strconv.Atoi(splitRow[index-1])
				if val > 13 {
					tooMuch = true
					break
				}
			}
		}

		if !tooMuch {
			fmt.Println(lineIdx)
			lineCount += (lineIdx + 1)
		}
	}
	log.Printf("Number sum: %d", lineCount)
}
