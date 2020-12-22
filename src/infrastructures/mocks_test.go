package infrastructures

import (
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
