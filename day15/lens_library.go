package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day15/input/input.txt")
	initSequences := strings.Split(lines[0], ",")
	total := 0
	for _, initSequence := range initSequences {
		total += runHashAlgorithm(initSequence)
	}
	fmt.Println(total)
}

func runHashAlgorithm(sequence string) int {
	elements := strings.Split(sequence, "")
	currentValue := 0
	for _, element := range elements {
		// get ASCII code of current element
		asciiCode := int(element[0])
		currentValue += asciiCode
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}
