package gateways

import (
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IReportsRepository .
type IReportsRepository interface {
	Update(id, title, body string, tags domain.Tags) error
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
func (r *ReportsRepository) Update(id, title, body string, tags domain.Tags) error {
	if err := r.qiitaAPI.UpdateItem(id, title, body, append(domain.Tags{{Name: "Qiita"}}, tags...)); err != nil {
		return err
	}

	return nil
}
