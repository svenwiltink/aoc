package common

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
