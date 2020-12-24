package gateways

import (
	"fmt"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IReportsRepository .
type IReportsRepository interface {
	Update(id, body, tag string) error
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
func (r *ReportsRepository) Update(id, body, tag string) error {
	if err := r.qiitaAPI.UpdateItem(id, fmt.Sprintf("【%s】Qiita 週間LGTM数ランキング【自動更新】", tag), body, domain.Tags{{Name: "Qiita"}, {Name: tag}}); err != nil {
		return err
	}

	return nil
}
