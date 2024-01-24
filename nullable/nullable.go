package nullable

import (
	"encoding/json"
)

type Nullable[T any] struct {
	set   bool
	value T
}

func (n Nullable[T]) Value() T {
	if !n.set {
		var a T
		return a
	}

	return n.value
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.set {
		return []byte("null"), nil
	}

	return json.Marshal(n.value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.set = false
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	n.value = v
	n.set = true
	return nil
}
