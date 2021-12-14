package main

import (
	"2021/common"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	content := common.ReadFile("7/input.txt")
	splittedVals := strings.Split(string(content), ",")
	var values []int
	for _, val := range splittedVals {
		x, _ := strconv.Atoi(val)
		values = append(values, x)
	}
	min, max := common.MinMax(values)
	minFuel := -1
	validPos := min
	for i := min; i <= max; i++ {
		tmpFuel := 0
		for _, val := range values {
			dist := int(math.Abs(float64(val) - float64(i)))
			tmpFuel += (dist * (dist + 1)) / 2
		}
		if i == min {
			minFuel = tmpFuel
		}
		if minFuel > tmpFuel {
			minFuel = tmpFuel
			validPos = i
		}
	}
	log.Printf("Min fuel: %d at pos %d", minFuel, validPos)
}
