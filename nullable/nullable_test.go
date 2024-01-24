package nullable

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	tests := []struct {
		input    Nullable[any]
		expected any
	}{
		{Nullable[any]{true, 42}, 42},
		{Nullable[any]{set: false}, nil},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.input.Value(), test.expected) {
			t.Errorf("for input %v, got %#v, expected %#v", test.input, test.input.Value(), test.expected)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected Nullable[any]
	}{
		{`"test"`, Nullable[any]{value: "test", set: true}},
		{`42`, Nullable[any]{value: 42.0, set: true}},
		{`true`, Nullable[any]{value: true, set: true}},
		{`null`, Nullable[any]{set: false}},
		{`{}`, Nullable[any]{set: true, value: map[string]interface{}{}}},
	}

	for _, test := range tests {
		var result Nullable[any]
		err := json.Unmarshal([]byte(test.input), &result)
		if err != nil {
			t.Errorf("Unexpected error for input %s: %v", test.input, err)
			continue
		}

		if !reflect.DeepEqual(test.expected, result) {
			t.Errorf("for input %s, got %#v, expected %#v", test.input, result, test.expected)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		input    Nullable[any]
		expected string
	}{
		{Nullable[any]{set: true, value: "test"}, `"test"`},
		{Nullable[any]{set: true, value: 42}, `42`},
		{Nullable[any]{set: true, value: true}, `true`},
		{Nullable[any]{set: false}, `null`},
	}

	for _, test := range tests {
		resultJSON, err := json.Marshal(test.input)
		if err != nil {
			t.Errorf("Unexpected error for input %#v: %v", test.input, err)
			continue
		}

		if string(resultJSON) != test.expected {
			t.Errorf("For input %#v, got %s, expected %s", test.input, string(resultJSON), test.expected)
		}
	}
}
