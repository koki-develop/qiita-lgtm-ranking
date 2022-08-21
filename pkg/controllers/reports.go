package controllers

import (
	"fmt"
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/qiita"
	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/report"
)

type ReportController struct {
	builder     *report.Builder
	qiitaClient *qiita.Client
}

type ReportControllerConfig struct {
	QiitaAccessToken string
}

func NewReportController(cfg *ReportControllerConfig) *ReportController {
	return &ReportController{
		builder:     report.NewBuilder(),
		qiitaClient: qiita.New(cfg.QiitaAccessToken),
	}
}

func (ctrl *ReportController) UpdateDaily(rptID string) error {
	now := time.Now()
	from := now.AddDate(0, 0, -1)
	query := fmt.Sprintf("created:>=%s", from.Format("2006-01-02"))

	var items qiita.Items
	for i := 0; i < 100; i++ {
		rslt, err := ctrl.qiitaClient.GetItems(&qiita.GetItemsOptions{Page: i + 1, PerPage: 100, Query: query})
		if err != nil {
			return err
		}
		items = append(items, rslt.FilterWithMinLiked(1)...)
		if len(rslt) < 100 {
			break
		}
	}

	for _, item := range items {
		cnt, err := ctrl.qiitaClient.GetStockersCount(item.ID)
		if err != nil {
			return err
		}
		item.StockersCount = cnt
	}

	rpt, err := ctrl.builder.Build(from, now, items)
	if err != nil {
		return err
	}

	if err := ctrl.qiitaClient.UpdateItem(rptID, &qiita.UpdateItemPayload{
		Title: "Qiita デイリー LGTM 数ランキング【自動更新】",
		Body:  rpt,
		Tags: qiita.Tags{
			{Name: "Qiita"},
			{Name: "lgtm"},
			{Name: "ランキング"},
		},
	}); err != nil {
		return err
	}

	return nil
}
