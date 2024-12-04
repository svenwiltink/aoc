package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
)

func main() {
	lines := common.GetLines()

	var part1, part2 int

	part2lookup := make(map[[2]int]bool)
	for x := range len(lines[0]) {
		for y := range len(lines) {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					valid := checkWord(lines, x, y, "XMAS", dx, dy)
					if valid {
						part1++
					}

					// both dy and dx must be <> 0 in order for it to be a diagonal
					if dy != 0 && dx != 0 {
						valid = checkWord(lines, x, y, "MAS", dx, dy)
						coord := [2]int{x + dx, y + dy}
						if valid {
							current := part2lookup[coord]
							if current {
								part2++
							}

							part2lookup[coord] = true
						}
					}
				}
			}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func checkWord(lines []string, x int, y int, word string, dx int, dy int) bool {
	valid := true
	xLetter := x
	yLetter := y
	for _, letter := range []byte(word) {

		//bound check
		if xLetter < 0 || xLetter >= len(lines[0]) {
			valid = false
			break
		}

		if yLetter < 0 || yLetter >= len(lines) {
			valid = false
			break
		}

		if letter != lines[yLetter][xLetter] {
			valid = false
			break
		}

		xLetter += dx
		yLetter += dy
	}
	return valid
}
