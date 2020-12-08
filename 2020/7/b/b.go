package main

import (
	"2020/common"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := common.OpenInputFile("7/input.txt")
	buffer := common.ReadBuffer(file)
	slice := make([]map[string]interface{}, 0)
	counter := 0
	for buffer.Scan() {
		line := buffer.Text()
		re := regexp.MustCompile(`^(?P<bagName>.+) bags contain (no other bag(s?)|(?P<affiliateBags>.+))\.$`)
		match := re.FindStringSubmatch(line)
		paramsMap := make(map[string]interface{})
		for i, name := range re.SubexpNames() {
			if i > 0 && i <= len(match) {
				if name == "affiliateBags" {
					bagNamesContain := make(map[string]int, 0)
					bags := strings.Split(match[i], ", ")
					for _, bag := range bags {
						reBag := regexp.MustCompile(`^(?P<bagNb>^\d+) (?P<bagName>(\w|\s)+) bag(s)?`)
						bagMatch := reBag.FindStringSubmatch(bag)
						for j, bagRegexKey := range reBag.SubexpNames() {
							if j > 0 && j <= len(bagMatch) {
								if bagRegexKey == "bagNb" {
									val, _ := strconv.Atoi(bagMatch[j])
									bagNamesContain[bagMatch[j+1]] = val
								}
							}
						}
					}
					paramsMap[name] = bagNamesContain
				} else {
					paramsMap[name] = match[i]
				}
			}
		}
		slice = append(slice, paramsMap)
	}
	for i := 0; i < len(slice); i++ {
		if slice[i]["bagName"].(string) == "shiny gold" {
			counter += find(slice[i]["bagName"].(string), slice)
			break
		}

	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total count: %v", counter)
}

func contains(bagName string, tmp []string) bool {
	for _, val := range tmp {
		if val == bagName {
			return true
		}
	}
	return false
}

func find(name string, slice []map[string]interface{}) int {
	counter:= 0
	for i := 0; i < len(slice); i++ {
		if slice[i]["bagName"] == name {
			affiliateBags := slice[i]["affiliateBags"].(map[string]int)
			if len(affiliateBags) == 0 {
				break
			}
			for affiliateBag, nbBag := range affiliateBags {
				counter += nbBag
				counter += nbBag * find(affiliateBag, slice)
			}
			break
		}
	}
	return counter
}