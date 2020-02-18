package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"path"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

const s3putEvent = "ObjectCreated:Put"

func main(){
	lambda.Start(handler)
}


func handler(sqsEvent events.SQSEvent) error{
	for _, message := range sqsEvent.Records{
		ext, err := getExtFromMessage(message)
		if err != nil{
			log.Fatal(err)
		}
		log.Printf("extension of the file", ext)
	}
	return nil
}

func getExtFromMessage(e events.SQSMessage)(string, error){
	log.Printf("SQS message: %s", e.Body)

	var snsEvent events.SNSEntity
	if err := json.Unmarshal([]byte(e.Body), &snsEvent); err != nil{
		return "", errors.Wrapf(err, "failed to unmarshal: %s", e.Body)
	}
	log.Printf("SNS message %s", snsEvent.Message)
	var s3event events.S3Event
	if strings.Contains(snsEvent.Message, s3putEvent){
		return "", nil
	}
	if err := json.Unmarshal([]byte(snsEvent.Message), s3event); err != nil{
		return "", errors.Wrapf(err, "failed to unmarshal: %s", e.Body)
	}
	key, err := url.PathUnescape(s3event.Records[0].S3.Object.Key)
	if err != nil{
		return "", errors.Wrapf(err, "failed to unescaped file name: %s", s3event.Records[0].S3.Object.Key)
	}

	return path.Ext(key), nil

}