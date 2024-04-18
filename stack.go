package collections

type Stack[T any] interface {
	// Push legt ein Element auf den Stack
	Push(value T)
	// Pop entfernt das oberste Element vom Stack
	Pop() T
}
