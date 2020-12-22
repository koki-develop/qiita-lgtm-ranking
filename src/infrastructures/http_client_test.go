package infrastructures

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
 * NewHTTPClient()
 */

func Test_NewHTTPClient(t *testing.T) {
	c := NewHTTPClient()

	assert.NotNil(t, c)
	assert.NotNil(t, c.httpAPI)
	assert.IsType(t, &http.Client{}, c.httpAPI)
}

/*
 * HTTPClient.Do()
 */

func TestHTTPClient_Do_ReturnResponseWhenSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	mh := new(mockHTTPAPI)
	mh.On("Do", req).Return(&http.Response{
		Body:       ioutil.NopCloser(strings.NewReader("BODY")),
		StatusCode: 200,
	}, nil)

	c := &HTTPClient{httpAPI: mh}
	resp, err := c.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	assert.Equal(t, "BODY", buf.String())
	assert.Equal(t, 200, resp.StatusCode)
	assert.Nil(t, err)
	mh.AssertNumberOfCalls(t, "Do", 1)
}

func TestHTTPClient_Do_ReturnErrorWhenRequestFailed(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	mh := new(mockHTTPAPI)
	mh.On("Do", req).Return((*http.Response)(nil), errors.New("SOMETHING_WRONG"))

	c := &HTTPClient{httpAPI: mh}
	resp, err := c.Do(req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mh.AssertNumberOfCalls(t, "Do", 1)
}
