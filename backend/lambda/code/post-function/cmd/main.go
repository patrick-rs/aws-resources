package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
	"resources/post-function/transformer"
)

var (
	dynamodbClient dynamodbiface.DynamoDBAPI
	dynamodbTable  string
	awsSession     session.Session
	region         string
)

func main() {
	region := os.Getenv("AWS_REGION")
	dynamodbTable = os.Getenv("DYNAMODB_TABLE_NAME")

	//initialize dynamodb client
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	fmt.Println("from main: ", awsSession)

	if err != nil {
		fmt.Printf("Could not create aws session, %s", err)
		return
	}
	dynamodbClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return transformers.PostResources(req, region, dynamodbClient, dynamodbTable)
}
