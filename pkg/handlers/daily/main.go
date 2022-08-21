package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koki-develop/qiita-lgtm-ranking/pkg/controllers"
)

type Event struct {
	ReportID string `json:"report_id"`
}

func Handler(e *Event) error {
	ctrl := controllers.NewReportController(&controllers.ReportControllerConfig{
		QiitaAccessToken: os.Getenv("QIITA_ACCESS_TOKEN:"),
	})

	return ctrl.UpdateDaily(e.ReportID)
}

func main() {
	lambda.Start(Handler)
}
