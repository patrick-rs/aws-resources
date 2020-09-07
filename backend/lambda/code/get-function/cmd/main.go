package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
	"resources/get-function/transformer"
)

var (
	dynamodbClient dynamodbiface.DynamoDBAPI
	dynamodbTable  string
)

func main() {
	region := os.Getenv("AWS_REGION")
	dynamodbTable = os.Getenv("DYNAMODB_TABLE_NAME")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		return
	}

	dynamodbClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return transformers.GetResources(req, dynamodbClient, dynamodbTable)
}
