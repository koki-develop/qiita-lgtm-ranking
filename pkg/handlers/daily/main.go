package main

import (
	"fmt"
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

	if err := ctrl.UpdateDaily(e.ReportID); err != nil {
		fmt.Printf("err: %+v\n", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
