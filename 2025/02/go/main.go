package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"strconv"
	"strings"
)

type pair [2]int

func main() {
	ranges := getRanges()
	fmt.Println(common.Sum(common.Map(checkRange, ranges)))

}

func checkRange(p pair) int {
	var result int
	for i := p[0]; i <= p[1]; i++ {
		if checkNumber(i) {
			result += i
		}
	}

	return result
}

func checkNumber(n int) bool {
	num := strconv.Itoa(n)
	if len(num)%2 > 0 {
		return false
	}

	left, right := num[:len(num)/2], num[len(num)/2:]
	return left == right
}

func getRanges() []pair {
	input := common.GetInput()
	ranges := strings.Split(input, ",")
	return common.Map(func(k string) pair {
		var left, right int
		_, err := fmt.Sscanf(k, "%d-%d", &left, &right)
		if err != nil {
			panic(err)
		}
		return pair{left, right}
	}, ranges)
}
