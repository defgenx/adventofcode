package main

import (
	"2023/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	buffer := common.ReadFile("1/input.txt")
	var sum = 0
	lines := strings.Split(string(buffer), "\n")
	for _, line := range lines {
		var buff []string
		for _, char := range line {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			buff = append(buff, string(char))
		}
		concatBuff := buff[0] + buff[len(buff)-1]
		integer, _ := strconv.Atoi(concatBuff)
		sum += integer
	}
	log.Printf("Number sum: %d", sum)
}
