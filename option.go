package option

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
)

type opt[T any] interface {
	Get() T
	GetOrElse(v T) T
	OrElse(opt Option[T]) Option[T]
	Empty() bool
	NonEmpty() bool
	String() string
}

type Option[T any] struct {
	opt[T]
}

// None returns an Option of Type T which represents no value.
func None[T any]() Option[T] {
	return Option[T]{optNone[T]{}}
}

// Some returns an Option of Type T with value 'o'.
func Some[T any](o T) Option[T] {
	return Option[T]{optSome[T]{o}}
}

// NewOption creates an Option, initializing it with value 'o'.
// If 'o' is a nil pointer, map, slice, func, interface or channel, the Option will represent no value.
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

// NewOptionFromPointer creates an Option from input pointer 'o'.
// If 'o' is nil, it returns an Option representing no value.
func NewOptionFromPointer[T any](o *T) Option[T] {
	if o == nil {
		return None[T]()
	}
	return Some[T](*o)
}

// Map applies the function 'mapper' to the Option 'opt' if it is a Some and returns a new Option.
// If 'opt' is a None, the function 'mapper' is not applied.
func Map[T1, T2 any](opt Option[T1], mapper func(T1) T2) Option[T2] {
	if opt.NonEmpty() {
		return Option[T2]{optSome[T2]{mapper(opt.Get())}}
	} else {
		return Option[T2]{optNone[T2]{}}
	}
}

// FlatMap applies the function 'mapper' to the Option 'opt' if it is a Some and returns a new Option.
// If 'opt' is a None, the function 'mapper' is not applied and Option representing no value is returned.
func FlatMap[T1, T2 any](opt Option[T1], mapper func(T1) Option[T2]) Option[T2] {
	if opt.NonEmpty() {
		return mapper(opt.Get())
	} else {
		return Option[T2]{optNone[T2]{}}
	}
}

// Match takes in a current Option 'opt', a function 'onSome' and 'onNone'.
// 'onSome' gets executed when 'opt' is Some and 'onNone' gets executed when 'opt' is None.
func Match[T1, T2 any](opt Option[T1], onSome func(T1) T2, onNone func() T2) T2 {
	if opt.NonEmpty() {
		return onSome(opt.Get())
	} else {
		return onNone()
	}
}

// Equal checks if the calling Option 'x' and the provided Option 'y' are equal.
// Returns true if both Options are None or both Options are Some and hold the same value.
func (x Option[T]) Equal(y Option[T]) bool {
	if x.Empty() && y.Empty() {
		return true
	} else if !x.Empty() && !y.Empty() {
		return cmp.Equal(x.Get(), y.Get())
	}
	return false
}
