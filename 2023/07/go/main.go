package main

import (
	"fmt"
	"slices"
	"unicode"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	hands := getHands()
	part1(hands)
	part2(hands)
}

func part1(hands []Hand) {
	slices.SortFunc(hands, Hand.CompareP1)

	var total int
	for i, hand := range hands {
		total += hand.Bid * (i + 1)
	}

	fmt.Println(total)
}

func part2(hands []Hand) {
	slices.SortFunc(hands, Hand.CompareP2)

	var total int
	for i, hand := range hands {
		total += hand.Bid * (i + 1)
	}

	fmt.Println(total)
}

func getHands() []Hand {
	return common.Map(getHand, common.GetLines())
}

func getHand(line string) Hand {
	var hand Hand
	if _, err := fmt.Sscanf(line, "%s %d", &hand.Cards, &hand.Bid); err != nil {
		panic(err)
	}

	return hand
}

type Hand struct {
	Cards string
	Bid   int
}

func (h Hand) Score() int {
	groups := h.GroupCards()
	return h.calculateScores(groups)
}

func (h Hand) calculateScores(groups map[string]int) int {
	var max int

	for _, value := range groups {
		if value > max {
			max = value
		}
	}

	switch len(groups) {
	case 1: // 5 of a kind
		return 7
	case 2: // 4 of a kind or full house
		// AAAAB
		// AAABB
		// AABBB
		// ABBBB

		// 4 of a kind
		if max == 4 {
			return 6
		}

		// full house
		return 5
	case 3: // 3 of a kind or two pair
		// AABBC
		// AAABC
		if max == 3 {
			return 4
		}

		return 3
	case 4: // one pair
		// AABCD
		return 2
	case 5:
		return 1
	default:
		panic("never happens")
	}
}

func (h Hand) InflatedScore() int {
	groups := h.GroupCards()

	var maxString string
	var max int

	for key, value := range groups {
		if value > max && key != "J" {
			max = value
			maxString = key
		}
	}

	if maxString != "J" {
		groups[maxString] += groups["J"]
		delete(groups, "J")
	}

	return h.calculateScores(groups)

}

func (h Hand) GroupCards() map[string]int {
	result := make(map[string]int)
	for _, r := range h.Cards {
		result[string(r)]++
	}

	return result
}

func (h Hand) CompareP2(o Hand) int {
	hScore := h.InflatedScore()
	oScore := o.InflatedScore()

	if hScore < oScore {
		return -1
	}

	if hScore > oScore {
		return 1
	}

	for i, r := range h.Cards {
		if runeToPoint(r, true) < runeToPoint([]rune(o.Cards)[i], true) {
			return -1
		}

		if runeToPoint(r, true) > runeToPoint([]rune(o.Cards)[i], true) {
			return 1
		}
	}

	panic(fmt.Sprintln("edge case", h.Cards, o.Cards))
}

func (h Hand) CompareP1(o Hand) int {
	hScore := h.Score()
	oScore := o.Score()

	if hScore < oScore {
		return -1
	}

	if hScore > oScore {
		return 1
	}

	for i, r := range h.Cards {
		if runeToPoint(r, false) < runeToPoint([]rune(o.Cards)[i], false) {
			return -1
		}

		if runeToPoint(r, false) > runeToPoint([]rune(o.Cards)[i], false) {
			return 1
		}
	}

	panic(fmt.Sprintln("edge case", h.Cards, o.Cards))
}

func runeToPoint(r rune, p2 bool) int {
	if unicode.IsDigit(r) {
		return int(r - '0')
	}

	lookup := map[rune]int{
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	if p2 {
		lookup['J'] = 1
	}

	return lookup[r]
}
