package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"regexp"
)

var mul = regexp.MustCompile(`mul\(\d+,\d+\)`)
var onoff = regexp.MustCompile(`(don't\(\)|do\(\))`)

func main() {
	var sum1, sum2 int
	lines := common.GetLines()
	enabled := true
	for _, line := range lines {
		index := 0

		lookup := make([]bool, len(line))
		locs := onoff.FindAllStringIndex(line, -1)
		for _, l := range locs {
			value := line[l[0]:l[1]]

			for ; index < l[1]; index++ {
				lookup[index] = enabled
			}

			switch value {
			case `don't()`:
				enabled = false
			case `do()`:
				enabled = true
			default:
				panic(value)
			}
		}
		for ; index < len(line); index++ {
			lookup[index] = enabled
		}

		instances := mul.FindAllStringIndex(line, -1)
		for _, l := range instances {
			item := line[l[0]:l[1]]
			var a, b int
			_, err := fmt.Sscanf(item, "mul(%d,%d)", &a, &b)
			if err != nil {
				panic(err)
			}
			sum1 += a * b
			if !lookup[l[0]] {
				continue
			}
			sum2 += a * b
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
