package main

import (
	"fmt"

	"github.com/svenwiltink/aoc/common"
)

const (
	vertical   = '|'
	horizontal = '-'
	NE         = 'L'
	NW         = 'J'
	SW         = '7'
	SE         = 'F'
	ground     = '.'
	start      = 'S'
)

var directionChange = map[rune]map[rune]rune{
	'E': {
		horizontal: 'E',
		SW:         'S',
		NW:         'N',
		'S':        'D', // end marker
	},
	'N': {
		vertical: 'N',
		SW:       'W',
		SE:       'E',
		'S':      'D', //end marker
	},
	'S': {
		vertical: 'S',
		NW:       'W',
		NE:       'E',
		'S':      'D', // end marker
	},
	'W': {
		horizontal: 'W',
		SE:         'S',
		NE:         'N',
		'S':        'D', // end marker
	},
}

func main() {
	part1()
}

func part1() {
	m := maze(common.GetLines())

	startX, startY := m.findStart()
	direction := m.startDirection()

	x, y, direction := startX, startY, direction
	steps := 0
	for {
		steps++
		x, y, direction = m.getNextDirection(x, y, direction)
		if x == startX && y == startY {
			break
		}
	}

	fmt.Println(steps / 2)
}

type maze []string

func (m maze) Print() {
	for _, line := range m {
		fmt.Println(line)
	}
}

func (m maze) findStart() (x, y int) {
	for y, line := range m {
		for x, char := range line {
			if char == start {
				return x, y
			}
		}
	}

	panic("no start found")
}

func (m maze) startDirection() rune {
	startX, startY := m.findStart()
	bounds := len(m) - 1

	for _, direction := range []rune{'N', 'E', 'S', 'W'} {
		newX, newY := startX, startY
		switch direction {
		case 'N':
			newY--
		case 'E':
			newX++
		case 'S':
			newY++
		case 'W':
			newX--
		}

		if newX < 0 || newY < 0 || newX > bounds || newY > bounds {
			continue
		}

		pipe := m[newY][newX]

		if directionChange[direction][rune(pipe)] != 0 {
			return direction
		}
	}

	panic("no start direction")
}

func (m maze) getNextDirection(startX, startY int, direction rune) (newX, newY int, newDirection rune) {
	newX, newY = startX, startY
	switch direction {
	case 'N':
		newY--
	case 'E':
		newX++
	case 'S':
		newY++
	case 'W':
		newX--
	}

	bounds := len(m) - 1
	if newX < 0 || newY < 0 || newX > bounds || newY > bounds {
		panic("OUT OF BOUNDS")
	}

	pipe := m[newY][newX]
	if directionChange[direction][rune(pipe)] == 0 {
		panic("invalid puzzle!!!!")
	}
	return newX, newY, directionChange[direction][rune(pipe)]
}
