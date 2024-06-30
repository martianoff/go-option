package option

import (
	"reflect"
	"strconv"
	"testing"
)

func TestFlatMap(t *testing.T) {
	type args struct {
		opt    Option[Option[bool]]
		mapper func(Option[bool]) Option[bool]
	}
	tests := []struct {
		name string
		args args
		want Option[bool]
	}{
		{
			name: "FlatMap with Some[Some[true]] get converted to Some[true]",
			args: args{
				opt: Some[Option[bool]](
					Some[bool](true),
				),
				mapper: func(o1 Option[bool]) Option[bool] {
					return o1
				},
			},
			want: Some[bool](true),
		},
		{
			name: "FlatMap with Some[None] get converted to None",
			args: args{
				opt: Some[Option[bool]](
					None[bool](),
				),
				mapper: func(o1 Option[bool]) Option[bool] {
					return o1
				},
			},
			want: None[bool](),
		},
		{
			name: "FlatMap with None[Option[true]] get converted to None",
			args: args{
				opt: None[Option[bool]](),
				mapper: func(o1 Option[bool]) Option[bool] {
					return o1
				},
			},
			want: None[bool](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatMap(tt.args.opt, tt.args.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		opt    Option[int]
		mapper func(int) string
	}
	tests := []struct {
		name string
		args args
		want Option[string]
	}{
		{
			name: "Map with Some[int] get converted to Some[string]",
			args: args{
				opt: Some[int](4),
				mapper: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: Some[string]("4"),
		},
		{
			name: "Map with None[int] get converted to None[string]",
			args: args{
				opt: None[int](),
				mapper: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: None[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.opt, tt.args.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToOption(t *testing.T) {
	type args struct {
		v map[int]string
	}
	tests := []struct {
		name string
		args args
		want Option[map[int]string]
	}{
		{
			name: "NewOption with value is converted to Some(value)",
			args: args{v: map[int]string{}},
			want: Some[map[int]string](map[int]string{}),
		},
		{
			name: "NewOption without value is converted to None",
			args: args{v: nil},
			want: None[map[int]string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOption(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOptionFromPointer(t *testing.T) {
	type T = int

	cases := []struct {
		name     string
		input    *T
		expected Option[T]
	}{
		{
			name:     "Test with nil pointer should return None",
			input:    nil,
			expected: None[T](),
		},
		{
			name: "Test with valid pointer should return Some",
			input: func() *T {
				var v T = 5
				return &v
			}(),
			expected: Some[T](5),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := NewOptionFromPointer(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestOptionEqual(t *testing.T) {
	type T = int

	cases := []struct {
		name     string
		x        Option[T]
		y        Option[T]
		expected bool
	}{
		{
			name:     "Both options are None, should be equal",
			x:        None[T](),
			y:        None[T](),
			expected: true,
		},
		{
			name:     "One is None, one is Some, should not be equal",
			x:        None[T](),
			y:        Some[T](1),
			expected: false,
		},
		{
			name:     "Both are Some and same value, should be equal",
			x:        Some[T](1),
			y:        Some[T](1),
			expected: true,
		},
		{
			name:     "Both are Some but different values, should not be equal",
			x:        Some[T](1),
			y:        Some[T](2),
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.x.Equal(tc.y)
			if actual != tc.expected {
				t.Errorf("Equal() = %v, want %v", actual, tc.expected)
			}
		})
	}
}
