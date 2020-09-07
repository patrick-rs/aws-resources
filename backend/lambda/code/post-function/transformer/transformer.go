package transformers

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"net/http"
	"resources/post-function/backend"
)

// ErrorBody struct
type ErrorBody struct {
	ErrorMsg *string `jsong:"error,omitempy"`
}

// PostResources takes in regions and writes to dynamodb
func PostResources(req events.APIGatewayProxyRequest, region string, dynamodbClient dynamodbiface.DynamoDBAPI, tableName string) (*events.APIGatewayProxyResponse, error) {

	request := struct{ Regions []string }{}

	resource := req.QueryStringParameters["resource"]

	err := json.Unmarshal([]byte(req.Body), &request)

	requestRegions := request.Regions

	if err != nil {
		fmt.Println("ERROR: unmarshaling request", req.Body)
		return createResponse(http.StatusBadRequest, err)
	}

	if resource == "" {
		fmt.Printf("ERROR: no resource provided")
		return createResponse(http.StatusBadRequest, nil)
	}

	switch resource {
	case "lambda":
		err = backend.QueryLambda(requestRegions, region, dynamodbClient, tableName)
	case "ec2":
		err = backend.QueryEc2(requestRegions, region, dynamodbClient, tableName)
	case "s3":
		err = backend.QueryS3(region, dynamodbClient, tableName)
	case "dynamodb":
		err = backend.QueryTables(requestRegions, region, dynamodbClient, tableName)
	case "apigw":
		err = backend.QueryAPIGW(requestRegions, region, dynamodbClient, tableName)
	default:
		fmt.Printf("ERROR: bad request resource string, %s", resource)
		return createResponse(http.StatusBadRequest, err)
	}

	if err != nil {
		fmt.Printf("ERROR: %s", err)
		return createResponse(http.StatusInternalServerError, err)
	}

	return createResponse(http.StatusOK, nil)
}
