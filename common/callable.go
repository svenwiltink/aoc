package common

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

func Must[K any, L any](callable func(L) (K, error), l L) K {
	output, err := callable(l)
	if err != nil {
		panic(err)
	}

	return output
}

func Map[K any, V any](mapper func(K) V, values []K) []V {
	var result []V

	for _, value := range values {
		result = append(result, mapper(value))
	}

	return result
}
