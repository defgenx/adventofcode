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
	var values []int
	for _, val := range splittedVals {
		x, _ := strconv.Atoi(val)
		values = append(values, x)
	}
	for i := 0; i < 80; i++ {
		for j, val := range values {
			if val == 0 {
				values[j] = 6
				values = append(values, 8)
			} else {
				values[j] = val - 1
			}
		}
	}
	log.Printf("Fish count: %d", len(values))
}
