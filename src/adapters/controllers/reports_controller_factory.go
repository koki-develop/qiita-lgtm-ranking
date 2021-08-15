package controllers

import (
	"os"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways"
	itemsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/items"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/presenters"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/infrastructures"
)

// ReportsControllerFactory .
type ReportsControllerFactory struct {
	osGetenv func(string) string
}

// NewReportsControllerFactory .
func NewReportsControllerFactory() *ReportsControllerFactory {
	return &ReportsControllerFactory{
		osGetenv: os.Getenv,
	}
}

// Create .
func (f *ReportsControllerFactory) Create() IReportsController {
	qapi := infrastructures.NewQiitaClient(f.osGetenv("QIITA_ACCESS_TOKEN"))

	return &ReportsController{
		itemsRepository:   itemsrepo.New(&itemsrepo.Config{QiitaAPI: qapi}),
		reportsRepository: gateways.NewReportsRepository(qapi),
		reportsPresenter:  presenters.NewReportsPresenter(),
	}
}
