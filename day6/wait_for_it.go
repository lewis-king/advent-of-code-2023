package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day6/input/input.txt")

	var timeStrings []string
	var distanceStrings []string

	// Part 1
	digit := regexp.MustCompile(`\d+`)
	timeStrings = digit.FindAllString(lines[0], -1)
	distanceStrings = digit.FindAllString(lines[1], -1)
	total := calculateNumberOfWaysToWinTotal(timeStrings, distanceStrings)
	fmt.Println(total)

	// Part 2
	time := strings.Join(strings.Fields(strings.TrimPrefix(lines[0], "Time:")), "")
	distance := strings.Join(strings.Fields(strings.TrimPrefix(lines[1], "Distance:")), "")
	totalPart2 := calculateNumberOfWaysToWinTotal([]string{time}, []string{distance})
	fmt.Println(totalPart2)
}

func calculateNumberOfWaysToWinTotal(timeStrings []string, distanceStrings []string) int {
	total := 0
	for i := 0; i < len(timeStrings); i++ {
		raceTimeLimit, _ := strconv.Atoi(timeStrings[i])
		distanceRecord, _ := strconv.Atoi(distanceStrings[i])
		numberOfWinningDistances := 0
		for j := 1; j < raceTimeLimit; j++ {
			totalDistanceWillBe := j * (raceTimeLimit - j)
			if totalDistanceWillBe > distanceRecord {
				numberOfWinningDistances++
			}
		}
		if total == 0 {
			total = numberOfWinningDistances
		} else {
			total *= numberOfWinningDistances
		}
	}
	return total
}
