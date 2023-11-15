package option

type some[T any] struct {
	value T
}

func (s some[T]) IsSome() bool {
	return true
}

func (s some[T]) IsNone() bool {
	return false
}

func (s some[T]) Value() T {
	return s.value
}

func Some[T any](value T) Option[T] {
	return &some[T]{
		value: value,
	}
}
