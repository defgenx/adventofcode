package main

import (
	"2021/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	content := common.ReadFile("6/input.txt")
	splittedVals := strings.Split(string(content), ",")
	var values [9]int
	for _, val := range splittedVals {
		x, _ := strconv.Atoi(val)
		values[x]++
	}
	for i := 0; i < 256; i++ {
		for j, counter := range values {
			if j == 0 {
				values[6] += counter
				values[8] += counter
				values[j] -= counter
			} else {
				values[j-1] += counter
				values[j] -= counter
			}
		}
	}
	var count = 0
	for _, val := range values {
		count += val
	}
	log.Printf("Fish count: %d", count)
}
