package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
)

func main() {
	lines := common.GetLines()

	g := grid{
		height: len(lines),
		width:  len(lines[0]),
		data:   make(map[coords]rune),
	}

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				g.data[coords(complex(float64(x), float64(y)))] = char
			}
		}
	}

	part1 := make(map[coords]bool)
	part2 := make(map[coords]bool)

	for coordA, a := range g.data {
		for coordB, b := range g.data {
			if a != b {
				continue
			}

			if coordA == coordB {
				continue
			}

			anti := coordA + ((coordB - coordA) * 2)
			if g.inBounds(anti) {
				part1[anti] = true
			}

			vector := (coordB - coordA).normalise()
			anti = coordA
			for g.inBounds(anti) {
				part2[anti] = true
				anti += vector
			}
		}
	}

	fmt.Println(len(part1))
	fmt.Println(len(part2))
}

type coords complex128

func (c coords) normalise() coords {
	div := common.Abs(common.Gcd(int(real(c)), int(imag(c))))
	return c / (coords(complex(float64(div), 0)))
}

type grid struct {
	height, width int
	data          map[coords]rune
}

func (g grid) inBounds(c coords) bool {
	if real(c) < 0 || imag(c) < 0 {
		return false
	}

	if int(real(c)) >= g.width || int(imag(c)) >= g.height {
		return false
	}

	return true
}
