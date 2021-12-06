package main

import (
	"2021/common"
	"log"
	"strconv"
	"strings"
)

type Value struct {
	Number string
	Found  bool
}

type Grids [][][]*Value

var grids Grids

func main() {
	file := common.OpenInputFile("4/input.txt")
	var nbrs = make([]string, 0)
	var row = 0
	var gridNumber = 0
	var gridLineRow = 0
	var currentGrid [][]*Value
	buffer := common.ReadBuffer(file)
	for buffer.Scan() {
		rowtext := buffer.Text()
		if row == 0 {
			splittedVals := strings.Split(rowtext, ",")
			nbrs = splittedVals
			row++
			continue
		}
		if rowtext == "" {
			addGrid(gridNumber, currentGrid)
			if currentGrid != nil {
				gridNumber++
				currentGrid = make([][]*Value, 0)
			}
			row = 2
			gridLineRow = 0
			continue
		}
		splittedVals := strings.Split(rowtext, " ")
		var currentRow []*Value
		for _, val := range splittedVals {
			if val == "" {
				continue
			}
			currentRow = append(currentRow, &Value{
				Number: val,
				Found:  false,
			})
		}
		currentGrid = append(currentGrid, nil)
		currentGrid[gridLineRow] = currentRow
		gridLineRow++
	}

	addGrid(gridNumber, currentGrid)
	completedGridNumber, calledNumber := findFirstCompletedGrid(nbrs)
	sum := 0
	for _, col := range grids[completedGridNumber] {
		for _, val := range col {
			if !val.Found {
				intVal, _ := strconv.Atoi(val.Number)
				sum += intVal
			}
		}
	}
	calledNumberInt, _ := strconv.Atoi(calledNumber)
	log.Printf("Val is %d", calledNumberInt*sum)
}

func addGrid(pos int, grid [][]*Value) {
	if grid != nil {
		grids = append(grids, nil)
		grids[pos] = grid
	}
}

func findFirstCompletedGrid(numbers []string) (int, string) {
	var tmpRow = make([][5]int, len(grids))
	var tmpCol = make([][5]int, len(grids))
	for _, number := range numbers {
		for gridIndex, grid := range grids {
			for colIndex := 0; colIndex < 5; colIndex++ {
				for rowIndex := 0; rowIndex < 5; rowIndex++ {
					if grid[colIndex][rowIndex].Number == number {
						grid[colIndex][rowIndex].Found = true
						tmpRow[gridIndex][rowIndex]++
					}
					if grid[rowIndex][colIndex].Number == number {
						grid[rowIndex][colIndex].Found = true
						tmpCol[gridIndex][rowIndex]++
					}
					if tmpCol[gridIndex][rowIndex] == 5 || tmpRow[gridIndex][rowIndex] == 5 {
						return gridIndex, number
					}
				}

			}
		}

	}
	return 0, ""
}
