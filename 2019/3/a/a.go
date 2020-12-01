package main

import (
	"2019/common"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var lines [][]string
	var grid = make(map[int]map[int]string)
	var dist = 0
	file := common.OpenInputFile("../input.txt")
	buffer := common.ReadBuffer(file)
	for buffer.Scan() {
		lines = append(lines, strings.Split(buffer.Text(), ","))

	}
	direction := regexp.MustCompile(`[R|L|U|D]`)
	position := regexp.MustCompile(`[0-9]+`)

	var x = 0
	var y = 0
	for _, line := range lines {
		for _, val := range line {
			dir := direction.FindAllString(val, 1)[0]
			pos, _ := strconv.Atoi(position.FindAllString(val, 1)[0])
		}
	}

	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(dist + 1)
}

func move(dir string, pos int) (coorX int, coorY int) {
	switch dir {
	case "R":
		coorX = x + 1
		coorY = y
	case "L":
		coorX = x - 1
		coorY = y
	case "U":
		coorX = x
		coorY = y + 1
	case "D":
		coorX = x
		coorY = y - 1
	default:
		log.Fatalf("direction [%s] not implemented", dir)
	}
	return
}
