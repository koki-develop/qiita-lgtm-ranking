package infrastructures

import (
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
 * NewQiitaAPI
 */

func Test_NewQiitaAPI(t *testing.T) {
	api := NewQiitaAPI("TOKEN")

	assert.Equal(t, "TOKEN", api.accessToken)
	assert.NotNil(t, api.httpClient)
	assert.IsType(t, &HTTPClient{}, api.httpClient)
	assert.NotNil(t, api.jsonDecoder)
	assert.IsType(t, &JSONDecoder{}, api.jsonDecoder)
	assert.NotNil(t, api.jsonMarshaler)
	assert.IsType(t, &JSONMarshaler{}, api.jsonMarshaler)
}

/*
 * QiitaAPI.GetItems()
 */

func TestQiitaAPI_GetItems_ReturnItemsWhenSucceeded(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/items?page=1&per_page=1&query=QUERY", nil)
	req.Header.Set("Authorization", "Bearer TOKEN")

	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return(&HTTPResponse{
		Body:       strings.NewReader("RESPONSE"),
		StatusCode: 200,
	}, nil)

	mjd := new(mockJSONDecoder)
	mjd.On("Decode", strings.NewReader("RESPONSE"), &domain.Items{}).Return(nil).Run(func(args mock.Arguments) {
		items := args.Get(1).(*domain.Items)
		*items = domain.Items{{Title: "TITLE"}}
	})

	api := &QiitaAPI{
		accessToken: "TOKEN",
		httpClient:  mhc,
		jsonDecoder: mjd,
	}

	items, err := api.GetItems(1, 1, "QUERY")

	assert.Equal(t, &domain.Items{{Title: "TITLE"}}, items)
	assert.Nil(t, err)
	mhc.AssertNumberOfCalls(t, "Do", 1)
	mjd.AssertNumberOfCalls(t, "Decode", 1)
}

func TestQiitaAPI_GetItems_ReturnErrorWhenHTTPClientDoFailed(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/items?page=1&per_page=1&query=QUERY", nil)
	req.Header.Set("Authorization", "Bearer TOKEN")

	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return((*HTTPResponse)(nil), errors.New("SOMETHING_WRONG"))

	api := &QiitaAPI{
		accessToken: "TOKEN",
		httpClient:  mhc,
	}

	items, err := api.GetItems(1, 1, "QUERY")

	assert.Nil(t, items)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mhc.AssertNumberOfCalls(t, "Do", 1)
}

func TestQiitaAPI_GetItems_ReturnErrorReceivedFailedHTTPStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/items?page=1&per_page=1&query=QUERY", nil)
	req.Header.Set("Authorization", "Bearer TOKEN")

	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return(&HTTPResponse{
		Body:       strings.NewReader("RESPONSE"),
		StatusCode: 500,
	}, nil)

	api := &QiitaAPI{
		accessToken: "TOKEN",
		httpClient:  mhc,
	}

	items, err := api.GetItems(1, 1, "QUERY")

	assert.Nil(t, items)
	assert.EqualError(t, err, "RESPONSE")
	mhc.AssertNumberOfCalls(t, "Do", 1)
}

func TestQiitaAPI_GetItems_ReturnErrorWhenJSONDecodeFailed(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/items?page=1&per_page=1&query=QUERY", nil)
	req.Header.Set("Authorization", "Bearer TOKEN")

	mhc := new(mockHTTPClient)
	mhc.On("Do", req).Return(&HTTPResponse{
		Body:       strings.NewReader("RESPONSE"),
		StatusCode: 200,
	}, nil)

	mjd := new(mockJSONDecoder)
	mjd.On("Decode", strings.NewReader("RESPONSE"), &domain.Items{}).Return(errors.New("SOMETHING_WRONG"))

	api := &QiitaAPI{
		accessToken: "TOKEN",
		httpClient:  mhc,
		jsonDecoder: mjd,
	}

	items, err := api.GetItems(1, 1, "QUERY")

	assert.Nil(t, items)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mhc.AssertNumberOfCalls(t, "Do", 1)
	mjd.AssertNumberOfCalls(t, "Decode", 1)
}

