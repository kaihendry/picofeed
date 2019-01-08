package main

import (
	"context"

	"github.com/apex/log"
	jsonhandler "github.com/apex/log/handlers/json"
	"github.com/kaihendry/picofeed"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.SetHandler(jsonhandler.Default)
	lambda.Start(handler)
}

func handler(ctx context.Context, evt events.SNSEvent) (string, error) {
	err := picofeed.Refresh(ctx)
	return "", err
}
