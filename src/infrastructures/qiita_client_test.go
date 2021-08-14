package infrastructures

import (
	"testing"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/infrastructures/mocks"
	"github.com/stretchr/testify/assert"
)

type mocksForQiitaClient struct {
	HTTPAPI *mocks.HTTPAPI
}

func (ms *mocksForQiitaClient) AssertExpectations(t *testing.T) {
	ms.HTTPAPI.AssertExpectations(t)
}

func Test_NewQiitaClient(t *testing.T) {
	t.Run("return QiitaClient", func(t *testing.T) {
		dummyAccessToken := "ACCESS_TOKEN"
		c := NewQiitaClient(dummyAccessToken)

		assert.NotNil(t, c)
		assert.IsType(t, &QiitaClient{}, c)
		assert.Equal(t, dummyAccessToken, c.accessToken)
		assert.NotNil(t, c.httpAPI)
	})
}
