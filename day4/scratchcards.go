package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day4/input/input.txt")
	allNumbers, allWinningNumbers, allScratchcards := buildWinningsNumbersAndNumbers(lines)

	// part 1
	totalPrize := calculatePart1Prize(allNumbers, allWinningNumbers)
	fmt.Println(totalPrize)

	// part 2
	totalScratchCards := calculatePart2Prize(allScratchcards)
	fmt.Println(totalScratchCards)
}

func calculatePart2Prize(allScratchcards []string) int {
	return findScratchcardsWon(allScratchcards)
}

func findScratchcardsWon(allScratchcards []string) int {
	totalScratchcardsWon := 0
	remainingScratchcards := make([]string, 0)
	remainingScratchcards = append(remainingScratchcards, allScratchcards[0])
	index := 0
	hasSeen := make(map[int]bool)

	for len(remainingScratchcards) > 0 {
		scratchcardIndex, scratchcardsWon := findWinningNumbersInIndividualScratchcard(remainingScratchcards[0])
		totalScratchcardsWon += scratchcardsWon
		remainingScratchcards = remainingScratchcards[1:]
		if !hasSeen[scratchcardIndex] && scratchcardIndex < len(allScratchcards) {
			// add original
			remainingScratchcards = append(remainingScratchcards, allScratchcards[scratchcardIndex])
		}
		hasSeen[scratchcardIndex] = true
		if scratchcardsWon > 0 {
			remainingScratchcards = append(remainingScratchcards, allScratchcards[scratchcardIndex:scratchcardIndex+scratchcardsWon]...)
		}
		index++
	}
	return index
}

func findWinningNumbersInIndividualScratchcard(scratchcard string) (int, int) {
	scratchcardIndex, _ := strconv.Atoi(strings.TrimSpace(strings.Split(scratchcard, ":")[0]))
	winningNumbers := strings.Fields(strings.Split(strings.Split(scratchcard, ": ")[1], " | ")[0])
	candidateNumbers := make(map[string]int)
	numbers := strings.Fields(strings.Split(strings.Split(scratchcard, ": ")[1], " | ")[1])

	// add winningNumbers to winningNumbersList
	for _, number := range numbers {
		if number != "" {
			candidateNumbers[number]++
		}
	}

	scratchcardsWon := 0
	for _, winningNumber := range winningNumbers {
		if candidateNumbers[winningNumber] != 0 {
			scratchcardsWon++
		}
	}

	return scratchcardIndex, scratchcardsWon
}

func calculatePart1Prize(allNumbers []map[string]int, allWinningNumbers [][]string) int {
	totalPoints := 0
	for i, winningNumbers := range allWinningNumbers {
		scratchcardPoints := 0
		for _, winningNumber := range winningNumbers {
			if allNumbers[i][winningNumber] != 0 {
				if scratchcardPoints == 0 {
					scratchcardPoints = 1
				} else {
					scratchcardPoints *= 2
				}
			}
		}
		totalPoints += scratchcardPoints
	}
	return totalPoints
}

func buildWinningsNumbersAndNumbers(lines []string) ([]map[string]int, [][]string, []string) {
	allNumbersList := make([]map[string]int, 0)
	allWinningNumbersList := make([][]string, 0)
	allScratchcards := make([]string, 0)

	for _, line := range lines {
		numbersList := make(map[string]int)
		winningNumbersList := make([]string, 0)
		allScratchcards = append(allScratchcards, strings.Split(line, "Card ")[1])
		numbers := strings.Split(strings.Split(line, "| ")[1], " ")
		// add winningNumbers to winningNumbersList
		for _, number := range numbers {
			if number != "" {
				numbersList[number]++
			}
		}
		winningNumbers := strings.Split(strings.Split(strings.Split(line, " | ")[0], ": ")[1], " ")
		for _, winningNumber := range winningNumbers {
			if winningNumber != "" {
				winningNumbersList = append(winningNumbersList, winningNumber)
			}
		}
		allNumbersList = append(allNumbersList, numbersList)
		allWinningNumbersList = append(allWinningNumbersList, winningNumbersList)
	}
	return allNumbersList, allWinningNumbersList, allScratchcards
}
