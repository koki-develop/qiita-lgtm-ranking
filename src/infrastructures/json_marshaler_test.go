package infrastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testJSONMarshalerStruct struct {
	String  string            `json:"string"`
	Int     int               `json:"int"`
	Float64 float64           `json:"float64"`
	Bool    bool              `json:"bool"`
	Map     map[string]string `json:"map"`
	Struct  struct {
		Key string `json:"key"`
	} `json:"struct"`
	Array []struct {
		Key string `json:"key"`
	} `json:"array"`
}

/*
 * NewJSONMarshaler()
 */

func Test_NewJSONMarshaler(t *testing.T) {
	m := NewJSONMarshaler()

	assert.NotNil(t, m)
}

/*
 * JSONMarshaler.Marshal()
 */

func TestJSONMarshaler_Marshal_ReturnJSONAsBytes(t *testing.T) {
	m := new(JSONMarshaler)

	bs, err := m.Marshal(&testJSONMarshalerStruct{
		String:  "a",
		Int:     1,
		Float64: 1.1,
		Bool:    true,
		Map:     map[string]string{"key": "value"},
		Struct: struct {
			Key string `json:"key"`
		}{Key: "value"},
		Array: []struct {
			Key string `json:"key"`
		}{
			{Key: "value"},
		},
	})

	assert.Equal(t, []byte(`{"string":"a","int":1,"float64":1.1,"bool":true,"map":{"key":"value"},"struct":{"key":"value"},"array":[{"key":"value"}]}`), bs)
	assert.Nil(t, err)
}
