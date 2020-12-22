package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/interfaces/controllers"
)

// Event .
type Event struct {
	ItemID string `json:"item_id"`
	Tag    string `json:"tag"`
}

// Handler .
func Handler(ev *Event) error {
	c := controllers.NewItemsControllerFactory().Create()
	return c.UpdateWeekly(time.Now(), ev.ItemID, ev.Tag)
}

func main() {
	lambda.Start(Handler)
}
