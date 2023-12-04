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
