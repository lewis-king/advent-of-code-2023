package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day8/input/input.txt")

	directions := strings.Split(lines[0], "")

	lookupLines := lines[2:]
	lookups := make(map[string][]string)
	for _, line := range lookupLines {
		sanitizedOpenBrackets := strings.ReplaceAll(line, "(", "")
		sanitizedClosedBrackets := strings.ReplaceAll(sanitizedOpenBrackets, ")", "")
		lookup := strings.Split(sanitizedClosedBrackets, " = ")
		lookups[lookup[0]] = strings.Split(lookup[1], ", ")
	}

	fmt.Println(directions)
	fmt.Println(lookups)
	stepsToTake := findTheEnd(directions, lookups)
	fmt.Println(stepsToTake)
}

func findTheEnd(directions []string, lookups map[string][]string) int {
	directionsCounter := 0
	currentLookup := lookups["AAA"]
	steps := 0
	for directionsCounter < len(directions) {
		direction := directions[directionsCounter]
		newLookup := ""
		if direction == "L" {
			newLookup = currentLookup[0]
		} else if direction == "R" {
			newLookup = currentLookup[1]
		}
		steps++
		if newLookup == "ZZZ" {
			break
		}
		currentLookup = lookups[newLookup]
		if directionsCounter == len(directions)-1 {
			directionsCounter = 0
		} else {
			directionsCounter++
		}
	}
	return steps
}
