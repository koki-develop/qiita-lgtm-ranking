package controllers

import (
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/entities"
)

type ReportsRepository interface {
	UpdateWeekly(t time.Time, id string, items entities.Items) error
	UpdateWeeklyByTag(t time.Time, id string, items entities.Items, tag string) error
}
