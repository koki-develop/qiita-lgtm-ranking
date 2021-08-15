package gateways

import (
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
)

type ReportBuilder interface {
	Weekly(from time.Time, items entities.Items) (*entities.Report, error)
	WeeklyByTag(from time.Time, items entities.Items, tag string) (*entities.Report, error)
}
