package main

import (
	"2020/common"
	"log"
	"strings"
)

func main() {
	file := common.OpenInputFile("3/input.txt")
	buffer := common.ReadBuffer(file)
	fileContent := make([][]string, 0)
	var slopes = []map[string]int{
		{
			"trees": 0,
			"right": 1,
			"down": 1,
		},
		{
			"trees": 0,
			"right": 3,
			"down": 1,
		},
		{
			"trees": 0,
			"right": 5,
			"down": 1,
		},
		{
			"trees": 0,
			"right": 7,
			"down": 1,
		},
		{
			"trees": 0,
			"right": 1,
			"down": 2,
		},
	}

	for buffer.Scan() {
		// Split each char of the line
		slice := strings.Split(buffer.Text(), "")
		fileContent = append(fileContent, slice)
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	var trees = 1
	for _, slope := range slopes {
		var right = slope["right"]
		for indexLine, lineSlice := range fileContent {
			if indexLine % slope["down"] != 0 || indexLine == 0 {
				continue
			}
			// If end of line is reached, return to the beginning
			if len(lineSlice)-1 < right {
				right = right - len(lineSlice)
			}
			if lineSlice[right] == "#" {
				slope["trees"]++
			}
			right += slope["right"]
		}
		log.Print(slopes)
		trees *= slope["trees"]
	}
	log.Printf("Total tree encountered: %d", trees)
}
