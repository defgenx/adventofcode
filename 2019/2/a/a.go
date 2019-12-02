package main

import (
	"2019/2/lib"
	"2019/common"
	"fmt"
)

func main() {
	fileContent := common.ReadFile("../input.txt")
	splitedContent := lib.ReformatSliceToInt(common.SplitStringToSlice(string(fileContent)))
	newState := lib.ReadOpcodes(splitedContent, 12, 2)
	fmt.Printf("Val left at pos 0: %d", newState[0])
}