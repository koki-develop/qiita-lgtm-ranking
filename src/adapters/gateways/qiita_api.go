package gateways

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
)

// QiitaAPI .
type QiitaAPI interface {
	GetItems(page, perPage int, query string) (entities.Items, error)
	UpdateItem(id, title, body string, tags entities.Tags) error
}
