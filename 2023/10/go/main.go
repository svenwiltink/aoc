package main

import (
	"fmt"
	"strings"

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

func isHairpin(a, b rune) bool {
	switch {
	case a == NE && b == NW:
		fallthrough
	case a == SE && b == SW:
		return true
	default:
		return false
	}
}

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
	// part1()
	part2()
}

func part1() {
	m := maze(common.GetLines())

	startX, startY := m.findStart()
	direction := m.startDirection()

	x, y, direction := startX, startY, direction
	steps := 0
	for {
		steps++
		x, y, direction = m.traverse(x, y, direction)
		if x == startX && y == startY {
			break
		}
	}

	fmt.Println(steps / 2)
}

func part2() {
	m := maze(common.GetLines())

	startX, startY := m.findStart()
	direction := m.startDirection()

	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true

	x, y, direction := startX, startY, direction
	for {
		x, y, direction = m.traverse(x, y, direction)
		visited[[2]int{x, y}] = true
		if x == startX && y == startY {
			break
		}
	}

	// ugly hack to repair the input with the correct pipe
	m[startY] = strings.ReplaceAll(m[startY], "S", "J")

	count := 0
	for y, line := range m {
		var inside bool
		var previousChar rune

		for x, c := range line {
			v := visited[[2]int{x, y}]

			if !v {
				previousChar = '.'
				if inside {
					count++
				}
				continue
			}

			switch c {
			case '|':
				inside = !inside
			case '-':
			default:
				if previousChar == '.' {
					previousChar = c
					inside = !inside
					continue
				}

				if isHairpin(previousChar, c) {
					inside = !inside
				}

				previousChar = '.' // pretend the corner doesn't exist

			}

		}
	}

	fmt.Println(count)

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
			fmt.Println("out of bounds", newX, newY, bounds)
			continue
		}

		pipe := m[newY][newX]

		if directionChange[direction][rune(pipe)] != 0 {
			return direction
		}
	}

	panic("no start direction")
}

func (m maze) traverse(startX, startY int, direction rune) (newX, newY int, newDirection rune) {
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
