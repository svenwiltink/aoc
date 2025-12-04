package main

import (
	"fmt"
	. "github.com/svenwiltink/aoc/common"
)

func main() {
	p := getPuzzle()
	num := p.sweep()
	fmt.Println(num)

	for next := p.sweep(); next > 0; next = p.sweep() {
		num += next
	}

	fmt.Println(num)
}

func getPuzzle() puzzle {
	rolls := make(map[Coords]rune)
	for y, line := range GetLines() {
		for x, char := range line {
			rolls[Coords{x, y}] = char
		}
	}

	return puzzle{rolls: rolls}
}

type puzzle struct {
	rolls map[Coords]rune
}

func (p *puzzle) sweep() int {
	var cleanup []Coords
	for coords, char := range p.rolls {
		if char != '@' {
			continue
		}

		var count int
		for _, dir := range Directions8 {
			lookup := coords.Add(dir)
			value := p.rolls[lookup]
			if value == '@' {
				count++
			}
		}
		if count < 4 {
			cleanup = append(cleanup, coords)
		}
	}

	for _, c := range cleanup {
		delete(p.rolls, c)
	}

	return len(cleanup)
}
