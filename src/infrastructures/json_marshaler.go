package infrastructures

import (
	jsoniter "github.com/json-iterator/go"
)

// IJSONMarshaler ...
type IJSONMarshaler interface {
	Marshal(obj interface{}) ([]byte, error)
}

// JSONMarshaler ...
type JSONMarshaler struct{}

// NewJSONMarshaler ...
func NewJSONMarshaler() *JSONMarshaler {
	return new(JSONMarshaler)
}

// Marshal ...
func (m *JSONMarshaler) Marshal(obj interface{}) ([]byte, error) {
	return jsoniter.Marshal(obj)
}
