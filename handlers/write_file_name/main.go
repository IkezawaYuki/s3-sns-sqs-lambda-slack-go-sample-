package write_file_name

import "github.com/aws/aws-lambda-go/events"

const s3putEvent = "ObjectCreated:put"

func main(){

}

func handler(sqsEvent events.SQSEvent) error{
	
}