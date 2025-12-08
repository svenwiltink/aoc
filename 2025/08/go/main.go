package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"math"
	"slices"
)

func main() {
	part1()
	part2()
}

func part2() {
	coords := getCoords()
	edges := common.Pairs(coords)
	slices.SortFunc(edges, func(e common.Pair[coord3d], e2 common.Pair[coord3d]) int {
		return int(e[0].distanceSquared(e[1]) - e2[0].distanceSquared(e2[1]))
	})

	var currentEdges []common.Pair[coord3d]
	for _, edge := range edges {
		currentEdges = append(currentEdges, edge)
		g := graph{
			edges:  currentEdges,
			lookup: make(map[coord3d][]coord3d),
		}

		g.constructLookup()
		clusterSize := g.getCluster()
		if clusterSize == len(coords) {
			fmt.Println(edge[0][0] * edge[1][0])
			break
		}
	}
}

func part1() {
	coords := getCoords()
	edges := common.Pairs(coords)
	slices.SortFunc(edges, func(e common.Pair[coord3d], e2 common.Pair[coord3d]) int {
		return int(e[0].distanceSquared(e[1]) - e2[0].distanceSquared(e2[1]))
	})

	top10 := edges[:1000]

	g := graph{
		edges:  top10,
		lookup: map[coord3d][]coord3d{},
	}

	g.constructLookup()
	var clusters []int
	for len(g.lookup) > 0 {
		clusters = append(clusters, g.getCluster())
	}

	slices.Sort(clusters)
	slices.Reverse(clusters)
	result := clusters[0]
	for _, num := range clusters[1:3] {
		result *= num
	}

	fmt.Println(result)
}

type graph struct {
	edges []common.Pair[coord3d]

	lookup map[coord3d][]coord3d
}

func (g *graph) getCluster() int {
	var start coord3d
	for e := range g.lookup {
		start = e
		break
	}

	visited := map[coord3d]struct{}{}
	stack := []coord3d{start}
	for {
		if len(stack) == 0 {
			break
		}
		var current coord3d
		stack, current = stack[1:], stack[0]
		visited[current] = struct{}{}

		for _, connected := range g.lookup[current] {
			if _, v := visited[connected]; v {
				continue
			}
			stack = append(stack, connected)
		}

		// nothing else can connect to this
		delete(g.lookup, current)
	}

	return len(visited)
}

func (g *graph) constructLookup() {
	for _, edge := range g.edges {
		g.lookup[edge[0]] = append(g.lookup[edge[0]], edge[1])
		g.lookup[edge[1]] = append(g.lookup[edge[1]], edge[0])
	}
}

type coord3d [3]int

func (c coord3d) distanceSquared(o coord3d) float64 {
	return math.Pow(float64(o[0]-c[0]), 2) +
		math.Pow(float64(o[1]-c[1]), 2) +
		math.Pow(float64(o[2]-c[2]), 2)
}

func coordFromString(s string) coord3d {
	var x, y, z int
	_, err := fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z)
	if err != nil {
		panic(err)
	}

	return coord3d{x, y, z}
}
func getCoords() []coord3d {
	lines := common.GetLines()
	return common.Map(coordFromString, lines)
}
