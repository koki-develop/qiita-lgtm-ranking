package gateways

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IQiitaAPI .
type IQiitaAPI interface {
	GetItems(page, perPage int, query string) (*domain.Items, error)
	UpdateItem(id, title, body string, tags []string) error
}
