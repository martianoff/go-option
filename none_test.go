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
		{name: "optNone[T] Empty() returns true", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := optNone[int]{}
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
		{name: "optNone[T] Get throws an exception"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := optNone[int]{}
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
		{name: "optNone[T] GetOrElse() returns else value", args: args{v: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := optNone[int]{}
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
		{name: "optNone[T] Empty() returns false", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := optNone[int]{}
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
		{name: "optNone[T] OrElse() returns optNone if else condition is optNone", args: args{opt: optNone[int]{}}, want: optNone[int]{}},
		{name: "optNone[T] OrElse() returns optSome if else condition is optSome", args: args{opt: optSome[int]{2}}, want: optSome[int]{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := optNone[int]{}
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
		{name: "optNone[T] String() returns None", want: "None"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := optNone[int]{}
			if got := n.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
