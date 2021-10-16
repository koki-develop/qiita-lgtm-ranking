package gateways

import (
	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
)

// QiitaAPI .
type QiitaAPI interface {
	GetItems(page, perPage int, query string) (entities.Items, error)
	GetStockersOfItem(id string, page, perPage int) (entities.Users, error)
	UpdateItem(id, title, body string, tags entities.Tags) error
}
