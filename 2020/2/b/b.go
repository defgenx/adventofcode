package main

import (
	"2020/common"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("2/input.txt")
	buffer := common.ReadBuffer(file)
	var rows = make([]map[string]interface{}, 0)
	for buffer.Scan() {
		var mapRow = make(map[string]interface{}, 0)
		// Thought to use buffer.Split() but easier with string lib
		slice := strings.Split(buffer.Text(), " ")
		rangeValues := strings.Split(slice[0], "-")
		min, _ := strconv.Atoi(rangeValues[0])
		mapRow["min"] = min
		max, _ := strconv.Atoi(rangeValues[1])
		mapRow["max"] = max
		mapRow["letter"] = strings.Split(slice[1], ":")[0]
		mapRow["pwd"] = slice[2]
		rows = append(rows, mapRow)
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	var totalValidPwd = 0
	for _, row := range rows {
		minLetter := string([]rune(row["pwd"].(string))[row["min"].(int)-1])
		maxLetter := string([]rune(row["pwd"].(string))[row["max"].(int)-1])
		if minLetter == row["letter"].(string) &&
			maxLetter == row["letter"].(string) {
			continue
		}
		if minLetter == row["letter"].(string) ||
			maxLetter == row["letter"].(string) {
			log.Print(row)
			totalValidPwd++
		}
	}
	log.Printf("Valid password in list: %d", totalValidPwd)
}
