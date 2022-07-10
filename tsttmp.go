package gotesttmp

type Client interface {
	Inc(a, b int) int
}

// Inc returns sum of "a" and "b", it's an unexpected name but we just want to
// break the contract of V1.
func Inc(a, b int) int {
	return a + b
}

func Dec(i int) int {
	return i - 1
}
