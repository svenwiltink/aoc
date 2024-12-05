package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"slices"
	"strings"
)

func main() {
	input := common.GetInput()
	rulesSection, updatesSection := common.Split(input, "\n\n")

	rules := common.Map(common.Curry(common.ExtractNumbersSep, "|"), strings.Split(rulesSection, "\n"))
	updates := common.Map(common.Curry(common.ExtractNumbersSep, ","), strings.Split(updatesSection, "\n"))

	part1(rules, updates)
	part2(rules, updates)
}

func part1(rules [][]int, updates [][]int) {
	var total int
	for _, update := range updates {
		if isValid(update, rules) {
			total += update[len(update)/2]
		}
	}

	fmt.Println(total)
}

func part2(rules [][]int, updates [][]int) {
	lookup := make(map[[2]int]bool)
	for _, rule := range rules {
		lookup[[2]int{rule[0], rule[1]}] = true
	}

	var total int
	invalid := common.Filter(common.Curry(isInvalid, rules), updates)
	for _, update := range invalid {
		slices.SortFunc(update, func(a, b int) int {
			if lookup[[2]int{a, b}] {
				return -1
			}
			return 1
		})
		total += update[len(update)/2]
	}

	fmt.Println(total)
}

func isInvalid(update []int, rules [][]int) bool {
	return !isValid(update, rules)
}

func isValid(update []int, rules [][]int) bool {
	for _, rule := range rules {
		left, right := rule[0], rule[1]

		leftIndex := slices.Index(update, left)
		rightIndex := slices.Index(update, right)

		if leftIndex == -1 || rightIndex == -1 {
			continue
		}

		if rightIndex < leftIndex {
			return false
		}
	}

	return true
}
