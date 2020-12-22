package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler .
func Handler() error {
	return nil
}

func main() {
	lambda.Start(Handler)
}
