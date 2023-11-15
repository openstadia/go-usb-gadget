package option

type Option[T any] interface {
	IsSome() bool
	IsNone() bool
	Value() T
}
