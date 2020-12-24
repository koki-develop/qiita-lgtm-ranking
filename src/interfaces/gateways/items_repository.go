package gateways

import (
	"fmt"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IItemsRepository .
type IItemsRepository interface {
	GetAll(from time.Time) (*domain.Items, error)
	GetAllWithTag(from time.Time, tag string) (*domain.Items, error)
}

// ItemsRepository .
type ItemsRepository struct {
	qiitaAPI IQiitaAPI
}

// NewItemsRepository .
func NewItemsRepository(api IQiitaAPI) *ItemsRepository {
	return &ItemsRepository{qiitaAPI: api}
}

// GetAll .
func (r *ItemsRepository) GetAll(from time.Time) (*domain.Items, error) {
	query := fmt.Sprintf("created:>%s stocks:>=1", from.Format("2006-01-02"))
	items, err := r.getAll(query)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// GetAllWithTag .
func (r *ItemsRepository) GetAllWithTag(from time.Time, tag string) (*domain.Items, error) {
	query := fmt.Sprintf("created:>%s stocks:>=1 tag:%s", from.Format("2006-01-02"), tag)
	items, err := r.getAll(query)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r *ItemsRepository) getAll(query string) (*domain.Items, error) {
	items := &domain.Items{}

	for i := 1; i <= 100; i++ {
		resp, err := r.qiitaAPI.GetItems(i, 100, query)
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
