package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/svenwiltink/aoc/common"
)

var (
	lookup = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := common.GetLines()
	numbers := extractNumbers(lines)
	fmt.Println(common.Sum(numbers))
}

func part2() {
	lines := common.GetLines()
	lines = replaceAlphaNumbers(lines)
	numbers := extractNumbers(lines)
	fmt.Println(common.Sum(numbers))
}

func replaceAlphaNumbers(lines []string) []string {
	var output []string
	for _, line := range lines {
		output = append(output, replaceAlphaNumber(line))
	}

	return output
}

func replaceAlphaNumber(line string) string {
	for index, word := range lookup {
		line = strings.ReplaceAll(line, word, fmt.Sprintf("%s%d%s", word, index+1, word))
	}

	return line
}

func extractNumbers(lines []string) []int {
	var numbers []int
	for _, line := range lines {
		numbers = append(numbers, extractNumber(line))
	}

	return numbers
}

func extractNumber(line string) int {
	var digits []rune
	for _, digit := range line {
		if unicode.IsNumber(digit) {
			digits = append(digits, digit)
		}
	}

	num, err := strconv.Atoi(string([]rune{digits[0], digits[len(digits)-1]}))
	if err != nil {
		panic(err)
	}
	return num
}
