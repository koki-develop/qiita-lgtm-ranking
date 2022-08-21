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

type GetItemsOptions struct {
	Page    int
	PerPage int
	Query   string
}

func (cl *Client) GetItems(opts *GetItemsOptions) (Items, error) {
	req, err := http.NewRequest(http.MethodGet, "https://qiita.com/api/v2/items", nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))

	q := req.URL.Query()
	q.Add("page", strconv.Itoa(opts.Page))
	q.Add("per_page", strconv.Itoa(opts.PerPage))
	q.Add("query", opts.Query)
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

func (cl *Client) GetStockersCount(itemid string) (int, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://qiita.com/api/v2/items/%s/stockers", itemid), nil)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))

	resp, err := cl.httpAPI.Do(req)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	defer resp.Body.Close()

	cstr := resp.Header.Get("total-count")
	c, err := strconv.Atoi(cstr)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return c, nil
}
