package main

import (
	"fmt"
	"math"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	universe := parseUniverse()
	universe.Multiplier = 1

	pairs := universe.GalaxyPairs()
	distances := common.Map(universe.Distance, pairs)

	fmt.Println(common.Sum(distances))
}

func part2() {
	universe := parseUniverse()
	universe.Multiplier = 1_000_000 - 1

	pairs := universe.GalaxyPairs()
	distances := common.Map(universe.Distance, pairs)

	fmt.Println(common.Sum(distances))
}

func parseUniverse() Universe {
	var u Universe

	lines := common.GetLines()
	for y, line := range lines {
		var containsGalaxy bool
		for x, char := range line {
			if char == '#' {
				u.Galaxies = append(u.Galaxies, Galaxy{x, y})
				containsGalaxy = true
			}
		}

		if !containsGalaxy {
			u.EmptyRows = append(u.EmptyRows, y)
		}
	}

	for x := 0; x < len(lines); x++ {
		var containsGalaxy bool
		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				containsGalaxy = true
				break
			}
		}

		if !containsGalaxy {
			u.EmptyColumns = append(u.EmptyColumns, x)
		}
	}

	return u
}

type Universe struct {
	Galaxies     []Galaxy
	EmptyRows    []int
	EmptyColumns []int
	Multiplier   int
}

func (u Universe) GalaxyPairs() (pairs []common.Pair[Galaxy]) {
	for i, a := range u.Galaxies {
		for _, b := range u.Galaxies[i:] {
			if a == b {
				continue
			}

			pairs = append(pairs, common.Pair[Galaxy]{a, b})
		}
	}

	return pairs
}

func (u Universe) Distance(p common.Pair[Galaxy]) int {
	a, b := p[0], p[1]
	minX, minY := int(math.Min(float64(a.X), float64(b.X))), int(math.Min(float64(a.Y), float64(b.Y)))
	maxX, maxY := int(math.Max(float64(a.X), float64(b.X))), int(math.Max(float64(a.Y), float64(b.Y)))

	xDist := maxX - minX + checkBounds(u.EmptyColumns, minX, maxX)*u.Multiplier
	yDist := maxY - minY + checkBounds(u.EmptyRows, minY, maxY)*u.Multiplier

	return xDist + yDist
}

func checkBounds(bounds []int, min, max int) int {
	var sum int
	for _, num := range bounds {
		if num > min && num < max {
			sum++
		}
	}

	return sum
}

type Galaxy struct {
	X int
	Y int
}
