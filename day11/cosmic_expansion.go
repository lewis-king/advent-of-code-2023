package main

import (
	"advent-of-code-2023/common"
	"fmt"
)

// Part 1 set expansion to 1 for Part 2 set to 1000000
// 10
const expansion = 1_000_000 - 1

func main() {
	lines := common.ReadFileLines("day11/input/input.txt")
	universe := convertLinesTo2DArray(lines)
	// TODO: A more performant approach - rather than expanding the array to make the shortest path calculation simpler
	// I should just keep note of where the expansion rows and columns are and then consider them to add to the final answer during shortest path calculation
	// i.e in the calculation if the coordinate would have crossed that path increment by the expansion factor
	universe = expandRows(universe)
	universe = expandColumns(universe)

	galaxyCoordinates := identifyGalaxyCoordinates(universe)
	fmt.Println(galaxyCoordinates)

	distances := calculateShortestPathsBetweenGalaxies(galaxyCoordinates)
	totalShortestDistance := 0
	for _, value := range distances {
		totalShortestDistance += value
	}
	fmt.Println(totalShortestDistance)
}

func calculateShortestPathsBetweenGalaxies(galaxies []Coordinate) map[PairCoordinate]int {
	fmt.Println("Calculating shortest paths between galaxies...")
	distancesBetweenGalaxies := make(map[PairCoordinate]int)
	for i, galaxy := range galaxies {
		for j, secondGalaxy := range galaxies {
			if galaxy != secondGalaxy && j > i {
				xDistance := galaxy.x - secondGalaxy.x
				if xDistance < 0 {
					xDistance = -xDistance
				}
				yDistance := galaxy.y - secondGalaxy.y
				if yDistance < 0 {
					yDistance = -yDistance
				}
				distanceBetween := xDistance + yDistance
				fmt.Print(distanceBetween)
				distancesBetweenGalaxies[PairCoordinate{galaxy, secondGalaxy}] = distanceBetween
			}
		}
	}
	fmt.Println(distancesBetweenGalaxies)
	return distancesBetweenGalaxies
}

type Coordinate struct {
	x int
	y int
}

type PairCoordinate struct {
	firstCoordinate  Coordinate
	secondCoordinate Coordinate
}

func identifyGalaxyCoordinates(universe [][]string) []Coordinate {
	fmt.Println("Identifying galaxy coordinates...")
	galaxyCoordinates := make([]Coordinate, 0)
	for i, row := range universe {
		for j, char := range row {
			if string(char) == "#" {
				fmt.Print(Coordinate{i, j})
				galaxyCoordinates = append(galaxyCoordinates, Coordinate{i, j})
			}
		}
	}
	return galaxyCoordinates
}

func convertLinesTo2DArray(lines []string) [][]string {
	twoDArray := make([][]string, len(lines))
	for i, line := range lines {
		twoDArray[i] = make([]string, len(line))
		for j, char := range line {
			twoDArray[i][j] = string(char)
		}
	}
	return twoDArray
}

func expandRows(twoDArray [][]string) [][]string {
	fmt.Println("Expanding rows...")
	rowIndexes := make([]int, 0)
	i := 0
	for i < len(twoDArray) {
		fmt.Println(i)
		row := twoDArray[i]
		for j, char := range row {
			if char == "#" {
				i++
				break
			}
			if char == "." && j == len(row)-1 {
				rowIndexes = append(rowIndexes, i)
				x := 0
				for x < expansion {
					twoDArray = append(twoDArray[:i+1], twoDArray[i:]...)
					x++
				}
				i = i + expansion + 1
				break
			}
		}
	}
	//fmt.Println(twoDArray)
	return twoDArray
}

func expandColumns(twoDArray [][]string) [][]string {
	fmt.Println("Expanding columns...")
	columnIndexes := make([]int, 0)
	i := 0
	for i < len(twoDArray[0]) {
		fmt.Println(i)
		for j, _ := range twoDArray {
			element := twoDArray[j][i]
			if element == "#" {
				i++
				break
			}
			if element == "." && j == len(twoDArray)-1 {
				columnIndexes = append(columnIndexes, i)
				x := 0
				for x < expansion {
					for k, _ := range twoDArray {
						// insert "." at index i
						twoDArray[k] = append(twoDArray[k][:i+1], twoDArray[k][i:]...)
					}
					x++
				}
				i = i + 2 + (expansion - 1)
			}
		}
	}
	//fmt.Println(twoDArray)
	return twoDArray
}