/*
 * QiitaAPI.UpdateItem()
 */

func TestQiitaAPI_UpdateItem_ReturnNilWhenSucceeded(t *testing.T) {
	mjm := new(mockJSONMarshaler)
	mjm.On("Marshal", map[string]interface{}{
		"title": "TITLE",
		"body":  "BODY",
		"tags":  domain.Tags{{Name: "TAG"}},
	}).Return([]byte("BODY"), nil)

	mhc := new(mockHTTPClient)
	mhc.On("Do", mock.Anything).Return(&HTTPResponse{
		Body:       strings.NewReader("RESPONSE"),
		StatusCode: 200,
	}, nil)

	api := &QiitaAPI{
		accessToken:   "TOKEN",
		httpClient:    mhc,
		jsonMarshaler: mjm,
	}

	err := api.UpdateItem("ID", "TITLE", "BODY", domain.Tags{{Name: "TAG"}})

	assert.Nil(t, err)
	mjm.AssertNumberOfCalls(t, "Marshal", 1)
	mhc.AssertNumberOfCalls(t, "Do", 1)
}

func TestQiitaAPI_UpdateItem_ReturnErrorWhenJSONMarshalFailed(t *testing.T) {
	mjm := new(mockJSONMarshaler)
	mjm.On("Marshal", map[string]interface{}{
		"title": "TITLE",
		"body":  "BODY",
		"tags":  domain.Tags{{Name: "TAG"}},
	}).Return(([]byte)(nil), errors.New("SOMETHING_WRONG"))

	api := &QiitaAPI{
		accessToken:   "TOKEN",
		jsonMarshaler: mjm,
	}

	err := api.UpdateItem("ID", "TITLE", "BODY", domain.Tags{{Name: "TAG"}})

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mjm.AssertNumberOfCalls(t, "Marshal", 1)
}

func TestQiitaAPI_UpdateItem_ReturnErrorWhenHTTPClientDoFailed(t *testing.T) {
	mjm := new(mockJSONMarshaler)
	mjm.On("Marshal", map[string]interface{}{
		"title": "TITLE",
		"body":  "BODY",
		"tags":  domain.Tags{{Name: "TAG"}},
	}).Return([]byte("BODY"), nil)

	mhc := new(mockHTTPClient)
	mhc.On("Do", mock.Anything).Return((*HTTPResponse)(nil), errors.New("SOMETHING_WRONG"))

	api := &QiitaAPI{
		accessToken:   "TOKEN",
		httpClient:    mhc,
		jsonMarshaler: mjm,
	}

	err := api.UpdateItem("ID", "TITLE", "BODY", domain.Tags{{Name: "TAG"}})

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mjm.AssertNumberOfCalls(t, "Marshal", 1)
	mhc.AssertNumberOfCalls(t, "Do", 1)
}

func TestQiitaAPI_UpdateItem_ReturnErrorWhenReceivedHTTPFailedStatus(t *testing.T) {
	mjm := new(mockJSONMarshaler)
	mjm.On("Marshal", map[string]interface{}{
		"title": "TITLE",
		"body":  "BODY",
		"tags":  domain.Tags{{Name: "TAG"}},
	}).Return([]byte("BODY"), nil)

	mhc := new(mockHTTPClient)
	mhc.On("Do", mock.Anything).Return(&HTTPResponse{
		Body:       strings.NewReader("RESPONSE"),
		StatusCode: 500,
	}, nil)

	api := &QiitaAPI{
		accessToken:   "TOKEN",
		httpClient:    mhc,
		jsonMarshaler: mjm,
	}

	err := api.UpdateItem("ID", "TITLE", "BODY", domain.Tags{{Name: "TAG"}})

	assert.EqualError(t, err, "RESPONSE")
	mjm.AssertNumberOfCalls(t, "Marshal", 1)
	mhc.AssertNumberOfCalls(t, "Do", 1)
}
