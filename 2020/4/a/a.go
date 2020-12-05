package main

import (
	"2020/common"
	"log"
	"strings"
)

var mandatoryKeys = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var notMandatoryKeys = []string{
	"cid",
}

func main() {
	file := common.OpenInputFile("4/input.txt")
	buffer := common.ReadBuffer(file)
	var validPassports = 0
	passportKeys := make([]string, 0)
	for buffer.Scan() {
		line := string(buffer.Bytes())
		if line == "" {
			if isPassportValid(passportKeys) {
				validPassports++
			}
			passportKeys = make([]string, 0)
			continue
		}
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			key := strings.Split(token,":")
			passportKeys = append(passportKeys, key[0])
		}
	}
	// Handle EOF
	if isPassportValid(passportKeys) {
		validPassports++
	}

	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total valid passports: %d", validPassports)
}

func isPassportValid(keys []string) bool {
	count := 0
	for _, key := range keys {
		for _, mandatory := range mandatoryKeys {
			if key == mandatory {
				count++
			}
		}
	}
	if count == len(mandatoryKeys) {
		return true
	}
	return false
}
