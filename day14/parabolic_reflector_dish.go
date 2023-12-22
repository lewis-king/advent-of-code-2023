package main

import (
	"advent-of-code-2023/common"
)

func main() {
	lines := common.ReadFileLines("day14/input/input.txt")
	rockPlane := buildRockPlane(lines)
	transposedRockPlane := transposeRockPlane(rockPlane)
	common.PrettyPrint2DArray(transposedRockPlane)
	load := countLoad(transposedRockPlane)
	println(load)
}

func countLoad(rockPlane [][]string) int {
	totalLoad := 0
	for i := 0; i < len(rockPlane); i++ {
		for _, char := range rockPlane[i] {
			if char == "O" {
				totalLoad += len(rockPlane) - i
			}
		}
	}
	return totalLoad
}

func transposeRockPlane(rockPlane [][]string) [][]string {
	for i := 0; i < len(rockPlane); i++ {
		for j := 0; j < len(rockPlane[i]); j++ {
			selectedObject := rockPlane[i][j]
			if selectedObject == "O" {
				// shift 0 up
				for k := i; k > 0; k-- {
					if rockPlane[k-1][j] == "." {
						rockPlane[k-1][j] = "O"
						rockPlane[k][j] = "."
					} else {
						break
					}
				}
			}
		}
	}
	return rockPlane
}

func buildRockPlane(lines []string) [][]string {
	rockPlane := make([][]string, len(lines))
	for i, line := range lines {
		rockPlane[i] = make([]string, len(line))
		for j, char := range line {
			rockPlane[i][j] = string(char)
		}
	}
	return rockPlane
}
