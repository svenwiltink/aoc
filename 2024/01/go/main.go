package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"slices"
)

func main() {
	//part1()
	part2()
}

func part1() {
	left, right := getInput()
	slices.Sort(left)
	slices.Sort(right)

	var distances []int

	for i, left := range left {
		distance := common.Abs(right[i] - left)
		distances = append(distances, distance)
	}

	fmt.Println(distances)
	fmt.Println(common.Sum(distances))
}

func part2() {
	left, right := getInput()

	lookup := make(map[int]int)
	for _, num := range right {
		lookup[num]++
	}

	var total int
	for _, num := range left {
		total += num * lookup[num]
	}

	fmt.Println(total)
}

func getInput() ([]int, []int) {
	lines := common.GetLines()

	var left, right []int
	for _, line := range lines {
		var l, r int
		_, err := fmt.Sscanf(line, "%d %d", &l, &r)
		if err != nil {
			panic(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}
