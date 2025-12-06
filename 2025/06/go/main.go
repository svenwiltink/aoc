package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	lines := common.GetLines()
	matrix := common.Map(func(k string) []rune {
		return []rune(k)
	}, lines)

	matrix = common.Transpose(matrix)
	start := true
	operation := '0'
	result := 0
	row := 0
	for _, l := range matrix {
		// row is empty, add result
		oString := strings.TrimSpace(string(l))
		if len(oString) == 0 {
			start = true
			result += row
			row = 0
			continue
		}

		if start {
			start = false
			l, operation = l[:len(l)-1], l[len(l)-1]
			var err error
			row, err = strconv.Atoi(strings.TrimSpace(string(l)))
			if err != nil {
				panic(err)
			}
			continue
		}

		num := common.Must(strconv.Atoi, strings.TrimSpace(string(l)))

		switch operation {
		case '*':
			row *= num
		case '+':
			row += num
		}

		start = false
	}

	result += row
	fmt.Println(result)
}

func part1() {
	lines := common.GetLines()
	matrix := common.Map(strings.Fields, lines)
	matrix = common.Transpose(matrix)
	fmt.Println(common.Sum(common.Map(value, matrix)))
}

func value(line []string) int {
	values, operation := line[:len(line)-1], line[len(line)-1]
	numbers := common.Map(common.MustClosure(strconv.Atoi), values)
	return calculate(numbers, operation)
}

func calculate(nums []int, operation string) int {
	result := nums[0]

	for _, num := range nums[1:] {
		switch operation {
		case "+":
			result += num
		case "*":
			result *= num
		}
	}

	return result
}
