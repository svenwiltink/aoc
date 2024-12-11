package common

func Memoize2[A, B comparable, O any, F func(A, B) O](fn func(F, A, B) O) func(A, B) O {
	type key struct {
		input1 A
		input2 B
	}

	cache := make(map[key]O)

	var f F
	f = func(input1 A, input2 B) O {
		k := key{input1, input2}
		if result, found := cache[k]; found {
			return result
		}

		result := fn(f, input1, input2)
		cache[k] = result
		return result
	}
	return f
}
