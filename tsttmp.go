package gotesttmp

type Client interface {
	Inc(int) int
}

func Inc(i int) int {
	return i + 1
}
