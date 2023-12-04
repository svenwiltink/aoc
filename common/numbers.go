package common

import (
	"strconv"
	"strings"
)

func ExtractNumbers(input string) []int {
	numbers := strings.Fields(input)
	return Map(MustClosure(strconv.Atoi), numbers)
}

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}

	return sum
}
