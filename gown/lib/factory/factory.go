package factory

type Factory[T any] interface {
	Create() T
}
