package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"strings"
)

func main() {
	p := getPuzzle()
	p.trace(p.start)
	fmt.Println(len(p.hit))
}

type puzzle struct {
	start common.Coords
	input map[common.Coords]rune

	hit map[common.Coords]struct{}
}

func (p *puzzle) trace(c common.Coords) {
	for {
		fmt.Println("tracing", c)
		c = c.Add(common.S)
		char, exists := p.input[c]
		if !exists {
			return
		}

		if char == '.' {
			continue
		}

		if char != '^' {
			panic("we messed up")
		}

		p.hit[c] = struct{}{}
		p.trace(c.Add(common.W))
		p.trace(c.Add(common.E))
		return
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
		hit:   map[common.Coords]struct{}{},
	}
}
