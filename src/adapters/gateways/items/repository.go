package items

import (
	"github.com/koki-develop/qiita-lgtm-ranking/src/adapters/gateways"
	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
	"github.com/pkg/errors"
)

type Repository struct {
	config *Config
}

type Config struct {
	QiitaAPI gateways.QiitaAPI
}

func New(cfg *Config) *Repository {
	return &Repository{config: cfg}
}

func (repo *Repository) FindAll(query string) (entities.Items, error) {
	var items entities.Items

	for i := 1; i <= 100; i++ {
		resp, err := repo.config.QiitaAPI.GetItems(i, 100, query)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(resp) == 0 {
			break
		}

		items = append(items, resp.FilterOnlyHasLGTM()...)
	}

	for _, item := range items {
		for i := 1; i <= 100; i++ {
			stks, err := repo.config.QiitaAPI.GetStockersOfItem(item.ID, i, 100)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			item.Stockers = stks
		}
	}

	return items, nil
}
