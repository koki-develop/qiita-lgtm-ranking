package controllers

import (
	"errors"
	"testing"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * ReportsController.UpdateWeekly()
 */

func TestReportsController_UpdateWeekly_ReturnNilWhenSucceeded(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=10").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("Weekly", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}).Return("BODY", nil)

	mrr := new(mockReportsRepository)
	mrr.On("Update", "REPORT_ID", "Qiita 週間 LGTM 数ランキング【自動更新】", "BODY", domain.Tags{}).Return(nil)

	c := &ReportsController{
		itemsRepository:   mir,
		reportsPresenter:  mrp,
		reportsRepository: mrr,
	}

	err := c.UpdateWeekly(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID")

	assert.Nil(t, err)
	mir.AssertNumberOfCalls(t, "GetAll", 1)
	mrp.AssertNumberOfCalls(t, "Weekly", 1)
	mrr.AssertNumberOfCalls(t, "Update", 1)
}

func TestReportsController_UpdateWeekly_ReturnErrorWhenGetAllFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=10").Return((*domain.Items)(nil), errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository: mir,
	}

	err := c.UpdateWeekly(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAll", 1)
}

func TestReportsController_UpdateWeekly_ReturnErrorWhenWeeklyFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=10").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("Weekly", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}).Return("", errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository:  mir,
		reportsPresenter: mrp,
	}

	err := c.UpdateWeekly(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAll", 1)
	mrp.AssertNumberOfCalls(t, "Weekly", 1)
}

func TestReportsController_UpdateWeekly_ReturnErrorWhenUpdateFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=10").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("Weekly", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}).Return("BODY", nil)

	mrr := new(mockReportsRepository)
	mrr.On("Update", "REPORT_ID", "Qiita 週間 LGTM 数ランキング【自動更新】", "BODY", domain.Tags{}).Return(errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository:   mir,
		reportsPresenter:  mrp,
		reportsRepository: mrr,
	}

	err := c.UpdateWeekly(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAll", 1)
	mrp.AssertNumberOfCalls(t, "Weekly", 1)
	mrr.AssertNumberOfCalls(t, "Update", 1)
}

/*
 * ReportsController.UpdateWeeklyByTag()
 */

func TestReportsController_UpdateWeeklyByTag_ReturnNilWhenSucceeded(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=2 tag:TAG").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("WeeklyByTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}, "TAG").Return("BODY", nil)

	mrr := new(mockReportsRepository)
	mrr.On("Update", "REPORT_ID", "【TAG】Qiita 週間 LGTM 数ランキング【自動更新】", "BODY", domain.Tags{{Name: "TAG"}}).Return(nil)

	c := &ReportsController{
		itemsRepository:   mir,
		reportsPresenter:  mrp,
		reportsRepository: mrr,
	}

	err := c.UpdateWeeklyByTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.Nil(t, err)
	mir.AssertNumberOfCalls(t, "GetAll", 1)
	mrp.AssertNumberOfCalls(t, "WeeklyByTag", 1)
	mrr.AssertNumberOfCalls(t, "Update", 1)
}

func TestReportsController_UpdateWeeklyByTag_ReturnErrorWhenGetAllFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=2 tag:TAG").Return((*domain.Items)(nil), errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository: mir,
	}

	err := c.UpdateWeeklyByTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAll", 1)
}

func TestReportsController_UpdateWeeklyByTag_ReturnErrorWhenPresenterFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=2 tag:TAG").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("WeeklyByTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}, "TAG").Return("", errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository:  mir,
		reportsPresenter: mrp,
	}

	err := c.UpdateWeeklyByTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAll", 1)
	mrp.AssertNumberOfCalls(t, "WeeklyByTag", 1)
}

func TestReportsController_UpdateWeeklyByTag_ReturnUpdateFailed(t *testing.T) {
	mir := new(mockItemsRepository)
	mir.On("GetAll", "created:>=2020-01-01 stocks:>=2 tag:TAG").Return(&domain.Items{{Title: "TITLE"}}, nil)

	mrp := new(mockReportsPresenter)
	mrp.On("WeeklyByTag", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), &domain.Items{{Title: "TITLE"}}, "TAG").Return("BODY", nil)

	mrr := new(mockReportsRepository)
	mrr.On("Update", "REPORT_ID", "【TAG】Qiita 週間 LGTM 数ランキング【自動更新】", "BODY", domain.Tags{{Name: "TAG"}}).Return(errors.New("SOMETHING_WRONG"))

	c := &ReportsController{
		itemsRepository:   mir,
		reportsPresenter:  mrp,
		reportsRepository: mrr,
	}

	err := c.UpdateWeeklyByTag(time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), "REPORT_ID", "TAG")

	assert.EqualError(t, err, "SOMETHING_WRONG")
	mir.AssertNumberOfCalls(t, "GetAll", 1)
	mrp.AssertNumberOfCalls(t, "WeeklyByTag", 1)
	mrr.AssertNumberOfCalls(t, "Update", 1)
}
