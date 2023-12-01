package main

import (
	"2023/common"
	"log"
	"strconv"
	"strings"
)

var (
	strNumber = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
)

func main() {
	buffer := common.ReadFile("1/input.txt")
	var sum = 0
	lines := strings.Split(string(buffer), "\n")
	for _, line := range lines {
		var tmpChar []string
		var num []string
		for _, char := range line {
			tmpChar = append(tmpChar, string(char))
			for k, v := range strNumber {
				if strings.Contains(strings.Join(tmpChar, ""), k) {
					num = append(num, strconv.Itoa(v))
					tmpChar = []string{string(char)}
				}
			}
		}
		integer, _ := strconv.Atoi(num[0] + num[len(num)-1])
		sum += integer
	}
	log.Printf("Number sum: %d", sum)
}
