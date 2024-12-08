package common

import (
	"iter"
	"slices"
)

func GenerateCombinations[T any](src []T, size int) iter.Seq[[]T] {
	var res = make([]T, size)
	var generate func(int) bool

	return func(yield func([]T) bool) {
		generate = func(idx int) bool {
			if idx == size {
				if !yield(res) {
					return false
				}
				return true
			}

			for _, v := range src {
				res[idx] = v
				if !generate(idx + 1) {
					return false
				}
			}
			return true
		}

		generate(0)
	}
}

func Permutations[T any](src []T) iter.Seq[[]T] {
	dst := make([]T, len(src))
	return permutations(src, dst)
}

func permutations[T any](src []T, dst []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {

		if len(src) == 0 {
			yield(src)
			return
		}
		for i, item := range src {
			dst[0] = item
			remainder := slices.Clone(src)
			remainder = remainder[:i+copy(remainder[i:], remainder[i+1:])]
			for _ = range permutations(remainder, dst[1:]) {
				if !yield(dst) {
					return
				}
			}
		}
	}
}
