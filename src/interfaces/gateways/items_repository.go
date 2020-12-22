package gateways

import (
	"fmt"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IItemsRepository .
type IItemsRepository interface {
	GetAll(from time.Time, tag string) (*domain.Items, error)
}

// ItemsRepository .
type ItemsRepository struct {
	qiitaAPI IQiitaAPI
}

// GetAll .
func (r *ItemsRepository) GetAll(from time.Time, tag string) (*domain.Items, error) {
	items := &domain.Items{}

	for i := 1; i <= 100; i++ {
		resp, err := r.qiitaAPI.GetItems(i, 100, fmt.Sprintf("created:>%s stocks:>1 tag:%s", from.Format("2006-01-02"), tag))
		if err != nil {
			return nil, err
		}
		if len(*resp) == 0 {
			break
		}

		*items = append(*items, *resp...)
	}

	return items, nil
}
