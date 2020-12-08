package main

import (
	"2020/common"
	"log"
	"regexp"
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
					bagNamesContain := make([]string, 0)
					bags := strings.Split(match[i], ", ")
					for _, bag := range bags {
						reBag := regexp.MustCompile(`^\d+ (?P<bagName>(\w|\s)+) bag(s)?`)
						bagMatch := reBag.FindStringSubmatch(bag)
						for j, bagName := range re.SubexpNames() {
							if j > 0 && j <= len(bagMatch) {
								if bagName == "bagName" {
									bagNamesContain = append(bagNamesContain, bagMatch[j])
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
			continue
		}
		if find(slice[i]["bagName"].(string), slice) {
			counter++
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

func find(name string, slice []map[string]interface{}) bool {
	if name == "shiny gold" {
		return true
	}
	for i := 0; i < len(slice); i++ {
		if slice[i]["bagName"] == name {
			affiliateBags := slice[i]["affiliateBags"].([]string)
			if len(affiliateBags) == 0 {
				break
			}
			if contains("shiny gold", affiliateBags) {
				return true
			}
			for _, affiliateBag := range affiliateBags {
				if find(affiliateBag, slice) {
					return true
				}
			}
			break
		}
	}
	return false
}