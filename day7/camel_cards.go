package main

import (
	"advent-of-code-2023/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := common.ReadFileLines("day7/input/input.txt")
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
			fmt.Println(hand + " Five of a kind")
			rank = 0
		} else if isFourOfAKind(hand, isJokersWild) {
			fmt.Println(hand + " Four of a kind")
			rank = 1
		} else if isFullHouse(hand, isJokersWild) {
			fmt.Println(hand + " FullHouse")
			rank = 2
		} else if isThreeOfAKind(hand, isJokersWild) {
			fmt.Println(hand + " Three of a kind")
			rank = 3
		} else if isTwoPair(hand) {
			fmt.Println(hand + " Two Pair")
			rank = 4
		} else if isPair(hand, isJokersWild) {
			fmt.Println(hand + " Pair")
			rank = 5
		} else if isHighCard(hand) {
			fmt.Println(hand + " High Card")
			rank = 6
		}
		handRanksAndBids = append(handRanksAndBids, CardHandRankAndBid{Hand: hand, Rank: rank, Bid: bid})
	}
	fmt.Println(handRanksAndBids)

	cardValues := mapOfCardValues
	if isJokersWild {
		cardValues = mapOfCardValuesPart2
	}

	sort.Slice(handRanksAndBids, func(i, j int) bool {
		if handRanksAndBids[i].Rank == handRanksAndBids[j].Rank {
			hand1 := handRanksAndBids[i].Hand
			hand2 := handRanksAndBids[j].Hand
			for i := 0; i < len(hand1); i++ {
				if cardValues[string(hand1[i])] > cardValues[string(hand2[i])] {
					return true
				} else if cardValues[string(hand1[i])] < cardValues[string(hand2[i])] {
					return false
				}
			}
		}
		return handRanksAndBids[i].Rank < handRanksAndBids[j].Rank
	})

	fmt.Println(handRanksAndBids)
	for _, handRankAndBid := range handRanksAndBids {
		fmt.Println(handRankAndBid.Hand + " " + strconv.Itoa(handRankAndBid.Rank) + " " + strconv.Itoa(handRankAndBid.Bid))
	}
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

var mapOfCardValuesPart2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 11,
	"9": 10,
	"8": 9,
	"7": 8,
	"6": 7,
	"5": 6,
	"4": 5,
	"3": 4,
	"2": 3,
	"J": 2,
}

type CardHandRankAndBid struct {
	Hand string
	Rank int // 0 being the best/highest, 6 being the lowest
	Bid  int
}

func isFiveOfAKind(hand string, isJokersWild bool) bool {
	maxOccurrences := getMaxOccurrences(hand)
	numOfJokers := strings.Count(hand, "J")
	return maxOccurrences == 5 || (isJokersWild && maxOccurrences+numOfJokers == 5)
}

func isFourOfAKind(hand string, isJokersWild bool) bool {
	maxOccurrences := getMaxOccurrences(hand)
	numOfJokers := strings.Count(hand, "J")
	return maxOccurrences == 4 || (isJokersWild && maxOccurrences+numOfJokers == 4)
}

func isFullHouse(hand string, isJokersWild bool) bool {
	if isJokersWild {
		numOfJokers := strings.Count(hand, "J")
		if len(countOccurrencesOfEachCard(hand)) == 3 && numOfJokers == 1 {
			return true
		}
	}
	return len(countOccurrencesOfEachCard(hand)) == 2
}

func isThreeOfAKind(hand string, isJokersWild bool) bool {
	maxOccurrences := getMaxOccurrences(hand)
	numOfJokers := strings.Count(hand, "J")
	return maxOccurrences == 3 || (isJokersWild && maxOccurrences+numOfJokers == 3)
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
	maxOccurrences := getMaxOccurrences(hand)
	numOfJokers := strings.Count(hand, "J")
	return maxOccurrences == 2 || (isJokersWild && maxOccurrences+numOfJokers == 2)
}

func isHighCard(hand string) bool {
	return len(countOccurrencesOfEachCard(hand)) == 5
}

func getMaxOccurrences(hand string) int {
	max := 0
	for char, occurrences := range countOccurrencesOfEachCard(hand) {
		if occurrences > max && char != "J" {
			max = occurrences
		}
	}
	return max
}

func countOccurrencesOfEachCard(hand string) map[string]int {
	occurences := make(map[string]int)
	for _, card := range hand {
		occurences[string(card)]++
	}
	return occurences
}
