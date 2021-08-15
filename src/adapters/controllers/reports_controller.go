package controllers

import (
	"fmt"
	"time"
)

type IReportsController interface {
	UpdateWeekly(t time.Time, reportID string) error
	UpdateWeeklyByTag(t time.Time, reportID, tag string) error
}

type ReportsController struct {
	itemsRepository   ItemsRepository
	reportsRepository ReportsRepository
}

func (c *ReportsController) UpdateWeekly(t time.Time, reportID string) error {
	from := t.AddDate(0, 0, -7)
	query := fmt.Sprintf("created:>=%s stocks:>=10", from.Format("2006-01-02"))

	items, err := c.itemsRepository.FindAll(query)
	if err != nil {
		return err
	}

	if err := c.reportsRepository.UpdateWeekly(from, reportID, items); err != nil {
		return err
	}

	return nil
}

func (c *ReportsController) UpdateWeeklyByTag(t time.Time, reportID, tag string) error {
	from := t.AddDate(0, 0, -7)
	query := fmt.Sprintf("created:>=%s stocks:>=2 tag:%s", from.Format("2006-01-02"), tag)

	items, err := c.itemsRepository.FindAll(query)
	if err != nil {
		return err
	}

	if err := c.reportsRepository.UpdateWeeklyByTag(from, reportID, items, tag); err != nil {
		return err
	}

	return nil
}
