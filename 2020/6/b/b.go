package main

import (
	"2020/common"
	"log"
	"strings"
)

func main() {
	file := common.OpenInputFile("6/input.txt")
	buffer := common.ReadBuffer(file)
	mapByGroup := make(map[string]bool, 0)
	totalCount := 0
	firstGroupIter := true
	for buffer.Scan() {
		line := buffer.Text()
		if line == "" {
			firstGroupIter = true
			mapByGroup = make(map[string]bool, 0)
			continue
		}
		splittedLine := strings.Split(line, "")
		tmpSlice := make([]string, 0)
		for _, answer := range splittedLine {
			if firstGroupIter {
				mapByGroup[answer] = true
				totalCount++
			}
			tmpSlice = append(tmpSlice, answer)
		}
		for key := range mapByGroup {
			found := false
			for _, elt := range tmpSlice {
				if key == elt {
					found = true
				}
			}
			if !found {
				delete(mapByGroup, key)
				totalCount--
			}
		}
		firstGroupIter = false
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total count: %d", totalCount)
}