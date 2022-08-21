package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/koki-develop/qiita-lgtm-ranking/pkg/controllers"
)

type Event struct {
	ReportID string `json:"report_id"`
	Tag      string `json:"tag"`
}

func Handler(e *Event) error {
	ctrl := controllers.NewReportController(&controllers.ReportControllerConfig{
		QiitaAccessToken: os.Getenv("QIITA_ACCESS_TOKEN"),
	})

	if err := ctrl.UpdateWeeklyByTag(e.ReportID, e.Tag); err != nil {
		fmt.Printf("err: %+v\n", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
