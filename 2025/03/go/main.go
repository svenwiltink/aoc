package main

import (
	"fmt"
	. "github.com/svenwiltink/aoc/common"
)

func main() {
	lines := GetLines()
	numbers := Map(Curry(ExtractNumbersSep, ""), lines)
	fmt.Println(Sum(Map(Curry(getDigit, 2), numbers)))
	fmt.Println(Sum(Map(Curry(getDigit, 12), numbers)))
}

func getDigit(count int, numbers []int) int {
	result := 0
	for i := 1; i <= count; i++ {
		index, digit := Max(numbers[:len(numbers)-(count-i)])
		numbers = numbers[index+1:]
		result *= 10
		result += digit
	}
	return result
}
