package controllers

import (
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
)

type ReportsRepository interface {
	UpdateDaily(t time.Time, id string, items entities.Items) error
	UpdateWeekly(t time.Time, id string, items entities.Items) error
	UpdateWeeklyByTag(t time.Time, id string, items entities.Items, tag string) error
}
