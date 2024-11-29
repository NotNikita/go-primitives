package structures

type StructMethods[T comparable] interface {
	Push(T) []T
	Pop() (T, []T)
	Size() T
	Empty() bool
	// Front() T
	// Back() T
}

// Supported Structures
/*
1. Queue
*/
