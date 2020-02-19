package main

import (
	"github.com/IkezawaYuki/s3-sns-sqs-lambda-slack-go-sample-/handlers/notifier/slack"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

var client *slack.Client

func main() {
	lambda.Start(handler)
}

func init() {
	client = slack.NewClient(
		slack.Config{
			URL:       "",
			Username:  "",
			IconEmoji: "",
			Channel:   "",
		},
	)
}

func handler(snsEvent events.SNSEvent) error {
	record := snsEvent.Records[0]
	snsRecord := snsEvent.Records[0].SNS
	log.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

	if err := client.PostMessage(snsRecord.Message); err != nil {
		return err
	}
	return nil
}
