package common

import (
	"slices"
)

type Coords [2]int

func (c Coords) Add(o Coords) Coords {
	return Coords{c[0] + o[0], c[1] + o[1]}
}

func (c Coords) IsDiagonal() bool {
	return c[0] != 0 && c[1] != 0
}

func (c Coords) IsZero() bool {
	return c[0] == 0 && c[1] == 0
}

func (c Coords) Rotate90(clockWise bool) Coords {
	index := slices.Index(Directions4, c) + 4
	if clockWise {
		index++
	} else {
		index--
	}
	index = index % 4
	return Directions4[index]
}

func (c Coords) Reverse() Coords {
	return Coords{c[0] * -1, c[1] * -1}
}

var (
	N  = Coords{0, -1}
	NE = Coords{1, -1}
	E  = Coords{1, 0}
	SE = Coords{1, 1}
	S  = Coords{0, 1}
	SW = Coords{-1, 1}
	W  = Coords{-1, 0}
	NW = Coords{-1, -1}
)

var Directions4 = []Coords{
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
	{0, -1}, // up
}

var Directions8 = []Coords{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}
