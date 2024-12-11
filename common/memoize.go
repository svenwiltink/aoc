package common

type MemoizeType[A, B comparable, R any] func(MemoizeType[A, B, R], A, B) R

func Memoize[A, B comparable, R any](root MemoizeType[A, B, R], stone A, blinks B) R {
	cache := make(map[struct {
		p1 A
		p2 B
	}]R)

	var inner MemoizeType[A, B, R]
	inner = func(_ MemoizeType[A, B, R], s A, b B) R {
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
	return inner(root, stone, blinks)
}
