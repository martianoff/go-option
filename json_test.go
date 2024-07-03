package option

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"testing"
)

type S struct {
	Foo Option[string] `json:"foo"`
	Bar Option[string] `json:"bar"`
}

func TestOption_UnmarshalJSON(t *testing.T) {
	type T = int
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
				// validate with standard method
				assert.Equal(t, actual, tc.expected)
				// validate with custom equality
				if !cmp.Equal(actual, tc.expected) {
					t.Errorf("UnmarshalJSON() = %v, want %v", actual, tc.expected)
				}
			} else if err == nil {
				t.Error("UnmarshalJSON() expected error, but there is no error")
			}
		})
	}
}

func TestOption_UnmarshalJSONStr(t *testing.T) {
	type T = string
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
			jsonStr:     "\"option\"",
			expected:    Some[T]("option"),
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
				// validate with standard method
				assert.Equal(t, actual, tc.expected)
				// validate with custom equality
				if !cmp.Equal(actual, tc.expected) {
					t.Errorf("UnmarshalJSON() = %v, want %v", actual, tc.expected)
				}
			} else if err == nil {
				t.Error("UnmarshalJSON() expected error, but there is no error")
			}
		})
	}
}

func TestOption_UnmarshalJSON_fromStruct(t *testing.T) {
	cases := []struct {
		name        string
		jsonStr     string
		expected    S
		expectedErr bool
	}{
		{
			name:    "Test struct with empty values",
			jsonStr: "{\"foo\":null,\"bar\":null}",
			expected: S{
				Foo: None[string](),
				Bar: None[string](),
			},
			expectedErr: false,
		},
		{
			name:    "Test struct with filled values",
			jsonStr: "{\"foo\":\"bar\",\"bar\":\"baz\"}",
			expected: S{
				Foo: Some("bar"),
				Bar: Some("baz"),
			},
			expectedErr: false,
		},
		{
			name:    "Test struct with invalid values",
			jsonStr: "{\"foo\":1,\"bar\":\"baz\"}",
			expected: S{
				Foo: None[string](),
				Bar: None[string](),
			},
			expectedErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := S{}
			t.Logf("actual = %v", actual)
			err := json.Unmarshal([]byte(tc.jsonStr), &actual)
			if !tc.expectedErr {
				if err != nil {
					t.Errorf("UnmarshalJSON() error = %v", err)
				}
				// validate with standard method
				assert.Equal(t, actual, tc.expected)
				// validate with custom equality
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
	type T = int
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
