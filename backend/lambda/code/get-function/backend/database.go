package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"resources/types"
)

// FetchResources Fetches all resources of each source provided
/*
func FetchResources(dynamodbClient dynamodbiface.DynamoDBAPI, dynamodbTable string) (*[]types.Resource, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(dynamodbTable),
	}

	result, err := dynamodbClient.Scan(input)

	if err != nil {
		return nil, errors.New("Failed to fetch records")
	}

	resources := new([]types.Resource)

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, resources)

	if err != nil {
		return nil, errors.New("Failed to unmarshal data from dynamodb")
	}

	return resources, nil
}
*/

// FetchResource queries for all keys that begin with a keyword
func FetchResource(queryString string, dynamodbClient dynamodbiface.DynamoDBAPI, dynamodbTable string) (interface{}, error) {

	keyCond := expression.Key("PK").Equal(expression.Value(queryString))

	expression, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()

	if err != nil {
		return nil, fmt.Errorf("Failed to format a valid query: %s", err)
	}
	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(dynamodbTable),
		KeyConditionExpression:    expression.KeyCondition(),
		ExpressionAttributeValues: expression.Values(),
		FilterExpression:          expression.Filter(),
		ExpressionAttributeNames:  expression.Names(),
	}

	result, err := dynamodbClient.Query(queryInput)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch records: %s", err)
	}

	switch queryString {
	case "lambda":
		var resources []types.LambdaFn
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resources)
		if err != nil {
			return nil, err
		}
		return resources, nil
	case "ec2":
		var resources []types.Ec2Instance
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resources)
		if err != nil {
			return nil, err
		}
		return resources, nil

	case "s3":
		var resources []types.Bucket
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resources)
		if err != nil {
			return nil, err
		}
		return resources, nil

	case "dynamodb":
		var resources []types.DynamoTable
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resources)
		if err != nil {
			return nil, err
		}
		return resources, nil

	case "apigw":
		var resources []types.APIGW
		err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resources)
		if err != nil {
			return nil, err
		}
		return resources, nil
	}

	return nil, fmt.Errorf("Unsupported request")
}
