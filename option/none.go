package option

type none[T any] struct {
}

func (n none[T]) IsSome() bool {
	return false
}

func (n none[T]) IsNone() bool {
	return true
}

func (n none[T]) Value() T {
	panic("can't get value of none")
}

func None[T any]() Option[T] {
	return &none[T]{}
}
