package common

// Curry the first parameter of the function. This may have been a mistake but we're here for the experiment
func Curry[A, B, C any](f func(a A, b B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}

// Curry the first parameter of the function. This may have been a mistake but we're here for the experiment
func Curry2[A, B, C, D any](f func(a A, b B, c C) D, a A) func(B, C) D {
	return func(b B, c C) D {
		return f(a, b, c)
	}
}
