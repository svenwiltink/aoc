package common

func Memoize2[A, B comparable, R any](root func(func(A, B) R, A, B) R, a A, b B) R {
	cache := make(map[struct {
		p1 A
		p2 B
	}]R)

	var inner func(A, B) R
	inner = func(s A, b B) R {
		cacheKey := struct {
			p1 A
			p2 B
		}{s, b}
		if item, exists := cache[cacheKey]; exists {
			return item
		}

		result := root(inner, s, b)
		cache[cacheKey] = result
		return result
	}
	return inner(a, b)
}
