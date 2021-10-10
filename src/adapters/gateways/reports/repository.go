package reports

import (
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/src/adapters/gateways"
	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
	"github.com/pkg/errors"
)

type Repository struct {
	config *Config
}

type Config struct {
	QiitaAPI      gateways.QiitaAPI
	ReportBuilder gateways.ReportBuilder
}

func New(cfg *Config) *Repository {
	return &Repository{config: cfg}
}

func (repo *Repository) UpdateWeekly(from time.Time, id string, items entities.Items) error {
	rpt, err := repo.config.ReportBuilder.Weekly(from, items)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := repo.update(id, rpt); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *Repository) UpdateWeeklyByTag(from time.Time, id string, items entities.Items, tag string) error {
	rpt, err := repo.config.ReportBuilder.WeeklyByTag(from, items, tag)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := repo.update(id, rpt); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *Repository) update(id string, rpt *entities.Report) error {
	if err := repo.config.QiitaAPI.UpdateItem(id, rpt.Title, rpt.Body, rpt.Tags); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
