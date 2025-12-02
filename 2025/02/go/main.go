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
	fmt.Println(common.Sum(common.Map(common.Curry(checkRange, checkNumber1), ranges)))
	fmt.Println(common.Sum(common.Map(common.Curry(checkRange, checkNumber2), ranges)))
}

func checkRange1(p pair) int {
	var result int
	for i := p[0]; i <= p[1]; i++ {
		if checkNumber1(i) {
			result += i
		}
	}

	return result
}

func checkRange(checkFunc func(int) bool, p pair) int {
	var result int
	for i := p[0]; i <= p[1]; i++ {
		if checkFunc(i) {
			result += i
		}
	}

	return result
}

func checkNumber1(n int) bool {
	num := strconv.Itoa(n)
	if len(num)%2 > 0 {
		return false
	}

	left, right := num[:len(num)/2], num[len(num)/2:]
	return left == right
}

func checkNumber2(n int) bool {
	num := strconv.Itoa(n)
	for i := 1; i <= len(num)/2; i++ {
		chunk := num[:i]
		count := len(num) / i
		repeated := strings.Repeat(chunk, count)
		if repeated == num {
			return true
		}
	}

	return false
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
