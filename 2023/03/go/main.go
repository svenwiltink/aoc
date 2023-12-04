package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := common.GetLines()
	matrix := common.Map(func(item string) []rune {
		return []rune(item)
	}, lines)

	var numbers []int

	for x, line := range matrix {
		started := false
		start := 0
		end := 0
		matches := false
		for y, value := range line {
			if unicode.IsDigit(value) {
				if !started {
					started, start = true, y
				}

				if !matches { // no need to do this more often
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							newx := x + dx
							newy := y + dy

							// bounds checking
							if newx < 0 || newx > len(matrix)-1 || newy < 0 || newy > len(matrix[newx])-1 {
								continue
							}

							char := matrix[newx][newy]
							if !unicode.IsDigit(char) && char != '.' {
								matches = true
							}
						}
					}
				}

				if y != len(line)-1 {
					continue
				}

				y++
			}

			if !started {
				continue
			}

			end = y

			if matches {
				numbers = append(numbers, common.MustClosure(strconv.Atoi)(string(line[start:end])))
			}

			started = false
			matches = false

		}

	}

	fmt.Println(common.Sum(numbers))
}

func part2() {
	lines := common.GetLines()
	matrix := common.Map(func(item string) []rune {
		return []rune(item)
	}, lines)

	var cache [141][141]int
	var sum int

	for x, line := range matrix {
		started := false
		start := 0
		end := 0
		matches := false
		for y, value := range line {
			if unicode.IsDigit(value) {
				if !started {
					started, start = true, y
				}

				if !matches { // no need to do this more often
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							newx := x + dx
							newy := y + dy

							// bounds checking
							if newx < 0 || newx > len(matrix)-1 || newy < 0 || newy > len(matrix[newx])-1 {
								continue
							}

							char := matrix[newx][newy]
							if !unicode.IsDigit(char) && char != '.' {
								matches = true
							}
						}
					}
				}

				if y != len(line)-1 {
					continue
				}

				y++
			}

			if !started {
				continue
			}

			end = y

			if matches {
				for tmpy := start; tmpy < end; tmpy++ {
					cache[x][tmpy] = common.Must(strconv.Atoi, string(line[start:end]))
				}
			}

			started = false
			matches = false
		}
	}

	for x, line := range matrix {
		for y, value := range line {
			if value != '*' {
				continue
			}

			var matches int
			var ratio int
			var prev int
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					newx := x + dx
					newy := y + dy

					// bounds checking
					if newx < 0 || newx > len(matrix)-1 || newy < 0 || newy > len(matrix[newx])-1 {
						continue
					}

					num := cache[newx][newy]
					if num > 0 && num != prev {
						matches++

						if ratio == 0 {
							ratio = num
						} else {
							ratio *= num
						}
					}

					prev = num
				}
			}

			if matches == 2 {
				sum += ratio
			}
		}
	}

	fmt.Println(sum)
}
