package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	// part1()
	part2()
}

func part1() {
	line := common.GetLines()[0]
	parts := strings.Split(line, ",")
	numbers := common.Map(hash, parts)
	fmt.Println(common.Sum(numbers))
}

type lens struct {
	label string
	focal int
}

func (l lens) Equal(o lens) bool {
	return l.label == o.label
}

type box []lens

func (b *box) add(l lens) {
	index := slices.IndexFunc(*b, l.Equal)

	if index == -1 {
		*b = append(*b, l)
		return
	}

	(*b)[index] = l
}

func (b *box) remove(l lens) {
	index := slices.IndexFunc(*b, l.Equal)

	if index == -1 {
		return
	}

	(*b) = slices.Delete(*b, index, index+1)
}

func (b box) power() int {
	var total int
	for i, lens := range b {
		total += (i + 1) * lens.focal
	}

	return total
}

type boxes [256]box

func (b *boxes) handleInstruction(instruction string) {
	switch instruction[len(instruction)-1] {
	case '-':
		label := instruction[:len(instruction)-1]
		h := hash(label)
		l := lens{label, 0}

		b[h].remove(l)
	default:
		label := instruction[:len(instruction)-2]
		h := hash(label)
		focal := int(instruction[len(instruction)-1] - '0')

		l := lens{label, focal}
		b[h].add(l)
	}
}

func part2() {
	line := common.GetLines()[0]
	parts := strings.Split(line, ",")

	var b boxes
	for _, instruction := range parts {
		b.handleInstruction(instruction)
	}

	var total int
	for i, box := range b {
		total += (i + 1) * box.power()
	}

	fmt.Println(total)
}

func hash(part string) int {
	var num int
	for _, char := range part {
		num += int(char)
		num *= 17
		num = num % 256
	}

	return num
}
