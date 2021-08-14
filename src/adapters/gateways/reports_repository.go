package gateways

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
)

// IReportsRepository .
type IReportsRepository interface {
	Update(id, title, body string, tags entities.Tags) error
}

// ReportsRepository .
type ReportsRepository struct {
	qiitaAPI IQiitaAPI
}

// NewReportsRepository .
func NewReportsRepository(api IQiitaAPI) *ReportsRepository {
	return &ReportsRepository{
		qiitaAPI: api,
	}
}

// Update .
func (r *ReportsRepository) Update(id, title, body string, tags entities.Tags) error {
	if err := r.qiitaAPI.UpdateItem(id, title, body, append(entities.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}}, tags...)); err != nil {
		return err
	}

	return nil
}
