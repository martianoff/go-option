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
			name: "FlatMap with optSome[optSome[true]] get converted to optSome[true]",
			args: args{
				opt: optSome[Option[bool]]{
					optSome[bool]{true},
				},
				mapper: func(o1 Option[bool]) Option[bool] {
					return o1
				},
			},
			want: optSome[bool]{true},
		},
		{
			name: "FlatMap with optSome[optNone] get converted to optNone",
			args: args{
				opt: optSome[Option[bool]]{
					optNone[bool]{},
				},
				mapper: func(o1 Option[bool]) Option[bool] {
					return o1
				},
			},
			want: optNone[bool]{},
		},
		{
			name: "FlatMap with optNone[Option[true]] get converted to optNone",
			args: args{
				opt: optNone[Option[bool]]{},
				mapper: func(o1 Option[bool]) Option[bool] {
					return o1
				},
			},
			want: optNone[bool]{},
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
			name: "Map with optSome[int] get converted to optSome[string]",
			args: args{
				opt: optSome[int]{4},
				mapper: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: optSome[string]{"4"},
		},
		{
			name: "Map with optNone[int] get converted to optNone[string]",
			args: args{
				opt: optNone[int]{},
				mapper: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: optNone[string]{},
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
			name: "NewOption with value is converted to optSome(value)",
			args: args{v: map[int]string{}},
			want: optSome[map[int]string]{map[int]string{}},
		},
		{
			name: "NewOption without value is converted to optNone",
			args: args{v: nil},
			want: optNone[map[int]string]{},
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
