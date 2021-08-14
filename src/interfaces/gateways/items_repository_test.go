package gateways

import (
	"errors"
	"testing"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
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
	mapi.On("GetItems", 1, 100, "QUERY").Return(&entities.Items{{Title: "TITLE_1"}}, nil)
	mapi.On("GetItems", 2, 100, "QUERY").Return(&entities.Items{{Title: "TITLE_2"}}, nil)
	mapi.On("GetItems", 3, 100, "QUERY").Return(&entities.Items{}, nil)

	r := &ItemsRepository{qiitaAPI: mapi}

	items, err := r.GetAll("QUERY")

	assert.Equal(t, &entities.Items{
		{Title: "TITLE_1"},
		{Title: "TITLE_2"},
	}, items)
	assert.Nil(t, err)
	mapi.AssertNumberOfCalls(t, "GetItems", 3)
}

func TestItemsRepository_GetAll_ReturnErrorWhenGetItemsFailed(t *testing.T) {
	mapi := new(mockQiitaAPI)
	mapi.On("GetItems", 1, 100, "QUERY").Return((*entities.Items)(nil), errors.New("SOMETHING_WRONG"))

	r := &ItemsRepository{qiitaAPI: mapi}

	items, err := r.GetAll("QUERY")

	assert.Nil(t, items)
	assert.EqualError(t, err, "SOMETHING_WRONG")
	mapi.AssertNumberOfCalls(t, "GetItems", 1)
}
