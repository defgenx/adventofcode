package lib

import (
	"fmt"
	"strconv"
)

func ReformatSliceToInt(input []string) []int {
	var inputInt = make([]int, len(input))
	for i, val := range input {
		inputInt[i], _ = strconv.Atoi(val)
	}
	return inputInt
}

func ReadOpcodes(input []int, noun, verb int) []int {
	input[1] = noun
	input[2] = verb
	for i := 0; i < len(input); i = i + 4 {
		action := input[i]
		posFirst := input[input[i+1]]
		posSec := input[input[i+2]]

		switch action {
		case 1:
			input[input[i+3]] = add(posFirst, posSec)
		case 2:
			input[input[i+3]] = multiply(posFirst, posSec)
		case 99:
			return input
		default:
			fmt.Print("Unknown opcode")
		}
	}
	return input
}

func add(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}
