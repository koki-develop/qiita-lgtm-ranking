package gateways

import (
	"errors"
	"testing"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
	"github.com/stretchr/testify/assert"
)

/*
 * NewItemsRepository()
 */

func Test_NewItemsRepository(t *testing.T) {
	mapi := new(mockQiitaAPI)

	r := NewItemsRepository(mapi)

	assert.NotNil(t, r)
	assert.Equal(t, mapi, r.qiitaAPI)
}

/*
 * ItemsRepository.GetAll()
 */

func TestItemsRepository_GetAll_ReturnItemsWhenSucceeded(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("GetItems", 1, 100, "created:>=2020-01-01 stocks:>=1").Return(&domain.Items{{Title: "TITLE_1"}}, nil)
	mapi.On("GetItems", 2, 100, "created:>=2020-01-01 stocks:>=1").Return(&domain.Items{{Title: "TITLE_2"}}, nil)
	mapi.On("GetItems", 3, 100, "created:>=2020-01-01 stocks:>=1").Return(&domain.Items{}, nil)

	r := &ItemsRepository{qiitaAPI: mapi}

	items, err := r.GetAll(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, &domain.Items{
		{Title: "TITLE_1"},
		{Title: "TITLE_2"},
	}, items)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "GetItems", 3)
}

func TestItemsRepository_GetAll_ReturnErrorWhenGetItemsFailed(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("GetItems", 1, 100, "created:>=2020-01-01 stocks:>=1").Return((*domain.Items)(nil), errors.New("SOMETHING_WRONG"))

	r := &ItemsRepository{qiitaAPI: mapi}

	items, err := r.GetAll(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

	assert.Nil(t, items)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "GetItems", 1)
}

/*
 * ItemsRepository.GetAllWithTag()
 */

func TestItemsRepository_GetAllWithTag_ReturnItemsWhenSucceeded(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("GetItems", 1, 100, "created:>=2020-01-01 stocks:>=1 tag:TAG").Return(&domain.Items{{Title: "TITLE_1"}}, nil)
	mapi.On("GetItems", 2, 100, "created:>=2020-01-01 stocks:>=1 tag:TAG").Return(&domain.Items{{Title: "TITLE_2"}}, nil)
	mapi.On("GetItems", 3, 100, "created:>=2020-01-01 stocks:>=1 tag:TAG").Return(&domain.Items{}, nil)

	r := &ItemsRepository{qiitaAPI: mapi}

	items, err := r.GetAllWithTag(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "TAG")

	assert.Equal(t, &domain.Items{
		{Title: "TITLE_1"},
		{Title: "TITLE_2"},
	}, items)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "GetItems", 3)
}

func TestItemsRepository_GetAllWithTag_ReturnErrorWhenGetItemsFailed(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("GetItems", 1, 100, "created:>=2020-01-01 stocks:>=1 tag:TAG").Return((*domain.Items)(nil), errors.New("SOMETHING_WRONG"))

	r := &ItemsRepository{qiitaAPI: mapi}

	items, err := r.GetAllWithTag(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "TAG")

	assert.Nil(t, items)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "GetItems", 1)
}
