package infrastructures

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// QiitaAPI .
type QiitaAPI struct {
	accessToken string
	httpClient  IHTTPClient
	jsonDecoder IJSONDecoder
}

// NewQiitaAPI .
func NewQiitaAPI(accessToken string) *QiitaAPI {
	return &QiitaAPI{
		accessToken: accessToken,
		httpClient:  NewHTTPClient(),
		jsonDecoder: NewJSONDecoder(),
	}
}

// GetItems .
func (api *QiitaAPI) GetItems(page, perPage int, query string) (*domain.Items, error) {
	req, err := http.NewRequest("GET", "https://qiita.com/api/v2/items", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("per_page", strconv.Itoa(perPage))
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api.accessToken))

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		return nil, errors.New(buf.String())
	}

	items := &domain.Items{}
	if err := api.jsonDecoder.Decode(resp.Body, items); err != nil {
		return nil, err
	}

	return items, nil
}
