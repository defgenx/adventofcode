package main

import (
	"2019/common"
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	file := common.OpenInputFile("../input.txt")
	buffer := common.ReadBuffer(file)
	var fuelSum = 0
	for buffer.Scan() {
		mass, _ := strconv.Atoi(buffer.Text())
		fuelSum = fuelSum + (int(math.Floor(float64(mass) / 3)) - 2)
	}

	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total fuel needed: %d", fuelSum)
}