package controllers

import (
	"fmt"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/gateways"
)

// IItemsController .
type IItemsController interface {
	UpdateWeeklyPerTag(t time.Time, reportID, tag string) error
}

// ItemsController .
type ItemsController struct {
	itemsRepository gateways.IItemsRepository
}

// UpdateWeeklyPerTag .
func (c *ItemsController) UpdateWeeklyPerTag(t time.Time, reportID, tag string) error {
	items, err := c.itemsRepository.GetAll(t.AddDate(0, 0, -7), tag)
	if err != nil {
		return err
	}

	// TODO: update item
	for _, item := range *items {
		fmt.Printf("%+v\n", item)
	}

	return nil
}
