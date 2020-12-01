package main

import (
	"2019/2/lib"
	"2019/common"
	"fmt"
	"os"
)

func main() {
	fileContent := common.ReadFile("../input.txt")
	for verb := 0; verb <= 99; verb += 1 {
		for noun := 0; noun <= 99; noun += 1 {
			splitedContent := lib.ReformatSliceToInt(common.SplitStringToSlice(string(fileContent)))
			newState := lib.ReadOpcodes(splitedContent, noun, verb)
			if newState[0] == 19690720 {
				fmt.Printf("Val left at pos 0 for gravity assist: %d", 100*newState[1]+newState[2])
				os.Exit(0)
			}
		}
	}
}
