package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Result struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		var result Result
		err := json.Unmarshal([]byte(message.Body), &result)
		if err != nil {
			return err
		}

		if result.Status {
			storeInDb(result)
		}
	}

	return nil
}

func storeInDb(result Result) {
	//TODO
}
