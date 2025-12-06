package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"strconv"
	"strings"
)

func main() {
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
