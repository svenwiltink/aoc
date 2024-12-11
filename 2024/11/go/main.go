package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"math"
)

func main() {
	numbers := common.ExtractNumbers(common.GetLines()[0])
	var part1 int
	var part2 int
	memoized := common.Memoize2(splitRock)
	for _, num := range numbers {
		part1 += memoized(num, 25)
		part2 += memoized(num, 75)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func splitRock(self func(int, int) int, stone int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	if stone == 0 {
		return self(1, blinks-1)
	}

	digits := int(math.Floor(math.Log10(float64(stone)))) + 1
	isEven := digits%2 == 0

	if isEven {
		left := stone / int(math.Pow(10, float64(digits/2)))
		right := stone % int(math.Pow(10, float64(digits/2)))
		return self(left, blinks-1) + self(right, blinks-1)
	}

	return self(stone*2024, blinks-1)
}
