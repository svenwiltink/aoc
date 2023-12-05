package main

import (
	"fmt"
	"slices"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}

func part2() {
	lines := common.GetLines()
	seeds := parseSeeds(lines)
	maps := parseMaps(lines[2:])

	var results []int
	for i := seeds[0]; i < seeds[0]+seeds[1]; i++ {
		results = append(results, placeSeed(maps, i))
	}

	for i := seeds[2]; i < seeds[2]+seeds[3]; i++ {
		results = append(results, placeSeed(maps, i))
	}

	fmt.Println(slices.Min(results))
}

func part1() {
	lines := common.GetLines()
	seeds := parseSeeds(lines)
	maps := parseMaps(lines[2:])

	seeds = common.Map(func(seed int) int {
		return placeSeed(maps, seed)
	}, seeds)

	fmt.Println(slices.Min(seeds))
}

func placeSeed(maps []Map, seed int) int {
	for _, entry := range maps {
		seed = entry.Map(seed)
	}

	return seed
}

func parseSeeds(lines []string) []int {
	line := lines[0]
	_, nums := common.Split(line, ":")
	return common.ExtractNumbers(nums)
}

func parseMaps(lines []string) []Map {
	var result []Map

	if len(lines) == 0 {
		return result
	}

	var entry Map
	lines = lines[1:]
	for index, line := range lines {
		if line == "" {
			result = append([]Map{entry}, parseMaps(lines[index+1:])...)
			return result
		}

		nums := common.ExtractNumbers(line)
		entry.Recipes = append(entry.Recipes, Recipe{nums[0], nums[1], nums[2]})

	}

	return []Map{entry}
}

type Map struct {
	From    string
	To      string
	Recipes []Recipe
}

func (m Map) Map(seed int) int {
	for _, recipe := range m.Recipes {
		if seed < recipe.Start || seed > recipe.Start+recipe.Range {
			continue
		}
		return seed + recipe.Destination - recipe.Start
	}

	return seed
}

type Recipe struct {
	Destination int
	Start       int
	Range       int
}
