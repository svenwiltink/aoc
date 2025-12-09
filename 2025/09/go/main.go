package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
)

func main() {
	lines := common.GetLines()
	points := common.Map(func(s string) common.Coords {
		var x, y int
		_, err := fmt.Sscanf(s, "%d,%d", &x, &y)
		if err != nil {
			panic(err)
		}
		return common.Coords{x, y}
	}, lines)

	rectangles := common.Pairs(points)
	sizes := common.Map(size, rectangles)
	_, value := common.Max(sizes)
	fmt.Println(value)

	current, points := points[0], append(points[1:], points[0])

	var edges []common.Pair[common.Coords]
	for _, point := range points {
		edges = append(edges, common.Pair[common.Coords]{current, point})
		current = point
	}

	var largest int
	for _, r := range rectangles {
		if isValid(r, edges) {
			area := size(r)
			if area > largest {
				largest = area
			}
		}
	}

	fmt.Println("largest", largest)
}

func isValid(rectangle common.Pair[common.Coords], edges []common.Pair[common.Coords]) bool {

	for _, edge := range edges {
		if intersects(rectangle, edge) {
			return false
		}
	}
	return true
}

// based on https://kishimotostudios.com/articles/aabb_collision/
func intersects(rectangle common.Pair[common.Coords], edge common.Pair[common.Coords]) bool {

	eMinX, eMaxX := min(edge[0][0], edge[1][0]), max(edge[0][0], edge[1][0])
	eMinY, eMaxY := min(edge[0][1], edge[1][1]), max(edge[0][1], edge[1][1])

	rMinX, rMaxX := min(rectangle[0][0], rectangle[1][0]), max(rectangle[0][0], rectangle[1][0])
	rMinY, rMaxY := min(rectangle[0][1], rectangle[1][1]), max(rectangle[0][1], rectangle[1][1])

	EdgeRightOfRec := eMinX >= rMaxX
	EdgeLeftOfRec := eMaxX <= rMinX
	EdgeAboveRec := eMaxY <= rMinY
	EdgeBelowRec := eMinY >= rMaxY
	return !(EdgeRightOfRec || EdgeLeftOfRec || EdgeAboveRec || EdgeBelowRec)
}

func size(pair common.Pair[common.Coords]) int {
	xd := pair[0].Sub(pair[1])

	return (common.Abs(xd[0]) + 1) * (common.Abs(xd[1]) + 1)
}
