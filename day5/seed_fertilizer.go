package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines := common.ReadFileLines("day5/input/input.txt")
	seeds, allMappings := buildData(lines)
	fmt.Println(seeds)
	fmt.Println(allMappings)
	locations := calculateLocations(seeds, allMappings)
	sort.Ints(locations)
	fmt.Println(locations[0])

	// part 2
	length := len(allMappings)
	for i := 0; i < length/2; i++ {
		allMappings[i], allMappings[length-1-i] = allMappings[length-1-i], allMappings[i]
	}

	for location := 0; location < 999999999; location++ {
		seed := calculateSeedGivenLocation(location, allMappings)
		if doesSeedExist(seed, seeds) {
			fmt.Println(location)
			break
		}
	}
}

func doesSeedExist(seed int, seeds []int) bool {
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := seeds[i+1]

		if start <= seed && start+end >= seed {
			return true
		}
	}
	return false
}

func buildData(lines []string) ([]int, [][]Mapping) {
	input := lines[0]

	// Extract numbers after "seeds:"
	fields := strings.Fields(strings.TrimPrefix(input, "seeds:"))

	// Convert strings to integers
	var seeds []int
	for _, field := range fields {
		num, _ := strconv.Atoi(field)
		seeds = append(seeds, num)
	}
	allMappings := make([][]Mapping, 0)
	mappings := make([]Mapping, 0)
	currentKey := -1
	for _, line := range lines[2:] {
		if line == "" {
			allMappings = append(allMappings, mappings)
		} else if unicode.IsLetter(rune(string(line[0])[0])) {
			mappings = make([]Mapping, 0)
			currentKey++
		} else {
			m := strings.Split(line, " ")
			drs, _ := strconv.Atoi(m[0])
			srs, _ := strconv.Atoi(m[1])
			rl, _ := strconv.Atoi(m[2])
			mappings = append(mappings, Mapping{drs, srs, rl})
		}
	}
	allMappings = append(allMappings, mappings)
	return seeds, allMappings
}

func calculateLocations(seeds []int, allMappings [][]Mapping) []int {
	locations := make([]int, 0)

	for _, seed := range seeds {
		currentNumber := seed
		for _, mappings := range allMappings {
			for _, mapping := range mappings {
				if currentNumber >= mapping.SourceRangeStart && currentNumber < mapping.SourceRangeStart+mapping.RangeLength {
					currentNumber = mapping.DestinationRangeStart + currentNumber - mapping.SourceRangeStart
					break
				}
			}
		}
		locations = append(locations, currentNumber)
	}
	return locations
}

func calculateSeedGivenLocation(location int, allMappings [][]Mapping) int {
	currentNumber := location
	for _, mappings := range allMappings {
		for _, mapping := range mappings {
			if currentNumber >= mapping.DestinationRangeStart && currentNumber < mapping.DestinationRangeStart+mapping.RangeLength {
				currentNumber = mapping.SourceRangeStart + currentNumber - mapping.DestinationRangeStart
				break
			}
		}
	}
	return currentNumber
}

type Mapping struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}
