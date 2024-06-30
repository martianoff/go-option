package option

import (
	"reflect"
	"testing"
)

func TestSome_Empty(t *testing.T) {
	type fields struct {
		underlyingValue int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Some[T] Empty() returns false", fields: struct{ underlyingValue int }{underlyingValue: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := optSome[int]{
				underlyingValue: tt.fields.underlyingValue,
			}
			if got := s.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome_Get(t *testing.T) {
	type fields struct {
		underlyingValue int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "Some[T] Get returns underlying value", fields: struct{ underlyingValue int }{underlyingValue: 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := optSome[int]{
				underlyingValue: tt.fields.underlyingValue,
			}
			if got := s.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome_GetOrElse(t *testing.T) {
	type fields struct {
		underlyingValue int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "Some[T] GetOrElse() returns underlying value", fields: struct{ underlyingValue int }{underlyingValue: 2}, args: struct{ v int }{v: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := optSome[int]{
				underlyingValue: tt.fields.underlyingValue,
			}
			if got := s.GetOrElse(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome_NonEmpty(t *testing.T) {
	type fields struct {
		underlyingValue int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "Some[T] NonEmpty() returns true", fields: struct{ underlyingValue int }{underlyingValue: 2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := optSome[int]{
				underlyingValue: tt.fields.underlyingValue,
			}
			if got := s.NonEmpty(); got != tt.want {
				t.Errorf("NonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome_OrElse(t *testing.T) {
	type fields struct {
		underlyingValue int
	}
	type args struct {
		opt Option[int]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Option[int]
	}{
		{name: "Some[T] OrElse() returns original value with another Some", fields: struct{ underlyingValue int }{underlyingValue: 3}, args: args{opt: Some[int](2)}, want: Some[int](3)},
		{name: "Some[T] OrElse() returns original value with another None", fields: struct{ underlyingValue int }{underlyingValue: 3}, args: args{opt: None[int]()}, want: Some[int](3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := optSome[int]{
				underlyingValue: tt.fields.underlyingValue,
			}
			if got := s.OrElse(tt.args.opt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome_String(t *testing.T) {
	type fields struct {
		underlyingValue int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Some[T] String() returns Some(T)", fields: struct{ underlyingValue int }{underlyingValue: 2}, want: "Some(2)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := optSome[int]{
				underlyingValue: tt.fields.underlyingValue,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
