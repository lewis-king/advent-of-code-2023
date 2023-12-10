package main

import (
	"advent-of-code-2023/common"
	"fmt"
)

func main() {
	lines := common.ReadFileLines("day10/input/input.txt")

	var allCoordinates [][]string
	startingCoordinate := Coordinate{}
	// find starting point
	for y, line := range lines {
		var rowCoordinates []string
		for x, e := range line {
			char := string(e)
			rowCoordinates = append(rowCoordinates, char)
			if char == "S" {
				startingCoordinate = Coordinate{x, y}
			}
		}
		allCoordinates = append(allCoordinates, rowCoordinates)
	}
	fmt.Println(allCoordinates)
	fmt.Println(startingCoordinate)
	followPipes(startingCoordinate, allCoordinates)
}

func followPipes(startingCoordinate Coordinate, pipes [][]string) {

	currentCoordinate := startingCoordinate
	visitedPipes := make(map[Coordinate]bool)
	visitedPipes[startingCoordinate] = true
	steps := 0
	for (currentCoordinate != startingCoordinate && steps != 0) || steps == 0 {
		currentCoordinate = move(startingCoordinate, currentCoordinate, pipes, visitedPipes)
		fmt.Println(currentCoordinate)
		visitedPipes[currentCoordinate] = true
		steps++
	}
	fmt.Println(steps)
	furthestPointAway := steps / 2
	fmt.Println(furthestPointAway)

}

func move(startingCoordinate Coordinate, currentCoordinate Coordinate, pipes [][]string, visitedPipes map[Coordinate]bool) Coordinate {
	currentPipe := pipes[currentCoordinate.y][currentCoordinate.x]
	// attempt move left
	if currentCoordinate.x > 0 && (!visitedPipes[Coordinate{currentCoordinate.x - 1, currentCoordinate.y}] || startingCoordinate == Coordinate{currentCoordinate.x - 1, currentCoordinate.y}) && isValidLeftMove(currentPipe, pipes[currentCoordinate.y][currentCoordinate.x-1]) {
		currentCoordinate.x--
	} else if currentCoordinate.y < len(pipes)-1 && (!visitedPipes[Coordinate{currentCoordinate.x, currentCoordinate.y + 1}] || startingCoordinate == Coordinate{currentCoordinate.x, currentCoordinate.y + 1}) && isValidDownMove(currentPipe, pipes[currentCoordinate.y+1][currentCoordinate.x]) {
		currentCoordinate.y++
	} else if currentCoordinate.x < len(pipes[currentCoordinate.y])-1 && (!visitedPipes[Coordinate{currentCoordinate.x + 1, currentCoordinate.y}] || startingCoordinate == Coordinate{currentCoordinate.x + 1, currentCoordinate.y}) && isValidRightMove(currentPipe, pipes[currentCoordinate.y][currentCoordinate.x+1]) {
		currentCoordinate.x++
	} else if currentCoordinate.y > 0 && (!visitedPipes[Coordinate{currentCoordinate.x, currentCoordinate.y - 1}] || startingCoordinate == Coordinate{currentCoordinate.x, currentCoordinate.y - 1}) && isValidUpMove(currentPipe, pipes[currentCoordinate.y-1][currentCoordinate.x]) {
		currentCoordinate.y--
	}
	return currentCoordinate
}

func isValidLeftMove(currentPipe string, newPipe string) bool {
	return (currentPipe == "S" || currentPipe == "-" || currentPipe == "J" || currentPipe == "7") && (newPipe == "-" || newPipe == "L" || newPipe == "F" || newPipe == "S")
}

func isValidDownMove(currentPipe string, newPipe string) bool {
	return (currentPipe == "S" || currentPipe == "|" || currentPipe == "F" || currentPipe == "7") && (newPipe == "|" || newPipe == "L" || newPipe == "J" || newPipe == "S")
}

func isValidRightMove(currentPipe string, newPipe string) bool {
	return (currentPipe == "S" || currentPipe == "-" || currentPipe == "F" || currentPipe == "L") && (newPipe == "-" || newPipe == "J" || newPipe == "7" || newPipe == "S")
}

func isValidUpMove(currentPipe string, newPipe string) bool {
	return (currentPipe == "S" || currentPipe == "|" || currentPipe == "J" || currentPipe == "L") && (newPipe == "|" || newPipe == "F" || newPipe == "7" || newPipe == "S")
}

type Coordinate struct {
	x int
	y int
}
