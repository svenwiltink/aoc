package common

import (
	"slices"
	"strconv"
	"strings"
)

func ExtractNumbers(input string) []int {
	numbers := strings.Fields(input)
	return Map(MustClosure(strconv.Atoi), numbers)
}

func Fold[K any, L any](callable func(previous L, current K) L, items []K, start L) L {
	current := start
	for _, item := range items {
		current = callable(current, item)
	}

	return current
}

func MustClosure[K any, L any](callable func(L) (K, error)) func(L) K {
	return func(l L) K {
		output, err := callable(l)
		if err != nil {
			panic(err)
		}

		return output
	}
}

func Map[K any, V any](mapper func(K) V, values []K) []V {
	var result []V

	for _, value := range values {
		result = append(result, mapper(value))
	}

	return result
}

// Intersect gets all item of B that are also in A
func Intersect[K comparable](A, B []K) []K {
	var result []K
	for _, item := range B {
		if slices.Contains(A, item) {
			result = append(result, item)
		}
	}

	return result
}
