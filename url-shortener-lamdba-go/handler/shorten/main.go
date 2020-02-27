package main

import (
	"github.com/IkezawaYuki/s3-sns-sqs-lambda-slack-go-sample-/url-shortener-lamdba-go/db"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type reqeust struct {
	URL string `json:"url"`
}

type Response struct {
	ShortenResource string `json:"shorten_resource"`
}

type Link struct {
	ShortenResource string `json:"shorten_resource"`
	OriginURL       string `json:"origin_url"`
}

var DynamoDB db.DB

func init() {
	DynamoDB = db.New()
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	p, err := parseRequest(request)
	if err != nil {
		return nil, err
	}
}

func parseRequest(req events.APIGatewayProxyResponse) (*reqeust, error) {

}
