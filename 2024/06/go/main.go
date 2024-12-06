package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"maps"
)

func main() {
	lines := common.GetLines()
	maze := Map{
		Height:        len(lines),
		Width:         len(lines[0]),
		Obstructions:  make(map[coords]bool),
		Visited:       make(map[coords]bool),
		LoopDetection: make(map[Guard]bool),
	}
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				maze.Obstructions[coords{x, y}] = true
			}

			if char == '^' {
				maze.StartGuard.X = x
				maze.StartGuard.Y = y
				maze.StartGuard.Direction = '^'
			}
		}
	}

	maze.reset()
	result := maze.simulate()
	fmt.Println(len(maze.Visited))
	fmt.Println(result)

}

type coords [2]int

type Map struct {
	Height int
	Width  int

	Obstructions map[coords]bool
	Guard        Guard
	StartGuard   Guard
	Visited      map[coords]bool

	LoopDetection map[Guard]bool

	PlacedBlock bool
}

func (m *Map) obstructionLocations() []coords {
	var visited []coords
	for v, _ := range m.Visited {
		if v[0] == m.StartGuard.X && v[1] == m.StartGuard.Y {
			continue
		}
		visited = append(visited, v)
	}

	return visited
}

// simulate the guard and return true if a loop was detected
func (m *Map) simulate() int {
	var total int
	for {
		canContinue, count := m.move()
		total += count
		if !canContinue {
			break
		}
	}

	if m.inBounds(m.Guard.X, m.Guard.Y) {
		total++
	}

	return total
}

func (m *Map) reset() {
	m.Visited = make(map[coords]bool)
	m.LoopDetection = make(map[Guard]bool)
	m.Guard = m.StartGuard

	m.Visited[coords{m.Guard.X, m.Guard.Y}] = true
	m.LoopDetection[m.Guard] = true
}

func (m *Map) clone() *Map {
	return &Map{
		Height:        m.Height,
		Width:         m.Width,
		Obstructions:  maps.Clone(m.Obstructions),
		Guard:         m.Guard,
		StartGuard:    m.StartGuard,
		Visited:       maps.Clone(m.Visited),
		LoopDetection: maps.Clone(m.LoopDetection),
	}
}

// move the guard. Return true if the can continue moving or false if the guard is looping or out of bounds
func (m *Map) move() (bool, int) {
	x, y := m.Guard.X, m.Guard.Y
	switch m.Guard.Direction {
	case '^':
		y--
	case '>':
		x++
	case 'v':
		y++
	case '<':
		x--
	}

	if !m.inBounds(x, y) {
		m.Guard.X, m.Guard.Y = x, y
		return false, 0
	}

	if m.Obstructions[coords{x, y}] {
		switch m.Guard.Direction {
		case '^':
			m.Guard.Direction = '>'
		case '>':
			m.Guard.Direction = 'v'
		case 'v':
			m.Guard.Direction = '<'
		case '<':
			m.Guard.Direction = '^'
		}

		m.LoopDetection[m.Guard] = true
		return true, 0
	}

	var subTotal int
	// instead of moving the guard, try adding an obstruction as well if we haven't placed a block yet
	if !m.PlacedBlock && !m.Visited[coords{x, y}] {
		newMap := m.clone()
		newMap.Obstructions[coords{x, y}] = true
		newMap.PlacedBlock = true
		subTotal = newMap.simulate()
	}

	m.Guard.X, m.Guard.Y = x, y
	m.Visited[coords{x, y}] = true

	if m.LoopDetection[m.Guard] {
		return false, subTotal
	}

	m.LoopDetection[m.Guard] = true

	return true, subTotal
}

func (m *Map) inBounds(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if x >= m.Width {
		return false
	}

	if y >= m.Height {
		return false
	}

	return true
}

type Guard struct {
	X, Y      int
	Direction rune
}
