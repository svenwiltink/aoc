package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"slices"
	"strconv"
)

func main() {
	data := getMap()
	fmt.Println(data.calculateScores())
}

func getMap() Map {
	lines := common.GetLines()
	result := make(Map)

	for y, line := range lines {
		for x, height := range line {
			num := common.Must(strconv.Atoi, string(height))
			result[coords{x, y}] = num
		}
	}

	return result
}

var directions = [...]coords{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

type Map map[coords]int

func (m Map) calculateScores() (int, int) {
	var part1, part2 int
	for coord, value := range m {
		if value != 0 {
			continue
		}

		reached, count := m.traversePath(coord, nil)
		part1 += len(reached)
		part2 += count
	}

	return part1, part2
}

func (m Map) traversePath(current coords, reached []coords) ([]coords, int) {
	curHeight := m[current]
	if curHeight == 9 {
		if !slices.Contains(reached, current) {
			return append(reached, current), 1
		}

		return reached, 1
	}

	target := curHeight + 1

	var total int

	for _, direction := range directions {
		//fmt.Println("checking direction for target height", target, direction)
		targetCoords := current.Add(direction)
		//fmt.Println("target coords", targetCoords)
		if m[targetCoords] == target {
			var sub int
			reached, sub = m.traversePath(targetCoords, reached)
			total += sub
		}

	}

	return reached, total
}

type coords [2]int

func (c coords) Add(o coords) coords {
	return coords{c[0] + o[0], c[1] + o[1]}
}
