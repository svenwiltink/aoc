package main

import (
	"fmt"
	"strings"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	lines := common.GetLines()
	games := parseGames(lines)
	part1(games)
	part2(games)
}

func part1(games []Game) {
	var total int

	for _, game := range games {
		if game.IsPossible(12, 13, 14) {
			total += game.ID
		}
	}

	fmt.Println(total)
}

func part2(games []Game) {
	var total int

	for _, game := range games {
		total += game.Power()
	}

	fmt.Println(total)
}

type Set map[string]int

type Game struct {
	ID   int
	Sets []Set
}

func (g Game) CalculateMax() Set {
	max := make(Set)

	for _, set := range g.Sets {
		for colour, num := range set {
			if num > max[colour] {
				max[colour] = num
			}
		}
	}

	return max
}

func (g Game) Power() int {
	max := g.CalculateMax()
	return max["red"] * max["blue"] * max["green"]
}

func (g Game) IsPossible(red, green, blue int) bool {
	max := g.CalculateMax()

	if max["red"] > red {
		return false
	}

	if max["green"] > green {
		return false
	}

	if max["blue"] > blue {
		return false
	}

	return true
}

func parseGames(lines []string) []Game {
	var games []Game
	for _, line := range lines {
		games = append(games, parseGame(line))
	}

	return games
}

func parseGame(line string) Game {
	game, sets := common.Split(line, ":")
	var id int
	if _, err := fmt.Sscanf(game, "Game %d", &id); err != nil {
		panic(err)
	}

	return Game{
		ID:   id,
		Sets: parseSets(sets),
	}
}

func parseSets(line string) []Set {
	var sets []Set
	for _, set := range strings.Split(line, ";") {
		sets = append(sets, parseSet(set))
	}

	return sets
}

func parseSet(input string) Set {
	set := make(Set)

	for _, part := range strings.Split(input, ",") {
		var num int
		var colour string

		if _, err := fmt.Sscanf(part, "%d %s", &num, &colour); err != nil {
			panic(err)
		}

		set[colour] = num
	}

	return set
}
