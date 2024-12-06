package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
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
	maze.simulate()
	fmt.Println(len(maze.Visited))

	var result int
	for _, v := range maze.obstructionLocations() {
		if maze.tryObstruction(v) {
			result++
		}
	}

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

func (m *Map) simulate() bool {
	for m.moveGuard() {
	}

	return m.loopDetected()
}

func (m *Map) tryObstruction(o coords) bool {
	m.reset()
	m.Obstructions[o] = true
	result := m.simulate()
	m.Obstructions[o] = false
	return result
}

func (m *Map) reset() {
	m.Visited = make(map[coords]bool)
	m.LoopDetection = make(map[Guard]bool)
	m.Guard = m.StartGuard

	m.Visited[coords{m.Guard.X, m.Guard.Y}] = true
	m.LoopDetection[m.Guard] = true
}

func (m *Map) loopDetected() bool {
	return m.inBounds(m.Guard.X, m.Guard.Y)
}

func (m *Map) moveGuard() bool {
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
		return false
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

		if m.LoopDetection[m.Guard] {
			return false
		}
		return true
	}

	m.Guard.X, m.Guard.Y = x, y
	m.Visited[coords{x, y}] = true

	if m.LoopDetection[m.Guard] {
		return false
	}

	m.LoopDetection[m.Guard] = true

	return true
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
