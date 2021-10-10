package controllers

import "github.com/koki-develop/qiita-lgtm-ranking/src/entities"

type ItemsRepository interface {
	FindAll(query string) (entities.Items, error)
}
