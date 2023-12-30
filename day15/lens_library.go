package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day15/input/input.txt")
	initSequences := strings.Split(lines[0], ",")

	// Part 1
	total := 0
	for _, initSequence := range initSequences {
		total += runHashAlgorithm(initSequence)
	}
	fmt.Println(total)

	// Part 2
	boxesAndLenses := populateBoxes(initSequences)
	focusingPower := calculateFocusingPower(boxesAndLenses)
	fmt.Println(focusingPower)
}

func calculateFocusingPower(boxesAndLenses map[int][]LabelAndLens) int {
	totalFocusingPower := 0
	for box, lenses := range boxesAndLenses {
		for slot, focalLength := range lenses {
			focusingPower := (box + 1) * (slot + 1) * focalLength.lens
			fmt.Println(focusingPower)
			totalFocusingPower += focusingPower
		}
	}
	return totalFocusingPower
}

func populateBoxes(initSequences []string) map[int][]LabelAndLens {
	boxesAndLenses := make(map[int][]LabelAndLens)
	for _, initSequence := range initSequences {
		if strings.Contains(initSequence, "=") {
			label := strings.Split(initSequence, "=")[0]
			box := runHashAlgorithm(label)
			lenses := boxesAndLenses[box]
			exists := false
			for index, lens := range lenses {
				if lens.label == label {
					focal, _ := strconv.Atoi(strings.Split(initSequence, "=")[1])
					newLens := LabelAndLens{label, focal}
					lenses[index] = newLens
					exists = true
				}
			}
			if !exists {
				focal, _ := strconv.Atoi(strings.Split(initSequence, "=")[1])
				newLens := LabelAndLens{label, focal}
				boxesAndLenses[box] = append(lenses, newLens)
			}
		} else if strings.Contains(initSequence, "-") {
			label := strings.Split(initSequence, "-")[0]
			box := runHashAlgorithm(label)
			lenses := boxesAndLenses[box]
			for index, lens := range lenses {
				if lens.label == label {
					boxesAndLenses[box] = append(lenses[:index], lenses[index+1:]...)
				}
			}
		}
		fmt.Println(boxesAndLenses)
	}
	return boxesAndLenses
}

type LabelAndLens struct {
	label string
	lens  int
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
