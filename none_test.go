package option

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNone_Empty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "None[T] Empty() returns true", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := None[int]()
			if got := n.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone_Get(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "None[T] Get throws an exception"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := None[int]()
			assert.Panics(t, func() { n.Get() })
		})
	}
}

func TestNone_GetOrElse(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "None[T] GetOrElse() returns else value", args: args{v: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := None[int]()
			if got := n.GetOrElse(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone_NonEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "None[T] Empty() returns false", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := None[int]()
			if got := n.NonEmpty(); got != tt.want {
				t.Errorf("NonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone_OrElse(t *testing.T) {
	type args struct {
		opt Option[int]
	}
	tests := []struct {
		name string
		args args
		want Option[int]
	}{
		{name: "None[T] OrElse() returns None if else condition is None", args: args{opt: None[int]()}, want: None[int]()},
		{name: "None[T] OrElse() returns Some if else condition is Some", args: args{opt: Some[int](2)}, want: Some[int](2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := None[int]()
			if got := n.OrElse(tt.args.opt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNone_String(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "None[T] String() returns None", want: "None"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := None[int]()
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
