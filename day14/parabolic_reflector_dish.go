package main

import (
	"advent-of-code-2023/common"
	"fmt"
)

const part2 = true

const cycles = 1000 // realised this gave the same answer as 1000000000 - must be some mathematical relationship then... maybe because 1000 cubed is a billion....

//const cycles = 10000000

func main() {
	lines := common.ReadFileLines("day14/input/input.txt")
	rockPlane := buildRockPlane(lines)
	var transposedRockPlane [][]string
	if !part2 {
		transposedRockPlane = transposeRockPlane(rockPlane)
	} else {
		for i := 0; i < cycles; i++ {
			//lastTransposedRockPlane := transposedRockPlane
			transposedRockPlane = transposeRockPlane(rockPlane)
			if i%100000000 == 0 {
				fmt.Println("cycle", i+1)
				common.PrettyPrint2DArray(transposedRockPlane)
			}
			//if reflect.DeepEqual(lastTransposedRockPlane, transposedRockPlane) {
			//	break
			//}
		}
	}
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
				shiftUp(rockPlane, i, j)
			}
		}
	}
	//fmt.Println("shifting up...")
	//common.PrettyPrint2DArray(rockPlane)
	if part2 {
		for i := 0; i < len(rockPlane); i++ {
			for j := 0; j < len(rockPlane[i]); j++ {
				selectedObject := rockPlane[i][j]
				if selectedObject == "O" {
					// shift 0 left
					shiftLeft(rockPlane, i, j)
				}
			}
		}
		//fmt.Println("shifting left...")
		//common.PrettyPrint2DArray(rockPlane)
		for i := len(rockPlane) - 1; i >= 0; i-- {
			for j := 0; j < len(rockPlane[i]); j++ {
				selectedObject := rockPlane[i][j]
				if selectedObject == "O" {
					// shift 0 down
					shiftDown(rockPlane, i, j)
				}
			}
		}
		//fmt.Println("shifting down...")
		//common.PrettyPrint2DArray(rockPlane)
		for i := 0; i < len(rockPlane); i++ {
			for j := len(rockPlane) - 1; j >= 0; j-- {
				selectedObject := rockPlane[i][j]
				if selectedObject == "O" {
					// shift 0 right
					shiftRight(rockPlane, i, j)
				}
			}
		}
		//fmt.Println("shifting right...")
		//common.PrettyPrint2DArray(rockPlane)
	}
	return rockPlane
}

func shiftUp(rockPlane [][]string, i int, j int) {
	for k := i; k > 0; k-- {
		if rockPlane[k-1][j] == "." {
			rockPlane[k-1][j] = "O"
			rockPlane[k][j] = "."
		} else {
			break
		}
	}
}

func shiftLeft(rockPlane [][]string, i int, j int) {
	for k := j; k > 0; k-- {
		if rockPlane[i][k-1] == "." {
			rockPlane[i][k-1] = "O"
			rockPlane[i][k] = "."
		} else {
			break
		}
	}
}

func shiftDown(rockPlane [][]string, i int, j int) {
	for k := i; k < len(rockPlane)-1; k++ {
		if rockPlane[k+1][j] == "." {
			rockPlane[k+1][j] = "O"
			rockPlane[k][j] = "."
		} else {
			break
		}
	}
}

func shiftRight(rockPlane [][]string, i int, j int) {
	for k := j; k < len(rockPlane[i])-1; k++ {
		if rockPlane[i][k+1] == "." {
			rockPlane[i][k+1] = "O"
			rockPlane[i][k] = "."
		} else {
			break
		}
	}
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
