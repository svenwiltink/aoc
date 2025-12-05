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
	for {
		var changed bool
		for i, iItem := range fresh {
			for j, jItem := range fresh {
				if i == j {
					continue
				}

				if iItem.count() == 0 && jItem.count() == 0 {
					continue
				}

				merged := iItem.merge(jItem)
				// o my god, we merged!
				if len(merged) == 1 {
					fresh[i] = merged[0]
					fresh[j] = FreshRange{} // this sucks
					changed = true
					break
				}
			}
		}

		if !changed {
			break
		}
	}

	fmt.Println(common.Sum(common.Map(FreshRange.count, fresh)))
}

type FreshRange [2]int

func (f FreshRange) isFresh(i int) bool {
	return i >= f[0] && i <= f[1]
}

func (f FreshRange) count() int {
	if f[0] == 0 && f[1] == 0 {
		return 0
	}
	return f[1] - f[0] + 1
}

func (f FreshRange) merge(o FreshRange) []FreshRange {
	// o end earlier than f, no overlap
	if o[1] < f[0] {
		return []FreshRange{f, o}
	}

	// o starts after f, no overlap
	if o[0] > f[1] {
		return []FreshRange{f, o}
	}

	// o is entirely inside f, fully contained
	if o[0] >= f[0] && o[1] <= f[1] {
		return []FreshRange{f}
	}

	// f is entirely inside o, fully contained but reversed
	if f[0] >= o[0] && f[1] <= o[1] {
		return []FreshRange{o}
	}

	// we are (partially) overlapping, check lower and higher end
	// lower side, cut to start of f
	if o[0] < f[0] {
		return []FreshRange{{o[0], f[1]}}
	}

	// high side, take lower f and higher o
	if o[1] > f[1] {
		return []FreshRange{{f[0], o[1]}}
	}

	panic("did I mess up?")
}
