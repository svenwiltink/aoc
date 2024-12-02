package common

import (
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

func ExtractNumbers(input string) []int {
	numbers := strings.Fields(input)
	return Map(MustClosure(strconv.Atoi), numbers)
}

func ExtractNumbersSep(input, sep string) []int {
	numbers := strings.Split(input, sep)
	return Map(MustClosure(strconv.Atoi), numbers)
}

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}

	return a
}

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
}

// greatest common divisor (GCD) via Euclidean algorithm
func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func Lcm(integers ...int) int {
	result := integers[0] * integers[1] / Gcd(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = Lcm(result, integers[i])
	}

	return result
}
