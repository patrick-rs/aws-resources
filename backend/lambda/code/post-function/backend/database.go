package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// PutDatabase takes in an array of items and attempts to put them in the database
func PutDatabase(payload interface{}, dynamodbClient dynamodbiface.DynamoDBAPI, dynamodbTable string) error {
	av, err := dynamodbattribute.MarshalMap(payload)

	if err != nil {
		return fmt.Errorf("ERROR marshaling, data: %s, error: %s", payload, err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(dynamodbTable),
	}

	_, err = dynamodbClient.PutItem(input)

	if err != nil {
		return fmt.Errorf("ERROR putting item in the database, input: %s, error: %s", input, err)
	}
	return nil
}
