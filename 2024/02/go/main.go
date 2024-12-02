package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"slices"
)

func main() {
	part1()
	part2()
}

func part1() {
	input := getInput()
	safe := common.Filter(isSafe, input)
	fmt.Println(len(safe))
}

func part2() {
	input := getInput()
	safe := common.Filter(dampened, input)
	fmt.Println(len(safe))
}

func differences(level []int) []int {
	var diffs []int

	current, remainder := level[0], level[1:]

	for len(remainder) > 0 {
		next := remainder[0]
		diffs = append(diffs, next-current)

		current, remainder = remainder[0], remainder[1:]
	}

	return diffs
}

func dampened(level []int) bool {
	if isSafe(level) {
		return true
	}

	for i := range len(level) {
		mutation := slices.Clone(level)
		mutation = slices.Delete(mutation, i, i+1)
		if isSafe(mutation) {
			return true
		}
	}

	return false
}

func isSafe(level []int) bool {
	diffs := differences(level)
	positives := common.Filter(common.IsPositive[int], diffs)
	negatives := common.Filter(common.IsNegative[int], diffs)

	if len(positives) > 0 && len(negatives) > 0 {
		return false
	}

	for _, d := range diffs {
		abs := common.Abs(d)
		if abs < 1 || abs > 3 {
			return false
		}
	}

	return true
}

func getInput() [][]int {
	lines := common.GetLines()
	input := common.Map(common.ExtractNumbers, lines)
	return input
}
