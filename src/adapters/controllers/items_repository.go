package controllers

import "github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"

type ItemsRepository interface {
	FindAll(query string) (entities.Items, error)
}
