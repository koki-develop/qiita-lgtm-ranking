package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/qiita"
	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/report"
	"github.com/pkg/errors"
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

	files, err := filepath.Glob("./events/updateDailyByTag/*.prod.json")
	if err != nil {
		return errors.WithStack(err)
	}

	var tags report.Tags
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return errors.WithStack(err)
		}
		defer f.Close()
		var tag report.Tag
		if err := json.NewDecoder(f).Decode(&tag); err != nil {
			return errors.WithStack(err)
		}
		tags = append(tags, &tag)
	}

	rpt, err := ctrl.builder.Build(&report.BuildOptions{
		Tags: tags,
		Conditions: report.Conditions{
			{Key: "期間", Value: fmt.Sprintf("%s ~ %s", from.Format("2006-01-02"), now.Format("2006-01-02"))},
		},
		Items: items,
	})
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

func (ctrl *ReportController) UpdateDailyByTag(rptID, tag string) error {
	now := time.Now()
	from := now.AddDate(0, 0, -1)
	query := fmt.Sprintf("created:>=%s tag:%s", from.Format("2006-01-02"), tag)

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

	files, err := filepath.Glob("./events/updateDailyByTag/*.prod.json")
	if err != nil {
		return errors.WithStack(err)
	}

	var tags report.Tags
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return errors.WithStack(err)
		}
		defer f.Close()
		var tag report.Tag
		if err := json.NewDecoder(f).Decode(&tag); err != nil {
			return errors.WithStack(err)
		}
		tags = append(tags, &tag)
	}

	rpt, err := ctrl.builder.Build(&report.BuildOptions{
		Tags: tags,
		Conditions: report.Conditions{
			{Key: "期間", Value: fmt.Sprintf("%s ~ %s", from.Format("2006-01-02"), now.Format("2006-01-02"))},
		},
		Items: items,
	})
	if err != nil {
		return err
	}

	if err := ctrl.qiitaClient.UpdateItem(rptID, &qiita.UpdateItemPayload{
		Title: fmt.Sprintf("【%s】Qiita デイリー LGTM 数ランキング【自動更新】", tag),
		Body:  rpt,
		Tags: qiita.Tags{
			{Name: "Qiita"},
			{Name: "lgtm"},
			{Name: "ランキング"},
			{Name: tag},
		},
	}); err != nil {
		return err
	}

	return nil
}
