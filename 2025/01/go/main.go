package main

import (
	"fmt"
	"github.com/svenwiltink/aoc/common"
	"math"
)

func main() {
	moves := getMoves()
	d := Dial{position: 50}
	for _, move := range moves {
		d.rotate(move)
	}

	fmt.Println(d.zeroes, d.clicks)
}

type Dial struct {
	position int

	clicks int
	zeroes int
}

func (d *Dial) rotate(n int) {
	revolutions := (int)(math.Abs(float64(n))) / 100
	d.clicks += revolutions

	wasZero := d.position == 0
	d.position += n % 100

	if d.position < 0 {
		d.position += 100
		if !wasZero {
			d.clicks++
		}
	}

	if d.position > 100 && !wasZero {
		d.clicks++
	}

	d.position %= 100

	if d.position == 0 {
		d.clicks++
		d.zeroes++
	}
}

func getMoves() []int {
	lines := common.GetLines()

	var moves []int
	for _, line := range lines {
		var dir rune
		var num int
		_, err := fmt.Sscanf(line, "%c%d", &dir, &num)
		if err != nil {
			panic(err)
		}

		if dir == 'L' {
			num *= -1
		}

		moves = append(moves, num)
	}

	return moves
}
