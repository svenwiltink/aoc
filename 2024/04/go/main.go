package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}
func part1() {
	lines := common.GetLines()

	word := `XMAS`
	var part1 int
	for x := range len(lines[0]) {
		for y := range len(lines) {
			if lines[y][x] != word[0] {
				continue
			}
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue
					}
					remainder := []byte(word[1:])
					valid := true
					xLetter := x
					yLetter := y
					for _, letter := range remainder {
						xLetter += dx
						yLetter += dy

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
					}

					if valid {
						part1++
					}
				}
			}
		}
	}

	fmt.Println(part1)
}
func part2() {
	lines := common.GetLines()
	var part2 int
	for x := 1; x < len(lines[0])-1; x++ {
		for y := 1; y < len(lines)-1; y++ {
			if lines[y][x] != 'A' {
				continue
			}

			tl := lines[y-1][x-1]
			tr := lines[y-1][x+1]
			bl := lines[y+1][x-1]
			br := lines[y+1][x+1]

			word := string([]byte{tl, 'A', br})
			word2 := string([]byte{bl, 'A', tr})

			if word != `MAS` && word != `SAM` {
				continue
			}

			if word2 != `MAS` && word2 != `SAM` {
				continue
			}

			part2++
		}
	}

	fmt.Println(part2)
}
