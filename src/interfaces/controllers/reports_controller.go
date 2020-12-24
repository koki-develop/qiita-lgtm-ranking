package controllers

import (
	"fmt"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/gateways"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/presenters"
)

// IReportsController .
type IReportsController interface {
	UpdateWeeklyPerTag(t time.Time, reportID, tag string) error
}

// ReportsController .
type ReportsController struct {
	itemsRepository   gateways.IItemsRepository
	reportsRepository gateways.IReportsRepository
	reportsPresenter  presenters.IReportsPresenter
}

// UpdateWeeklyPerTag .
func (c *ReportsController) UpdateWeeklyPerTag(t time.Time, reportID, tag string) error {
	from := t.AddDate(0, 0, -7)

	items, err := c.itemsRepository.GetAllWithTag(from, tag)
	if err != nil {
		return err
	}

	body, err := c.reportsPresenter.WeeklyPerTag(from, items, tag)
	if err != nil {
		return err
	}

	if err := c.reportsRepository.Update(reportID, fmt.Sprintf("【%s】Qiita 週間LGTM数ランキング【自動更新】", tag), body, tag); err != nil {
		return err
	}

	return nil
}
