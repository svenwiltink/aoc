package common

import (
	"slices"
)

func TransposeInPlace[e any](a [][]e) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}
}

func Transpose[e any](slice [][]e) [][]e {
	xl := len(slice[0])
	yl := len(slice)

	result := make([][]e, xl)

	for i := range result {
		result[i] = make([]e, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func FlipVert[e any](a [][]e) {
	for _, line := range a {
		slices.Reverse(line)
	}
}
