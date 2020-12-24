package controllers

import (
	"errors"
	"testing"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * ReportsController.UpdateWeeklyPerTag()
 */

func TestReportsController_UpdateWeeklyPerTag_ReturnNilWhenSucceeded(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAllWithTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "TAG").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("WeeklyPerTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}, "TAG").Return("BODY", nil)

	mrr := new(mockReportsRepository)
	mrr.On("Update", "REPORT_ID", "【TAG】Qiita 週間LGTM数ランキング【自動更新】", "BODY", "TAG").Return(nil)

	c := &ReportsController{
		itemsRepository:   mir,
		reportsPresenter:  mrp,
		reportsRepository: mrr,
	}

	err := c.UpdateWeeklyPerTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.Nil(t, err)
	mir.AssertNumberOfCalls(t, "GetAllWithTag", 1)
	mrp.AssertNumberOfCalls(t, "WeeklyPerTag", 1)
	mrr.AssertNumberOfCalls(t, "Update", 1)
}

func TestReportsController_UpdateWeeklyPerTag_ReturnErrorWhenGetAllWithTagFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAllWithTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "TAG").Return((*domain.Items)(nil), errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository: mir,
	}

	err := c.UpdateWeeklyPerTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAllWithTag", 1)
}

func TestReportsController_UpdateWeeklyPerTag_ReturnErrorWhenPresenterFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAllWithTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "TAG").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("WeeklyPerTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}, "TAG").Return("", errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository:  mir,
		reportsPresenter: mrp,
	}

	err := c.UpdateWeeklyPerTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAllWithTag", 1)
	mrp.AssertNumberOfCalls(t, "WeeklyPerTag", 1)
}

func TestReportsController_UpdateWeeklyPerTag_ReturnUpdateFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAllWithTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "TAG").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("WeeklyPerTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}, "TAG").Return("BODY", nil)

	mrr := new(mockReportsRepository)
	mrr.On("Update", "REPORT_ID", "【TAG】Qiita 週間LGTM数ランキング【自動更新】", "BODY", "TAG").Return(errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository:   mir,
		reportsPresenter:  mrp,
		reportsRepository: mrr,
	}

	err := c.UpdateWeeklyPerTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAllWithTag", 1)
	mrp.AssertNumberOfCalls(t, "WeeklyPerTag", 1)
	mrr.AssertNumberOfCalls(t, "Update", 1)
}
