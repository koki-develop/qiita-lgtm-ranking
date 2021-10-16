package gateways

import (
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
)

type ReportBuilder interface {
	Daily(from time.Time, items entities.Items) (*entities.Report, error)
	Weekly(from time.Time, items entities.Items) (*entities.Report, error)
	WeeklyByTag(from time.Time, items entities.Items, tag string) (*entities.Report, error)
}
