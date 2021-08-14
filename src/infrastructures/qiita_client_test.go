package infrastructures

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/infrastructures/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mocksForQiitaClient struct {
	httpAPI *mocks.HTTPAPI
}

func (ms *mocksForQiitaClient) AssertExpectations(t *testing.T) {
	ms.httpAPI.AssertExpectations(t)
}

func setupQiitaClientAndMocks() (*QiitaClient, *mocksForQiitaClient) {
	ms := &mocksForQiitaClient{
		httpAPI: &mocks.HTTPAPI{},
	}
	return &QiitaClient{
		accessToken: "ACCESS_TOKEN",
		httpAPI:     ms.httpAPI,
	}, ms
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

func TestQiitaClient_GetItems(t *testing.T) {
	t.Run("return items when succeeded", func(t *testing.T) {
		c, ms := setupQiitaClientAndMocks()
		dummyPage := 10
		dummyPerPage := 100
		dummyQuery := "QUERY"
		dummyItems := entities.Items{
			{
				Title:      "ITEM_1",
				LikesCount: 1,
				URL:        "URL_1",
				User:       entities.User{ID: "USER_1", ProfileImageURL: "PROFILE_IMAGE_URL_1"},
				Tags:       entities.Tags{{Name: "TAG_1"}},
				CreatedAt:  time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				Title:      "ITEM_2",
				LikesCount: 2,
				URL:        "URL_2",
				User:       entities.User{ID: "USER_2", ProfileImageURL: "PROFILE_IMAGE_URL_2"},
				Tags:       entities.Tags{{Name: "TAG_2"}},
				CreatedAt:  time.Date(1998, 2, 1, 0, 0, 0, 0, time.UTC),
			},
		}

		ms.httpAPI.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
			Body: ioutil.NopCloser(strings.NewReader(`[
        {
          "title": "ITEM_1",
          "likes_count": 1,
          "url": "URL_1",
          "user": {
            "id": "USER_1",
            "profile_image_url": "PROFILE_IMAGE_URL_1"
          },
          "tags": [
            {
              "name": "TAG_1"
            }
          ],
          "created_at": "1998-01-01T00:00:00.000Z"
        },
        {
          "title": "ITEM_2",
          "likes_count": 2,
          "url": "URL_2",
          "user": {
            "id": "USER_2",
            "profile_image_url": "PROFILE_IMAGE_URL_2"
          },
          "tags": [
            {
              "name": "TAG_2"
            }
          ],
          "created_at": "1998-02-01T00:00:00.000Z"
        }
      ]`)),
			StatusCode: http.StatusOK,
		}, nil).Run(func(args mock.Arguments) {
			req := args.Get(0).(*http.Request)
			assert.Equal(t, "https://qiita.com/api/v2/items?page=10&per_page=100&query=QUERY", req.URL.String())
			assert.Equal(t, http.MethodGet, req.Method)
			assert.Equal(t, http.Header{"Authorization": {"Bearer ACCESS_TOKEN"}}, req.Header)
		})

		items, err := c.GetItems(dummyPage, dummyPerPage, dummyQuery)
		assert.Equal(t, dummyItems, items)
		assert.NoError(t, err)
		ms.AssertExpectations(t)
	})
}

func TestQiitaClient_UpdateItem(t *testing.T) {
	t.Run("return nil when succeeded", func(t *testing.T) {
		c, ms := setupQiitaClientAndMocks()
		dummyItemID := "ITEM_ID"
		dummyTitle := "TITLE"
		dummyBody := "BODY"
		dummyTags := entities.Tags{{Name: "TAG"}}

		ms.httpAPI.On("Do", mock.AnythingOfType("*http.Request")).Return(&http.Response{
			Body:       ioutil.NopCloser(strings.NewReader("{}")),
			StatusCode: http.StatusOK,
		}, nil).Run(func(args mock.Arguments) {
			req := args.Get(0).(*http.Request)
			b, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, "https://qiita.com/api/v2/items/ITEM_ID", req.URL.String())
			assert.Equal(t, `{"body":"BODY","tags":[{"name":"TAG"}],"title":"TITLE"}`, string(b))
			assert.Equal(t, http.MethodPatch, req.Method)
			assert.Equal(t, http.Header{"Authorization": {"Bearer ACCESS_TOKEN"}, "Content-Type": {"application/json"}}, req.Header)
		})

		err := c.UpdateItem(dummyItemID, dummyTitle, dummyBody, dummyTags)
		assert.NoError(t, err)
		ms.AssertExpectations(t)
	})
}