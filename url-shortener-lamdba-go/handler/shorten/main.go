package main

import (
	"encoding/json"
	"fmt"
	"github.com/IkezawaYuki/s3-sns-sqs-lambda-slack-go-sample-/url-shortener-lamdba-go/db"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"

	"net/http"
	"net/url"
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
		return response(
			http.StatusBadRequest,
			errorResponseBody(err.Error()),
		), nil
	}
	shortenResource := shortid.MutGenerate()

	for shortenResource == "shorten" {
		shortenResource = shortid.MustGenerate()
	}
	link := &Link{
		ShortenResource: shortenResource,
		OriginURL:       p.URL,
	}

	_, err = DynamoDB.PutItem(link)
	if err != nil {
		return response(
			http.StatusInternalServerError,
			errorResponseBody(err.Error()),
		), nil
	}
	b, err := responseBody(shortenResource)
	if err != nil {
		return response(
			http.StatusInternalServerError,
			errorResponseBody(err.Error()),
		), nil
	}
	return response(http.StatusOK, b), nil
}

func parseRequest(req events.APIGatewayProxyRequest) (*reqeust, error) {
	if req.HTTPMethod != http.MethodPost {
		return nil, fmt.Errorf("use POST request")
	}

	var r reqeust
	err := json.Unmarshal([]byte(req.Body), &r)
	if err != nil {
		return nil, errors.Wrapf(err, "filed to parse request")
	}
	_, err = url.ParseRequestURI(r.URL)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid URL")
	}

	return &r, nil
}

func response(code int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}

func responseBody(shortenResource string) (string, error) {
	resp, err := json.Marshal(Response{ShortenResource: shortenResource})
	if err != nil {
		return "", nil
	}

	return string(resp), nil
}

func errorResponseBody(msg string) string {
	return fmt.Sprintf("{\"message\":\"%s\"}", msg)
}
