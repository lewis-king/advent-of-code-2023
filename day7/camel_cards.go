package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day7/input/example2.txt")
	var handRanksAndBids []CardHandRankAndBid

	// Part 1 set to false, Part 2 set to true
	isJokersWild := true

	for _, line := range lines {
		handAndBid := strings.Split(line, " ")
		hand := handAndBid[0]
		bid, _ := strconv.Atoi(handAndBid[1])
		rank := 0
		// NB: Checking of these predicates given a string is not very optimal because each check iterates the same hand. // TODO optimize i.e get number of occurrences of each card once and pass it to each predicate
		if isFiveOfAKind(hand, isJokersWild) {
			rank = 0
		} else if isFourOfAKind(hand, isJokersWild) {
			rank = 1
		} else if isFullHouse(hand, isJokersWild) {
			rank = 2
		} else if isThreeOfAKind(hand, isJokersWild) {
			rank = 3
		} else if isTwoPair(hand) {
			rank = 4
		} else if isPair(hand, isJokersWild) {
			rank = 5
		} else if isHighCard(hand) {
			rank = 6
		}
		handRanksAndBids = append(handRanksAndBids, CardHandRankAndBid{Hand: hand, Rank: rank, Bid: bid})
	}
	fmt.Println(handRanksAndBids)

	sort.Slice(handRanksAndBids, func(i, j int) bool {
		if handRanksAndBids[i].Rank == handRanksAndBids[j].Rank {
			hand1 := handRanksAndBids[i].Hand
			hand2 := handRanksAndBids[j].Hand
			for i := 0; i < len(hand1); i++ {
				if mapOfCardValues[string(hand1[i])] > mapOfCardValues[string(hand2[i])] {
					return true
				} else if mapOfCardValues[string(hand1[i])] < mapOfCardValues[string(hand2[i])] {
					return false
				}
			}
		}
		return handRanksAndBids[i].Rank < handRanksAndBids[j].Rank
	})

	fmt.Println(handRanksAndBids)
	totalWinnings := 0
	for i := 0; i < len(handRanksAndBids); i++ {
		globalRank := i + 1
		totalWinnings += handRanksAndBids[len(handRanksAndBids)-globalRank].Bid * globalRank
	}
	fmt.Println(totalWinnings)

}

var mapOfCardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type CardHandRankAndBid struct {
	Hand string
	Rank int // 0 being the best/highest, 6 being the lowest
	Bid  int
}

func isFiveOfAKind(hand string, isJokersWild bool) bool {
	if isJokersWild {
		numOfJokers := strings.Count(hand, "J")
		return len(countOccurrencesOfEachCard(hand)) == 2 && numOfJokers != 0
	} else {
		return len(countOccurrencesOfEachCard(hand)) == 1
	}
}

func isFourOfAKind(hand string, isJokersWild bool) bool {
	for _, occurrences := range countOccurrencesOfEachCard(hand) {
		if occurrences == 4 {
			return true
		} else if isJokersWild {
			numOfJokers := strings.Count(hand, "J")
			if (len(countOccurrencesOfEachCard(hand)) == 3 && numOfJokers == 2 && occurrences == 2) || occurrences == 3 && numOfJokers == 1 {
				return true
			}
		}
	}
	return false
}

func isFullHouse(hand string, isJokersWild bool) bool {
	if isJokersWild {
		numOfJokers := strings.Count(hand, "J")
		return len(countOccurrencesOfEachCard(hand)) == 3 && numOfJokers == 1
	}
	return len(countOccurrencesOfEachCard(hand)) == 2
}

func isThreeOfAKind(hand string, isJokersWild bool) bool {
	for _, occurrences := range countOccurrencesOfEachCard(hand) {
		if occurrences == 3 {
			return true
		} else if isJokersWild {
			numOfJokers := strings.Count(hand, "J")
			if occurrences+numOfJokers == 3 {
				return true
			}
		}
	}
	return false
}

func isTwoPair(hand string) bool {
	for _, occurrences := range countOccurrencesOfEachCard(hand) {
		if occurrences == 2 && len(countOccurrencesOfEachCard(hand)) == 3 {
			return true
		}
	}
	return false
}

func isPair(hand string, isJokersWild bool) bool {
	for _, occurrences := range countOccurrencesOfEachCard(hand) {
		if occurrences == 2 {
			return true
		} else if isJokersWild {
			numOfJokers := strings.Count(hand, "J")
			if occurrences+numOfJokers == 2 {
				return true
			}
		}
	}
	return false
}

func isHighCard(hand string) bool {
	return len(countOccurrencesOfEachCard(hand)) == 5
}

func countOccurrencesOfEachCard(hand string) map[string]int {
	occurences := make(map[string]int)
	for _, card := range hand {
		occurences[string(card)]++
	}
	return occurences
}
