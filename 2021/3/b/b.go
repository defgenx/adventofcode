package main

import (
	"2021/common"
	"github.com/juliangruber/go-intersect"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("3/input.txt")
	var storage = [][]int{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	buffer := common.ReadBuffer(file)
	var rawline []string
	for buffer.Scan() {
		row := buffer.Text()
		rawline = append(rawline, row)
		splittedVals := strings.Split(row, "")
		for i, val := range splittedVals {
			bit, _ := strconv.Atoi(val)
			storage[i] = append(storage[i], bit)
		}
	}
	oxygen, co2 := oxygenAndCo2(storage)
	oxygenVal, _ := strconv.ParseInt(rawline[oxygen[0]], 2, 64)
	co2Val, _ := strconv.ParseInt(rawline[co2[0]], 2, 64)
	log.Printf("Val is: %d", oxygenVal*co2Val)
}

func most(a []int, indices []int) ([]int, []int) {
	one := make([]int, 0)
	zero := make([]int, 0)
	for _, i := range indices {
		if a[i] == 1 {
			one = append(one, i)
		} else {
			zero = append(zero, i)
		}
	}
	return one, zero
}

func oxygenAndCo2(a [][]int) (newOxygen, newCo2 []int) {
	for i := 0; i < len(a[0]); i++ {
		newOxygen = append(newOxygen, i)
		newCo2 = append(newCo2, i)
	}
	for _, val := range a {
		if len(newOxygen) > 1 {
			one, zero := most(val, newOxygen)
			if len(one) >= len(zero) {
				newOxygen = toInt(intersect.Simple(newOxygen, one))
			} else {
				newOxygen = toInt(intersect.Simple(newOxygen, zero))
			}
		}
		if len(newCo2) > 1 {
			one, zero := most(val, newCo2)
			if len(one) >= len(zero) {
				newCo2 = toInt(intersect.Simple(newCo2, zero))
			} else {
				newCo2 = toInt(intersect.Simple(newCo2, one))
			}
		}
	}
	return
}

func toInt(a []interface{}) (b []int) {
	for _, val := range a {
		b = append(b, val.(int))
	}
	return
}
