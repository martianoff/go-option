package option

import (
	"encoding/json"
	"strings"
)

// UnmarshalJSON - custom JSON unmarshalling for Option
func (opt *Option[T]) UnmarshalJSON(data []byte) error {
	if strings.ToLower(string(data)) == "null" {
		*opt = None[T]()
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*opt = Some[T](v)
	return nil
}

// MarshalJSON - custom JSON marshalling for Option
func (opt Option[T]) MarshalJSON() ([]byte, error) {
	if opt.Empty() {
		return json.Marshal(nil)
	}
	return json.Marshal(opt.Get())
}
