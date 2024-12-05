package common

// this was a mistake but had to try
func Curry[A, B, C any](f func(a A, b B) C, b B) func(A) C {
	return func(input A) C {
		return f(input, b)
	}
}
