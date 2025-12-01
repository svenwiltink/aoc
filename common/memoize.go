package common

func Memoize[A comparable, O any, F func(A) O](fn func(F, A) O) func(A) O {
	cache := make(map[A]O)

	var f F
	f = func(input1 A) O {
		if result, found := cache[input1]; found {
			return result
		}

		result := fn(f, input1)
		cache[input1] = result
		return result
	}
	return f
}

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
