package main

import (
	"fmt"
	"strings"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	times, distances := parseInput1()
	simuluate(times, distances)
}

func part2() {
	times, distances := parseInput2()
	simuluate(times, distances)
}

func simuluate(times, distances []int) {
	winning := 1
	for i, time := range times {
		distance := distances[i]

		var wins int
		for t := 0; t < time; t++ {
			ourdistance := t * (time - t)
			if ourdistance > distance {
				wins++
			}
		}

		winning *= wins
	}

	fmt.Println(winning)
}

func parseInput1() ([]int, []int) {
	lines := common.GetLines()

	times, distances := lines[0], lines[1]
	_, timeNums := common.Split(times, ":")
	_, distancenums := common.Split(distances, ":")

	return common.ExtractNumbers(timeNums), common.ExtractNumbers(distancenums)
}

func parseInput2() ([]int, []int) {
	lines := common.GetLines()

	times, distances := lines[0], lines[1]
	_, timeNums := common.Split(times, ":")
	_, distancenums := common.Split(distances, ":")

	return common.ExtractNumbers(strings.ReplaceAll(timeNums, " ", "")), common.ExtractNumbers(strings.ReplaceAll(distancenums, " ", ""))
}
