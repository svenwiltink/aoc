package main

import (
	"fmt"
	"slices"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	lines := common.GetLines()
	numbers := common.Map(common.ExtractNumbers, lines)
	next := common.Map(getNext, numbers)

	fmt.Println(common.Sum(next))

	for _, num := range numbers {
		slices.Reverse(num)
	}

	next = common.Map(getNext, numbers)
	fmt.Println(common.Sum(next))
}

func getNext(numbers []int) int {
	diffs := getDiffs(numbers)

	last := numbers[len(numbers)-1]
	// if all numbers are the same the next number is the same as the current ones
	// [3,3,3] => [3,3,3,3]
	if common.Equal(diffs) {
		return diffs[0] + last
	}

	return getNext(diffs) + last
}

func getDiffs(numbers []int) []int {
	var diffs []int
	for i, num := range numbers[1:] {
		diffs = append(diffs, num-numbers[i])
	}

	return diffs
}
