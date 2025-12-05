package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := common.GetInput()
	rangeLines, ingredientLines := common.Split(input, "\n\n")
	ingredients := common.Map(
		common.MustClosure(strconv.Atoi),
		common.Map(
			strings.TrimSpace,
			slices.Collect(strings.Lines(ingredientLines)),
		),
	)

	var fresh []FreshRange
	for line := range strings.Lines(rangeLines) {
		var low, high int
		_, err := fmt.Sscanf(line, "%d-%d", &low, &high)
		if err != nil {
			panic(err)
		}
		fresh = append(fresh, FreshRange{low, high})
	}

	var p1 int
	for _, ingredient := range ingredients {
		for _, freshRange := range fresh {
			if freshRange.isFresh(ingredient) {
				p1++
				break
			}
		}
	}

	fmt.Println(p1)

	slices.SortFunc(fresh, func(a, b FreshRange) int {
		return a[0] - b[0]
	})

	var result int
	var prev FreshRange
	for _, f := range fresh {
		if f[1] <= prev[1] {
			continue
		}
		if f[0] <= prev[1] {
			f[0] = prev[1] + 1
		}

		result += f.count()
		prev = f
	}

	fmt.Println(result)
}

type FreshRange [2]int

func (f FreshRange) isFresh(i int) bool {
	return i >= f[0] && i <= f[1]
}

func (f FreshRange) count() int {
	return f[1] - f[0] + 1
}
