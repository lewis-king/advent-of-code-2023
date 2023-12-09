package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day9/input/input.txt")
	total := 0
	part2 := true
	for _, line := range lines {
		words := strings.Fields(line)
		if part2 {
			for i := 0; i < len(words)/2; i++ {
				words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
			}
		}
		var numbers []int
		for _, word := range words {
			num, _ := strconv.Atoi(word)
			numbers = append(numbers, num)
		}
		nextValue := calculateNextValue(numbers, numbers[len(numbers)-1])
		fmt.Println(nextValue)
		total += nextValue
	}
	fmt.Println(total)
}

func calculateNextValue(numbers []int, lastNum int) int {
	var diffs []int
	for i := 0; i < len(numbers)-1; i++ {
		diffs = append(diffs, numbers[i+1]-numbers[i])
	}
	if checkAllSame(diffs) {
		return lastNum + diffs[len(diffs)-1]
	} else {
		return numbers[len(numbers)-1] + calculateNextValue(diffs, diffs[len(diffs)-1])
	}
}

func checkAllSame(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] != numbers[i+1] {
			return false
		}
	}
	return true
}
