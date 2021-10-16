package infrastructures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
	"github.com/pkg/errors"
)

type QiitaClient struct {
	accessToken string
	httpAPI     HTTPAPI
}

func NewQiitaClient(accessToken string) *QiitaClient {
	return &QiitaClient{
		accessToken: accessToken,
		httpAPI:     http.DefaultClient,
	}
}

func (c *QiitaClient) GetItems(page, perPage int, query string) (entities.Items, error) {
	req, err := http.NewRequest("GET", "https://qiita.com/api/v2/items", nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(page))
	q.Add("per_page", strconv.Itoa(perPage))
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	resp, err := c.httpAPI.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return nil, errors.New(string(b))
	}

	fmt.Printf("Rate Limit Remaining: %s\n", resp.Header.Get("rate-remaining"))

	var items entities.Items
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, errors.WithStack(err)
	}

	return items, nil
}

func (c *QiitaClient) UpdateItem(id, title, body string, tags entities.Tags) error {
	p, err := json.Marshal(map[string]interface{}{
		"title": title,
		"body":  body,
		"tags":  tags,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://qiita.com/api/v2/items/%s", id), bytes.NewReader(p))
	if err != nil {
		return errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	resp, err := c.httpAPI.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.WithStack(err)
		}
		return errors.New(string(b))
	}

	return nil
}
