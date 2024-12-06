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
	lines := common.GetLines()
	var total int
	for _, line := range lines {
		springs, parts := common.Split(line, " ")
		numbers := common.ExtractNumbersSep(",", parts)
		total += getPossibilities(springs, numbers, 0, "")
	}
	fmt.Println(total)
}

func part2() {
	lines := common.GetLines()
	var total int
	for _, line := range lines {
		springs, parts := common.Split(line, " ")

		springs = strings.Repeat(springs+"?", 5)
		springs = springs[:len(springs)-1]

		parts = strings.Repeat(parts+",", 5)
		parts = parts[:len(parts)-1]

		numbers := common.ExtractNumbersSep(",", parts)
		total += getPossibilities(springs, numbers, 0, "")
	}

	fmt.Println(total)
}

type cacheKey struct {
	spring  string
	numbers [30]int
	count   int
}

var cache map[cacheKey]int = make(map[cacheKey]int)

func getPossibilities(springs string, numbers []int, count int, total string) (solutions int) {
	var numcache [30]int
	if copy(numcache[:], numbers) < len(numbers) {
		panic("cache save error, key too short")
	}

	key := cacheKey{springs, numcache, count}
	if value, exists := cache[key]; exists {
		return value
	}

	defer func() {
		cache[key] = solutions
	}()

	if springs == "" {
		if len(numbers) == 0 && count == 0 {
			return 1
		}

		if len(numbers) == 1 && numbers[0] == count {
			return 1
		}

		return 0
	}

	next := springs[0]
	branches := []byte{next}
	if next == '?' {
		branches = []byte{'.', '#'}
	}

	for _, next := range branches {
		if next == '#' {
			count++

			// number too big for group already. Don't need this case but speeds it up significantly
			if len(numbers) == 0 || count > numbers[0] {
				continue
			}

			solutions += getPossibilities(springs[1:], numbers, count, total+string(next))
			continue
		}

		// not expanding group. We might be in a number group
		if count > 0 {
			if len(numbers) > 0 && count == numbers[0] {
				solutions += getPossibilities(springs[1:], numbers[1:], 0, total+string(next))
			}
			continue
		}

		// not in a group, carry on
		solutions += getPossibilities(springs[1:], numbers, 0, total+string(next))
	}

	return solutions
}
