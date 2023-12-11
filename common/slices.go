package common

import "slices"

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

// Pairs creats a list distinc pairs from the provided items. The items must have a length of at least 2
func Pairs[K any](items []K) []Pair[K] {
	var pairs []Pair[K]
	for i, a := range items {
		for _, b := range items[i+1:] {
			pairs = append(pairs, Pair[K]{a, b})
		}
	}

	return pairs
}

func Map[K any, V any](mapper func(K) V, values []K) []V {
	var result []V

	for _, value := range values {
		result = append(result, mapper(value))
	}

	return result
}

func Fold[K any, L any](callable func(previous L, current K) L, items []K) L {
	var current L
	for _, item := range items {
		current = callable(current, item)
	}

	return current
}
