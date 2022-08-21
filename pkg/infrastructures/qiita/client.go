package qiita

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type Client struct {
	token   string
	httpAPI HTTPAPI
}

func New(token string) *Client {
	return &Client{
		token:   token,
		httpAPI: new(http.Client),
	}
}

func (cl *Client) GetItems(page, perPage int, query string) (Items, error) {
	req, err := http.NewRequest(http.MethodGet, "https://qiita.com/api/v2/items", nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("per_page", strconv.Itoa(perPage))
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	resp, err := cl.httpAPI.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return nil, errors.New(string(b))
	}

	var items Items
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, errors.WithStack(err)
	}

	return items, nil
}
