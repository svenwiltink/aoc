package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
)

func main() {
	lines := common.GetLines()
	points := common.Map(func(s string) common.Coords {
		var x, y int
		_, err := fmt.Sscanf(s, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		return common.Coords{x, y}
	}, lines)

	pairs := common.Pairs(points)
	sizes := common.Map(size, pairs)
	_, value := common.Max(sizes)
	fmt.Println(value)
}

func size(pair common.Pair[common.Coords]) int {
	xd := pair[0].Sub(pair[1])

	return (common.Abs(xd[0]) + 1) * (common.Abs(xd[1]) + 1)
}
