package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/adapters/controllers"
)

// Event .
type Event struct {
	ReportID string `json:"report_id"`
	Tag      string `json:"tag"`
}

// Handler .
func Handler(ev *Event) error {
	c := controllers.NewReportsControllerFactory().Create()
	return c.UpdateWeeklyByTag(time.Now(), ev.ReportID, ev.Tag)
}

func main() {
	lambda.Start(Handler)
}
