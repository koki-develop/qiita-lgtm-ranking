package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/controllers/reports"
)

// Event .
type Event struct {
	ReportID string `json:"report_id"`
}

// Handler .
func Handler(ev *Event) error {
	ctrl := reports.New()
	return ctrl.UpdateWeekly(time.Now(), ev.ReportID)
}

func main() {
	lambda.Start(Handler)
}
