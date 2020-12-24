package infrastructures

import (
	"io"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type mockHTTPAPI struct {
	mock.Mock
}

func (m *mockHTTPAPI) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

type mockHTTPClient struct {
	mock.Mock
}

func (m *mockHTTPClient) Do(req *http.Request) (*HTTPResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*HTTPResponse), args.Error(1)
}

type mockJSONDecoder struct {
	mock.Mock
}

func (m *mockJSONDecoder) Decode(r io.Reader, obj interface{}) error {
	args := m.Called(r, obj)
	return args.Error(0)
}

type mockJSONMarshaler struct {
	mock.Mock
}

func (m *mockJSONMarshaler) Marshal(obj interface{}) ([]byte, error) {
	args := m.Called(obj)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *mockJSONMarshaler) MarshalPretty(obj interface{}) ([]byte, error) {
	args := m.Called(obj)
	return args.Get(0).([]byte), args.Error(1)
}
