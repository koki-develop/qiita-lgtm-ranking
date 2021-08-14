package gateways

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
	"github.com/stretchr/testify/mock"
)

type mockQiitaAPI struct {
	mock.Mock
}

func (m *mockQiitaAPI) GetItems(page, perPage int, query string) (*entities.Items, error) {
	args := m.Called(page, perPage, query)
	return args.Get(0).(*entities.Items), args.Error(1)
}

func (m *mockQiitaAPI) UpdateItem(id, title, body string, tags entities.Tags) error {
	args := m.Called(id, title, body, tags)
	return args.Error(0)
}
