package main

import (
	"2023/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	buffer := common.ReadFile("2/input.txt")
	lines := strings.Split(string(buffer), "\n")
	lineCount := 0
	for _, line := range lines {
		rows := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line, ":", ","), ";", ","), ",", "")
		splitRow := strings.Split(rows, " ")
		red := 0
		blue := 0
		green := 0
		for index, rest := range splitRow {
			if "red" == rest {
				val, _ := strconv.Atoi(splitRow[index-1])
				if val > red {
					red = val
				}
			} else if "blue" == rest {
				val, _ := strconv.Atoi(splitRow[index-1])
				if val > blue {
					blue = val
				}
			} else if "green" == rest {
				val, _ := strconv.Atoi(splitRow[index-1])
				if val > green {
					green = val
				}
			}
		}

		lineCount += red * blue * green
	}
	log.Printf("Number sum: %d", lineCount)
}
