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
	panel := common.Map(func(a string) []rune { return []rune(a) }, lines)

	common.TransposeInPlace(panel)
	cycle(panel)
	common.TransposeInPlace(panel)

	score := 0
	max := len(panel)
	for index, row := range panel {
		score += strings.Count(string(row), "O") * (max - index)
	}

	fmt.Println(score)
}

func part2() {
	lines := common.GetLines()
	panel := common.Map(func(a string) []rune { return []rune(a) }, lines)

	// this must converge to some score. keep track of where each index was first found
	hashes := make(map[string]int)
	len := 1000000000
	for i := 0; i < len; i++ {
		rotate(panel)

		h := SprintPanel(panel)
		if _, exists := hashes[h]; exists {
			// hash hit, we now know how long each cycle is.
			// Use mod len to fast forward a lot of iterations
			i = len - (len-i)%(i-hashes[h])
		}
		hashes[h] = i
	}

	fmt.Println(score(panel))
}

func score(panel [][]rune) int {
	score := 0
	max := len(panel)
	for index, row := range panel {
		score += strings.Count(string(row), "O") * (max - index)
	}

	return score
}

func rotate(panel [][]rune) {

	common.TransposeInPlace(panel)

	cycle(panel)

	common.TransposeInPlace(panel)

	cycle(panel)

	common.TransposeInPlace(panel)
	common.FlipVert(panel)

	cycle(panel)

	common.FlipVert(panel)
	common.TransposeInPlace(panel)
	common.FlipVert(panel)

	cycle(panel)

	common.FlipVert(panel)
}

func cycle(panel [][]rune) {
	for _, row := range panel {
		var empties []int // index of empty spots
		for index, char := range row {
			switch char {
			case '#':
				empties = nil
			case 'O':
				if len(empties) == 0 {
					continue
				}

				var empty int
				empty, empties = empties[0], append(empties[1:], index)
				row[empty] = 'O'
				row[index] = '.'
			case '.':
				empties = append(empties, index)
			}
		}
	}
}

func printPanel(p [][]rune) {
	for _, line := range p {
		fmt.Println(string(line))
	}
}

func SprintPanel(p [][]rune) string {
	result := make([]rune, len(p)*len(p))
	for i, row := range p {
		copy(result[i*len(p):], row)
	}

	return string(result)
}
