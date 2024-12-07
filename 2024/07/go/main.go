package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"math"
	"strconv"
)

func main() {

	lines := common.GetLines()
	operatorsp1 := []byte{'*', '+'}
	operatorsp2 := []byte{'*', '+', '|'}

	var part1, part2 int
	for _, line := range lines {
		target, numbersString := common.Split(line, ":")
		numbers := common.ExtractNumbers(numbersString)

		t := common.Must(strconv.Atoi, target)

		if checkLine(t, numbers, operatorsp1) {
			part1 += t
		}

		if checkLine(t, numbers, operatorsp2) {
			part2 += t
		}

	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func checkLine(target int, numbers []int, operators []byte) bool {
	perms := make([]byte, len(numbers)-1)
	for perm := range permutations(perms, operators) {
		if isValid(target, numbers, perm) {
			return true
		}
	}

	return false
}

func isValid(target int, numbers []int, operators []byte) bool {
	current, remainder := numbers[0], numbers[1:]
	for i, num := range remainder {
		switch operators[i] {
		case '*':
			current *= num
		case '+':
			current += num
		case '|':
			current *= int(math.Pow(10, math.Floor(math.Log10(float64(num)))+1))
			current += num
		}

		if current > target {
			return false
		}
	}

	return current == target
}

func permutations(perms []byte, operators []byte) func(func([]byte) bool) {
	return func(yield func([]byte) bool) {
		if len(perms) == 0 {
			yield(perms)
			return
		}

		for _, op := range operators {
			perms[0] = op
			for _ = range permutations(perms[1:], operators) {
				if !yield(perms) {
					return
				}
			}
		}
	}
}
