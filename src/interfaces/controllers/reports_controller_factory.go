package controllers

import (
	"os"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/infrastructures"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/gateways"
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
	qapi := infrastructures.NewQiitaAPI(f.osGetenv("QIITA_ACCESS_TOKEN"))

	return &ReportsController{
		itemsRepository:   gateways.NewItemsRepository(qapi),
		reportsRepository: gateways.NewReportsRepository(qapi),
	}
}
