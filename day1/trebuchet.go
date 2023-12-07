package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFileLines() []string {
	content, err := os.ReadFile("day1/input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// read lines from file
	return strings.Split(string(content), "\n")
}

var mapOfTextualRepToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	lines := readFileLines()
	sumPart1 := 0
	sumPart2 := 0
	for _, s := range lines {
		// Part 1
		var firstNum = -1
		var lastNum = -1
		firstNum, lastNum = readDigits(s, firstNum, lastNum, false)
		concat, _ := strconv.Atoi(strconv.Itoa(firstNum) + strconv.Itoa(lastNum))
		sumPart1 += concat

		// Part 2
		firstNum = -1
		lastNum = -1
		firstNum, lastNum = readDigits(s, firstNum, lastNum, true)
		concat, _ = strconv.Atoi(strconv.Itoa(firstNum) + strconv.Itoa(lastNum))
		sumPart2 += concat
	}
	println(sumPart1)
	println(sumPart2)
}

func readDigits(s string, firstNum int, lastNum int, textualDigits bool) (int, int) {
	for i := 0; i < len(s); i++ {
		// is it a numerical digit?
		if v, err := strconv.Atoi(string(s[i])); err == nil {
			if firstNum == -1 {
				firstNum = v
			}
			lastNum = v
		} else if textualDigits {
			for k := range mapOfTextualRepToDigit {
				if i+len(k)-1 < len(s) && s[i:i+len(k)] == k {
					lastNum = mapOfTextualRepToDigit[k]
					if firstNum == -1 {
						firstNum = mapOfTextualRepToDigit[k]
					}
					break
				}
			}
		}
	}
	return firstNum, lastNum
}
