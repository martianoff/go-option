package option

type optNone[T any] struct {
}

func (n optNone[T]) Get() T {
	panic("called Get on optNone value")
}

func (n optNone[T]) GetOrElse(v T) T {
	return v
}

func (n optNone[T]) OrElse(opt Option[T]) Option[T] {
	return opt
}

func (n optNone[T]) Empty() bool {
	return true
}

func (n optNone[T]) NonEmpty() bool {
	return false
}

func (n optNone[T]) String() string {
	return "None"
}
