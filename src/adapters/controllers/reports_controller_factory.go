package controllers

import (
	"os"

	itemsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/items"
	rptsrepo "github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways/reports"
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
	rptb := infrastructures.NewReportBuilder()

	return &ReportsController{
		itemsRepository:   itemsrepo.New(&itemsrepo.Config{QiitaAPI: qapi}),
		reportsRepository: rptsrepo.New(&rptsrepo.Config{QiitaAPI: qapi, ReportBuilder: rptb}),
	}
}
