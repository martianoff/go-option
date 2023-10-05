package option

import (
	"fmt"
)

type optSome[T any] struct {
	underlyingValue T
}

func (s optSome[T]) Get() T {
	return s.underlyingValue
}

func (s optSome[T]) GetOrElse(v T) T {
	return s.underlyingValue
}

func (s optSome[T]) OrElse(opt Option[T]) Option[T] {
	return s
}

func (s optSome[T]) Empty() bool {
	return false
}

func (s optSome[T]) NonEmpty() bool {
	return true
}

func (s optSome[T]) String() string {
	return fmt.Sprintf("Some(%v)", s.underlyingValue)
}
