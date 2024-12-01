package internal

type TestData[T any] struct {
	Name  string
	Input string
	Want  T
}
