package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestNew() DB {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(Region),
		Region:   aws.String("http://localhost:8000")}),
	)
	return DB{Instance: dynamodb.New(sess)}
}

func (d DB) CreateLinkTable() error {

}
