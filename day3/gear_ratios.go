package main

import (
	common "advent-of-code-2023/common"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day3/input/input.txt")
	sumPartNumbers(lines)
}

type Coordinate struct {
	x int
	y int
}

var symbolRegex = regexp.MustCompile(`[^a-zA-Z0-9.]`)
var wildcardRegex = regexp.MustCompile(`[*]`)

var adjacentCoordinates = []Coordinate{
	{-1, -1}, // top left diagonal
	{-1, 0},  // top
	{-1, 1},  // top right diagonal
	{0, -1},  // left
	{0, 1},   // right
	{1, -1},  // bottom left diagonal
	{1, 0},   // bottom
	{1, 1},   // bottom right diagonal
}

func sumPartNumbers(lines []string) int {
	sum := 0
	coordinateToDigit := make(map[Coordinate]int)
	partNumbersAndCoordinates := make(map[Coordinate]int)
	gearRatioCoordinates := make(map[Coordinate][]int)
	// coordinates for symbols
	for i, line := range lines {
		for j, gear := range strings.Split(line, "") {
			if v, err := strconv.Atoi(gear); err == nil {
				coordinateToDigit[Coordinate{i, j}] = v
				// is adjacent to a symbol
				if condition, _ := isAdjacentToSymbol(Coordinate{i, j}, lines, symbolRegex); condition {
					var constructedFullNumber, adjacentCoordinate = fullNumber(Coordinate{i, j}, lines)
					partNumbersAndCoordinates[adjacentCoordinate] = constructedFullNumber
				}
				if condition, symbolCoordinate := isAdjacentToSymbol(Coordinate{i, j}, lines, wildcardRegex); condition {
					var constructedFullNumber, _ = fullNumber(Coordinate{i, j}, lines)
					if len(gearRatioCoordinates[symbolCoordinate]) > 0 && slices.Contains(gearRatioCoordinates[symbolCoordinate], constructedFullNumber) {
						continue
					}
					gearRatioCoordinates[symbolCoordinate] = append(gearRatioCoordinates[symbolCoordinate], constructedFullNumber)
				}
			}
		}
	}
	//fmt.Println(partNumbersAndCoordinates)
	for _, value := range partNumbersAndCoordinates {
		sum += value
	}
	//fmt.Println(sum)

	fmt.Println(gearRatioCoordinates)
	var gearRatioProduct int64 = 0
	for _, value := range gearRatioCoordinates {
		if len(value) == 2 {
			gearRatioProduct += int64(value[0]) * int64(value[1])
		}
	}
	fmt.Println(gearRatioProduct)
	return sum
}

func fullNumber(coordinate Coordinate, lines []string) (int, Coordinate) {
	fullNumber := 0
	startNumberCoordinate := coordinate
	var x1 = coordinate.y
	for x1 >= 0 {
		value := strings.Split(lines[coordinate.x], "")[x1]
		if _, err := strconv.Atoi(value); err == nil {
			startNumberCoordinate = Coordinate{coordinate.x, x1}
		} else {
			break
		}
		x1--
	}
	x3 := startNumberCoordinate.y
	fullNumberStr := ""
	for x3 < len(lines[0]) {
		value := strings.Split(lines[coordinate.x], "")[x3]
		if _, err := strconv.Atoi(value); err != nil {
			break
		}
		fullNumberStr += value
		x3++
	}
	if fullNum, err := strconv.Atoi(fullNumberStr); err == nil {
		return fullNum, Coordinate{coordinate.x, x3}
	}
	return fullNumber, Coordinate{coordinate.x, x3}
}

func isAdjacentToSymbol(coordinate Coordinate, lines []string, regex *regexp.Regexp) (bool, Coordinate) {
	for _, adjacentCoordinate := range adjacentCoordinates {
		var translateX = coordinate.x + adjacentCoordinate.x
		var translateY = coordinate.y + adjacentCoordinate.y
		if (translateX < 0 || translateX >= len(lines)) || (translateY < 0 || translateY >= len(lines)) {
			continue
		}
		value := strings.Split(lines[translateX], "")[translateY]
		if regex.MatchString(value) {
			return true, Coordinate{translateX, translateY}
		}
	}
	return false, Coordinate{-1, -1}

}
