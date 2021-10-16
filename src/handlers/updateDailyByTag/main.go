package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koki-develop/qiita-lgtm-ranking/src/adapters/controllers/reports"
)

// Event .
type Event struct {
	ReportID string `json:"report_id"`
	Tag      string `json:"tag"`
}

// Handler .
func Handler(ev *Event) error {
	ctrl := reports.New()
	return ctrl.UpdateDailyByTag(time.Now(), ev.ReportID, ev.Tag)
}

func main() {
	lambda.Start(Handler)
}