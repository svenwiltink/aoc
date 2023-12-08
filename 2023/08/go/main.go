package main

import (
	"fmt"
	"strings"

	"github.com/svenwiltink/aoc/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	instructions, nodes := parseInput()
	current := nodes["AAA"]
	fmt.Println(traverse(instructions, current, nodes))
}

func part2() {
	instructions, nodes := parseInput()
	var steps []int

	for _, node := range nodes {
		if strings.HasSuffix(node.Name, "A") {
			steps = append(steps, traverse(instructions, node, nodes))
		}
	}

	fmt.Println(common.Lcm(steps...))
}

func traverse(instructions string, current Node, nodes map[string]Node) int {
	steps := 0
	for {
		if strings.HasSuffix(current.Name, "Z") {
			return steps
		}

		switch instructions[steps%len(instructions)] {
		case 'R':
			current = nodes[current.Right]
		case 'L':
			current = nodes[current.Left]
		}

		steps++
	}
}

func parseInput() (string, map[string]Node) {
	lines := common.GetLines()
	instructions := lines[0]

	nodes := make(map[string]Node)
	for _, line := range lines[2:] {
		var node Node
		_, err := fmt.Sscanf(line, "%s = (%3s, %3s)", &node.Name, &node.Left, &node.Right)
		if err != nil {
			panic(err)
		}

		nodes[node.Name] = node
	}

	return instructions, nodes
}

type Node struct {
	Name  string
	Left  string
	Right string
}
