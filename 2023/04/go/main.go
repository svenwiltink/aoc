package main

import (
	"fmt"
	"math"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	games := getGames()
	fmt.Println(common.Fold(func(value int, current Game) int {
		return value + current.Score()
	}, games))
}

func part2() {
	games := getGames()
	fmt.Println(common.Fold(func(value int, current Game) int {
		return value + playGame(games, current) + 1
	}, games))
}

func playGame(games []Game, current Game) int {
	winnerCount := current.WinnerCount()

	return common.Fold(func(result int, currentGame Game) int {
		return result + playGame(games, currentGame)
	}, games[current.ID:current.ID+winnerCount]) + winnerCount
}

func getGames() []Game {
	var games []Game
	lines := common.GetLines()
	for _, line := range lines {
		games = append(games, parseGame(line))
	}

	return games
}

func parseGame(line string) Game {
	var game Game
	card, scores := common.Split(line, ":")

	if _, err := fmt.Sscanf(card, "Card %d", &game.ID); err != nil {
		panic(err)
	}

	winning, ours := common.Split(scores, "|")
	game.Winners = common.ExtractNumbers(winning)
	game.Ours = common.ExtractNumbers(ours)

	return game
}

type Game struct {
	ID      int
	Winners []int
	Ours    []int
}

func (g Game) WinnerCount() int {
	return len(common.Intersect(g.Winners, g.Ours))
}

func (g Game) Score() int {
	num := g.WinnerCount()
	if num == 0 {
		return 0
	}

	return int(math.Pow(2, float64(num-1)))
}
