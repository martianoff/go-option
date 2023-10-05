package option

import "reflect"

type Option[T any] interface {
	Get() T
	GetOrElse(v T) T
	OrElse(opt Option[T]) Option[T]
	Empty() bool
	NonEmpty() bool
	String() string
}

func None[T any]() Option[T] {
	return optNone[T]{}
}

func Some[T any](o T) Option[T] {
	return NewOption(o)
}

func NewOption[T any](o T) Option[T] {
	if v := reflect.ValueOf(o); (v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func) && v.IsNil() {
		return optNone[T]{}
	}
	return optSome[T]{o}
}

func Map[T1, T2 any](opt Option[T1], mapper func(T1) T2) Option[T2] {
	if opt.NonEmpty() {
		return optSome[T2]{mapper(opt.Get())}
	} else {
		return optNone[T2]{}
	}
}

func FlatMap[T1, T2 any](opt Option[T1], mapper func(T1) Option[T2]) Option[T2] {
	if opt.NonEmpty() {
		return mapper(opt.Get())
	} else {
		return optNone[T2]{}
	}
}
