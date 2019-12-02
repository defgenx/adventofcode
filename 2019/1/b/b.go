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
		//log.Print("Begin Row")
		mass, _ := strconv.Atoi(buffer.Text())
		fuelSum = fuelSum + recurseFuel(int(math.Floor(float64(mass) / 3)) - 2)
		//log.Print("End Row")
	}

	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total fuel needed: %d", fuelSum)
}

func recurseFuel(fuel int) int {
	//log.Print(fuel)
	if fuel < 3 && fuel >= 0 {
		return fuel
	}
	if fuel > 0 {
		return fuel + recurseFuel(int(math.Floor(float64(fuel) / 3)) - 2)
	}
	return 0
}