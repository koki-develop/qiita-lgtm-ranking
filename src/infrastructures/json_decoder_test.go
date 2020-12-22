package infrastructures

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testJSONDecoderStruct struct {
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
 * NewJSONDecoder()
 */

func Test_NewJSONDecoder(t *testing.T) {
	d := NewJSONDecoder()

	assert.NotNil(t, d)
}

/*
 * JSONDecoder.Decode()
 */

func TestJSONDecoder_Decode_ReturnNilWhenValidJSON(t *testing.T) {
	d := new(JSONDecoder)

	s := new(testJSONDecoderStruct)
	j := `{"string":"a","int":1,"float64":1.1,"bool":true,"struct":{"key":"value"},"array":[{"key":"value"}],"map":{"key":"value"}}`

	err := d.Decode(strings.NewReader(j), s)
	assert.Nil(t, err)
	assert.Equal(t, &testJSONDecoderStruct{
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
	}, s)
}

func TestJSONDecoder_Decode_InvalidJSON(t *testing.T) {
	d := new(JSONDecoder)

	j := "INVALID_JSON"

	err := d.Decode(strings.NewReader(j), new(testJSONDecoderStruct))

	assert.NotNil(t, err)
}
