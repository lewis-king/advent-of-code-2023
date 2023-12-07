package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileLines() []string {
	content, err := os.ReadFile("day2/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read lines from file
	return strings.Split(string(content), "\n")
}

var gameConfig = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGameStateAndCalcPossibleIds(lines []string) ([][]map[string]int, int) {
	allGameStates := make([][]map[string]int, 0)
	possibleIds := make(map[int]int)
	sumOfPossibleIds := 0

	for i, line := range lines {
		possible := true
		parts := strings.Split(strings.Split(line, ": ")[1], "; ")

		gameStates := make([]map[string]int, len(parts))
		for j, part := range parts {
			gameStates[j] = make(map[string]int)

			for _, pair := range strings.Split(part, ", ") {
				colorAndAmount := strings.Split(pair, " ")
				amount, _ := strconv.Atoi(colorAndAmount[0])
				color := colorAndAmount[1]

				gameStates[j][color] = amount

				if amount > gameConfig[color] {
					possible = false
				}
			}
		}

		if possible {
			possibleIds[i+1] = i + 1
			sumOfPossibleIds += i + 1
		}

		allGameStates = append(allGameStates, gameStates)
	}

	return allGameStates, sumOfPossibleIds
}

func calcMinSetOfCubes(gameStates [][]map[string]int) int {
	product := 0
	allMins := make([]map[string]int, 0)
	for _, gameState := range gameStates {
		mins := make(map[string]int)
		for _, individualGameState := range gameState {
			for colour, amount := range individualGameState {
				_, exists := mins[colour]
				if exists {
					if mins[colour] < amount {
						mins[colour] = amount
					}
				} else {
					mins[colour] = amount
				}
			}
		}
		allMins = append(allMins, mins)
	}
	for _, mins := range allMins {
		individualProduct := 0
		for _, amount := range mins {
			if individualProduct == 0 {
				individualProduct = amount
			} else {
				individualProduct *= amount
			}
		}
		product += individualProduct
	}
	return product
}

func main() {
	lines := readFileLines()
	gameStates, possibleIds := parseGameStateAndCalcPossibleIds(lines)
	fmt.Println(gameStates)
	fmt.Println(possibleIds)
	minimumSetOfCubes := calcMinSetOfCubes(gameStates)
	fmt.Println(minimumSetOfCubes)
}
