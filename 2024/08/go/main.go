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
				g.data[[2]int{x, y}] = char
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

			anti := coordA.add(coordA.distance(coordB).mul(2))
			if g.inBounds(anti) {
				part1[anti] = true
			}

			vector := coordA.distance(coordB).normalise()
			anti = coordA
			for g.inBounds(anti) {
				part2[anti] = true
				anti = anti.add(vector)
			}
		}
	}

	fmt.Println(len(part1))
	fmt.Println(len(part2))
}

type coords [2]int

func (c coords) distance(o coords) coords {
	return coords{o[0] - c[0], o[1] - c[1]}
}

func (c coords) add(d coords) coords {
	return coords{c[0] + d[0], c[1] + d[1]}
}

func (c coords) mul(mul int) coords {
	return coords{c[0] * mul, c[1] * mul}
}

func (c coords) div(div int) coords {
	return coords{c[0] / div, c[1] / div}
}

func (c coords) normalise() coords {
	div := common.Abs(common.Gcd(c[0], c[1]))
	return c.div(div)
}

type grid struct {
	height, width int
	data          map[coords]rune
}

func (g grid) inBounds(c coords) bool {
	if c[0] < 0 || c[1] < 0 {
		return false
	}

	if c[0] >= g.width || c[1] >= g.height {
		return false
	}

	return true
}
