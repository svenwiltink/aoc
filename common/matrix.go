package common

import "slices"

func Transpose[e any](a [][]e) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a[i][j], a[j][i] = a[j][i], a[i][j]
		}
	}
}

func FlipVert[e any](a [][]e) {
	for _, line := range a {
		slices.Reverse(line)
	}
}
