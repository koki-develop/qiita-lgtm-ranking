package items

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/gateways"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
	"github.com/pkg/errors"
)

type Repository struct {
	config *Config
}

type Config struct {
	QiitaAPI gateways.IQiitaAPI
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

	return items, nil
}
