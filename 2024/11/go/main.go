package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"math"
)

func main() {
	numbers := common.ExtractNumbers(common.GetLines()[0])
	var total int
	for _, num := range numbers {
		total += common.Memoize2(banana, num, 75)
	}

	fmt.Println(total)
}

func banana(callback func(int, int) int, stone int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	if stone == 0 {
		return callback(1, blinks-1)
	}

	digits := int(math.Floor(math.Log10(float64(stone)))) + 1
	isEven := digits%2 == 0

	if isEven {
		left := stone / int(math.Pow(10, float64(digits/2)))
		right := stone % int(math.Pow(10, float64(digits/2)))
		return callback(left, blinks-1) + callback(right, blinks-1)
	}

	return callback(stone*2024, blinks-1)
}
