package main

import (
	"2020/common"
	"log"
	"strings"
)

func main() {
	file := common.OpenInputFile("6/input.txt")
	buffer := common.ReadBuffer(file)
	sliceByGroup := make([]string, 0)
	totalCount := 0
	for buffer.Scan() {
		line := buffer.Text()
		if line == "" {
			sliceByGroup = make([]string, 0)
			continue
		}
		splittedLine := strings.Split(line, "")
		for _, answer := range splittedLine {
			if !containsAnswer(answer, sliceByGroup) {
				sliceByGroup = append(sliceByGroup, answer)
				totalCount++
			}
		}
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total count: %d", totalCount)
}

func containsAnswer(answer string, slice []string) bool {
	for _, val := range slice {
		if val == answer {
			return true
		}
	}
	return false
}