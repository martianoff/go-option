package option

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
)

type optionalValue[T any] interface {
	Get() T
	GetOrElse(v T) T
	OrElse(opt Option[T]) Option[T]
	Empty() bool
	NonEmpty() bool
	String() string
}

type Option[T any] struct {
	optionalValue[T]
}

func None[T any]() Option[T] {
	return Option[T]{optNone[T]{}}
}

func Some[T any](o T) Option[T] {
	return Option[T]{optSome[T]{o}}
}

func NewOption[T any](o T) Option[T] {
	if v := reflect.ValueOf(o); (v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func) && v.IsNil() {
		return Option[T]{optNone[T]{}}
	}
	return Option[T]{optSome[T]{o}}
}

func NewOptionFromPointer[T any](o *T) Option[T] {
	if o == nil {
		return None[T]()
	}
	return Some[T](*o)
}

func Map[T1, T2 any](opt Option[T1], mapper func(T1) T2) Option[T2] {
	if opt.NonEmpty() {
		return Option[T2]{optSome[T2]{mapper(opt.Get())}}
	} else {
		return Option[T2]{optNone[T2]{}}
	}
}

func FlatMap[T1, T2 any](opt Option[T1], mapper func(T1) Option[T2]) Option[T2] {
	if opt.NonEmpty() {
		return mapper(opt.Get())
	} else {
		return Option[T2]{optNone[T2]{}}
	}
}

func (x Option[T]) Equal(y Option[T]) bool {
	if x.Empty() && y.Empty() {
		return true
	} else if !x.Empty() && !y.Empty() {
		return cmp.Equal(x.Get(), y.Get())
	}
	return false
}
