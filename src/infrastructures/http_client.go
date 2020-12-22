package infrastructures

import (
	"bytes"
	"io"
	"net/http"
)

// HTTPAPI ...
type HTTPAPI interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPResponse ...
type HTTPResponse struct {
	StatusCode int
	Body       io.Reader
}

// IHTTPClient ...
type IHTTPClient interface {
	Do(req *http.Request) (*HTTPResponse, error)
}

// HTTPClient ...
type HTTPClient struct {
	httpAPI HTTPAPI
}

// NewHTTPClient ...
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{httpAPI: new(http.Client)}
}

// Do ...
func (c *HTTPClient) Do(req *http.Request) (*HTTPResponse, error) {
	resp, err := c.httpAPI.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return &HTTPResponse{
		Body:       buf,
		StatusCode: resp.StatusCode,
	}, nil
}
