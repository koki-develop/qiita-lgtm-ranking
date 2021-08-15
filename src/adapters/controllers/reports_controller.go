package controllers

import (
	"fmt"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/presenters"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
)

// IReportsController .
type IReportsController interface {
	UpdateWeekly(t time.Time, reportID string) error
	UpdateWeeklyByTag(t time.Time, reportID, tag string) error
}

// ReportsController .
type ReportsController struct {
	itemsRepository   ItemsRepository
	reportsRepository gateways.IReportsRepository
	reportsPresenter  presenters.IReportsPresenter
}

// UpdateWeekly .
func (c *ReportsController) UpdateWeekly(t time.Time, reportID string) error {
	from := t.AddDate(0, 0, -7)
	query := fmt.Sprintf("created:>=%s stocks:>=10", from.Format("2006-01-02"))

	items, err := c.itemsRepository.FindAll(query)
	if err != nil {
		return err
	}

	body, err := c.reportsPresenter.Weekly(from, &items)
	if err != nil {
		return err
	}

	if err := c.reportsRepository.Update(reportID, "Qiita 週間 LGTM 数ランキング【自動更新】", body, entities.Tags{}); err != nil {
		return err
	}

	return nil
}

// UpdateWeeklyByTag .
func (c *ReportsController) UpdateWeeklyByTag(t time.Time, reportID, tag string) error {
	from := t.AddDate(0, 0, -7)
	query := fmt.Sprintf("created:>=%s stocks:>=2 tag:%s", from.Format("2006-01-02"), tag)

	items, err := c.itemsRepository.FindAll(query)
	if err != nil {
		return err
	}

	body, err := c.reportsPresenter.WeeklyByTag(from, &items, tag)
	if err != nil {
		return err
	}

	if err := c.reportsRepository.Update(reportID, fmt.Sprintf("【%s】Qiita 週間 LGTM 数ランキング【自動更新】", tag), body, entities.Tags{{Name: tag}}); err != nil {
		return err
	}

	return nil
}
