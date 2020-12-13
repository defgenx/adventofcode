package main

import (
	"2020/common"
	"log"
	"strings"
)

func main() {
	file := common.OpenInputFile("11/input.txt")
	buffer := common.ReadBuffer(file)
	bufferMap := make([][]string, 0, 0)
	first := true
	tmp := make([]string, 0)
	for buffer.Scan() {
		line := buffer.Text()
		splittedLine := make([]string, 0)
		splittedLine = strings.Split("." + line + ".", "")
		if first {
			for i := 0; i < len(splittedLine); i++ {
				tmp = append(tmp, ".")
			}
			bufferMap = append(bufferMap, tmp)
			first = false
		}
		splittedLine = append(splittedLine, ".")
		bufferMap = append(bufferMap, splittedLine)
	}

	bufferMap = append(bufferMap, tmp)
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	for {
		move := 0
		for i := 1; i < len(bufferMap) -2; i++ {
			for j := 1; j < len(bufferMap[i]) -2; j++ {
				if bufferMap[i][j] == "L" {
					if bufferMap[i+1][j] == "L" || bufferMap[i+1][j] == "." ||
						bufferMap[i-1][j] == "L" || bufferMap[i-1][j] == "." ||
						bufferMap[i][j+1] == "L" || bufferMap[i][j+1] == "." ||
						bufferMap[i][j-1] == "L" || bufferMap[i][j-1] == "." ||
						bufferMap[i-1][j-1] == "L" || bufferMap[i-1][j-1] == "." ||
						bufferMap[i-1][j+1] == "L" || bufferMap[i-1][j+1] == "." ||
						bufferMap[i+1][j-1] == "L" || bufferMap[i+1][j-1] == "." ||
						bufferMap[i+1][j+1] == "L" || bufferMap[i+1][j+1] == "." {
						bufferMap[i][j] = "#"
						move++
					}
				}
			}

			log.Print("====================")
		}
		for i := 1; i < len(bufferMap) -2; i++ {
			for j := 1; j < len(bufferMap[i]) -2; j++ {
				if bufferMap[i][j] == "#" {
					if isFreeable(bufferMap, i, j) {
						bufferMap[i][j] = "L"
						move++
					}
				}
			}
		}
		if move == 0 {
			break
		}
	}
	occSeats := 0
	for i := 1; i < len(bufferMap); i++ {
		log.Print(bufferMap[i])
		for j := 1; j < len(bufferMap[i]); j++ {
			if bufferMap[i][j] == "#" {
				occSeats++
			}
		}
	}

	log.Printf("Occupied seats: %v", occSeats)
}

func isFreeable(bufferMap [][]string, i, j int) bool{
	count := 0
	if bufferMap[i+1][j] == "#" {
		count ++
	}
	if bufferMap[i+1][j+1] == "#" {
		count ++
	}
	if bufferMap[i][j+1] == "#" {
		count ++
	}
	if bufferMap[i-1][j] == "#" {
		count ++
	}
	if bufferMap[i-1][j-1] == "#" {
		count ++
	}
	if bufferMap[i][j-1] == "#" {
		count ++
	}
	if bufferMap[i-1][j+1] == "#" {
		count ++
	}
	if bufferMap[i+1][j-1] == "#" {
		count ++
	}
	return count >= 4
}