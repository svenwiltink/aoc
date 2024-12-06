package common

// Curry the first parameter of the function. This may have been a mistake but we're here for the experiment
func Curry[A, B, C any](f func(a A, b B) C, a A) func(B) C {
	return func(b B) C {
		return f(a, b)
	}
}
