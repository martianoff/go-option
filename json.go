package option

import (
	"encoding/json"
	"strings"
)

func (opt *Option[T]) UnmarshalJSON(data []byte) error {
	if strings.ToLower(string(data)) == "null" {
		*opt = Option[T]{None[T]()}
		return nil
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*opt = Option[T]{Some[T](v)}
	return nil
}

func (opt Option[T]) MarshalJSON() ([]byte, error) {
	if opt.Empty() {
		return json.Marshal(nil)
	}
	return json.Marshal(opt.Get())
}
