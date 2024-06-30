package option

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

type T = int

func TestOption_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name        string
		jsonStr     string
		expected    Option[T]
		expectedErr bool
	}{
		{
			name:        "Test UnmarshalJSON with null value",
			jsonStr:     "null",
			expected:    None[T](),
			expectedErr: false,
		},
		{
			name:        "Test UnmarshalJSON with normal value",
			jsonStr:     "123",
			expected:    Some[T](123),
			expectedErr: false,
		},
		{
			name:        "Test UnmarshalJSON with invalid value",
			jsonStr:     "-",
			expected:    None[T](),
			expectedErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var actual Option[T]
			err := actual.UnmarshalJSON([]byte(tc.jsonStr))
			if !tc.expectedErr {
				if err != nil {
					t.Errorf("UnmarshalJSON() error = %v", err)
				}
				if !cmp.Equal(actual, tc.expected) {
					t.Errorf("UnmarshalJSON() = %v, want %v", actual, tc.expected)
				}
			} else if err == nil {
				t.Error("UnmarshalJSON() expected error, but there is no error")
			}
		})
	}
}

func TestOption_MarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		obj      Option[T]
		expected string
	}{
		{
			name:     "Test MarshalJSON with None value",
			obj:      None[T](),
			expected: "null",
		},
		{
			name:     "Test MarshalJSON with Some value",
			obj:      Some[T](123),
			expected: "123",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.obj.MarshalJSON()
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
			}
			if string(actual) != tc.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(actual), tc.expected)
			}
		})
	}
}
