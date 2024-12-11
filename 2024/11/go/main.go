package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"math"
)

func main() {
	numbers := common.ExtractNumbers(common.GetLines()[0])
	var total int
	var cache = make(map[cachekey]int)
	for _, num := range numbers {
		total += splitStone(num, 75, cache)
	}

	fmt.Println(total)
}

type cachekey [2]int

func splitStone(stone int, blinks int, cache map[cachekey]int) int {
	if item, exists := cache[cachekey{stone, blinks}]; exists {
		return item
	}

	if blinks == 0 {
		return 1
	}

	if stone == 0 {
		result := splitStone(1, blinks-1, cache)
		cache[cachekey{stone, blinks}] = result
		return result
	}

	digits := int(math.Floor(math.Log10(float64(stone)))) + 1
	isEven := digits%2 == 0

	if isEven {
		left := stone / int(math.Pow(10, float64(digits/2)))
		right := stone % int(math.Pow(10, float64(digits/2)))
		result := splitStone(left, blinks-1, cache) + splitStone(right, blinks-1, cache)
		cache[cachekey{stone, blinks}] = result
		return result
	}

	result := splitStone(stone*2024, blinks-1, cache)
	cache[cachekey{stone, blinks}] = result
	return result
}
