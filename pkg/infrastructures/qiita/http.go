package qiita

import "net/http"

type HTTPAPI interface {
	Do(req *http.Request) (*http.Response, error)
}
