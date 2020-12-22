package controllers

import (
	"os"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/infrastructures"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/gateways"
)

// ItemsControllerFactory .
type ItemsControllerFactory struct {
	osGetenv func(string) string
}

// NewItemsControllerFactory .
func NewItemsControllerFactory() *ItemsControllerFactory {
	return &ItemsControllerFactory{
		osGetenv: os.Getenv,
	}
}

// Create .
func (f *ItemsControllerFactory) Create() IItemsController {
	return &ItemsController{
		itemsRepository: gateways.NewItemsRepository(
			infrastructures.NewQiitaAPI(f.osGetenv("QIITA_ACCESS_TOKEN")),
		),
	}
}
