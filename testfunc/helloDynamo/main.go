package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/rs/xid"
)

type Response struct {
	Message string `json:"message"`
}

func storeData() {
	// session
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)
	// PutItem = insert
	guid := xid.New()
	putParams := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("DATATABLE")),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(guid.String()),
			},
			"Title": {
				S: aws.String("testtest"),
			},
		},
	}

	_, putErr := svc.PutItem(putParams)
	if putErr != nil {
		panic(putErr)
	}
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	storeData()
	guid := xid.New()
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %#v %#v", os.Getenv("SEQUENCETABLE"), guid.String()),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
