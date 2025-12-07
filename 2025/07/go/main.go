package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"strings"
)

func main() {
	p := getPuzzle()
	p2 := p.trace(p.start)
	fmt.Println(len(p.hit))
	fmt.Println(p2)
}

type puzzle struct {
	start common.Coords
	input map[common.Coords]rune

	hit map[common.Coords]int
}

func (p *puzzle) trace(c common.Coords) (result int) {
	for {
		c = c.Add(common.S)
		char, exists := p.input[c]
		if !exists {
			return 1
		}

		if char == '.' {
			continue
		}

		if char != '^' {
			panic("we messed up")
		}

		if hits, exists := p.hit[c]; exists {
			return hits
		}

		left := p.trace(c.Add(common.W))
		right := p.trace(c.Add(common.E))
		p.hit[c] = left + right
		return left + right
	}
}

func getPuzzle() puzzle {
	lines := common.GetLines()
	start := strings.Index(lines[0], "S")

	input := make(map[common.Coords]rune)

	for y, line := range lines {
		for x, char := range line {
			input[common.Coords{x, y}] = char
		}
	}

	return puzzle{
		start: common.Coords{start, 0},
		input: input,
		hit:   map[common.Coords]int{},
	}
}
