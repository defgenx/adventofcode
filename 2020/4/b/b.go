package main

import (
	"2020/common"
	"log"
	"regexp"
	"strconv"
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
	passportKeyVals := make(map[string]string, 0)
	for buffer.Scan() {
		line := string(buffer.Bytes())
		if line == "" {
			if isPassportValid(passportKeyVals) {
				validPassports++
			}
			passportKeyVals = make(map[string]string, 0)
			continue
		}
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			key := strings.Split(token,":")
			passportKeyVals[key[0]] = key[1]
		}
	}
	// Handle EOF
	if isPassportValid(passportKeyVals) {
		validPassports++
	}
	if err := buffer.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total valid passports: %d", validPassports)
}

func isPassportValid(passport map[string]string) bool {
	if !passportMandatoryPresent(passport) {
		return false
	}
	for key, val := range passport {
		if !passportKeyContentValid(key, val) {
			return false
		}
	}
	return true
}

func passportMandatoryPresent(keys map[string]string) bool {
	count := 0
	for key := range keys {
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

func passportKeyContentValid(key, content string) bool {
	switch key {
	case "byr":
		val, _ := strconv.Atoi(content)
		if val >= 1920 && val <= 2002 {
			return true
		}
		return false
	case "iyr":
		val, _ := strconv.Atoi(content)
		if val >= 2010 && val <= 2020 {
			return true
		}
		return false
	case "eyr":
		val, _ := strconv.Atoi(content)
		if val >= 2020 && val <= 2030 {
			return true
		}
		return false
	case "hgt":
		reCm := regexp.MustCompile(`.+cm$`)
		reIn := regexp.MustCompile(`.+in$`)
		if reCm.MatchString(content) {
			res := strings.Split(content, "cm")
			intVal, _ := strconv.Atoi(res[0])
			if intVal >= 150 && intVal <= 193 {
				return true
			}
			return false
		} else if reIn.MatchString(content) {
			res := strings.Split(content, "in")
			intVal, _ := strconv.Atoi(res[0])
			if intVal >= 59 && intVal <= 76 {
				return true
			}
			return false
		}
		return false
	case "hcl":
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		if re.MatchString(content) {
			return true
		}
		return false
	case "ecl":
		colors := []string {
			"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
		}
		for _, color := range colors {
			if content == color {
				return true
			}
		}
		return false
	case "pid":
		res := strings.Split(content, "")
		if len(res) == 9 {
			return true
		}
		return false
	case "cid":
		return true
	default:
		log.Print("Unknown key")
		return false
	}
}
