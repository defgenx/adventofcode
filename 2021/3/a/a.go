package main

import (
	"2021/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("3/input.txt")
	var storage = [][]int{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	buffer := common.ReadBuffer(file)
	gammaRate := ""
	epsilonRate := ""
	for buffer.Scan() {
		splittedVals := strings.Split(buffer.Text(), "")
		for i, val := range splittedVals {
			bit, _ := strconv.Atoi(val)
			storage[i] = append(storage[i], bit)
		}
	}

	for _, val := range storage {
		res := most(val)
		if res {
			gammaRate = fmt.Sprintf("%s%d", gammaRate, 1)
			epsilonRate = fmt.Sprintf("%s%d", epsilonRate, 0)
		} else {
			gammaRate = fmt.Sprintf("%s%d", gammaRate, 0)
			epsilonRate = fmt.Sprintf("%s%d", epsilonRate, 1)
		}
	}
	gammaRateInt, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(epsilonRate, 2, 64)
	log.Printf("Val is: %d", gammaRateInt*epsilonRateInt)
}

func most(a []int) bool {
	one := 0
	zero := 0
	for _, val := range a {
		if val == 1 {
			one++
		} else {
			zero++
		}
	}
	return zero < one
}
