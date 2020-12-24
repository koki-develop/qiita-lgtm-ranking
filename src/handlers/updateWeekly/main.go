package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/controllers"
)

// Event .
type Event struct {
	ReportID string `json:"report_id"`
}

// Handler .
func Handler(ev *Event) error {
	c := controllers.NewReportsControllerFactory().Create()
	return c.UpdateWeekly(time.Now(), ev.ReportID)
}

func main() {
	lambda.Start(Handler)
}
