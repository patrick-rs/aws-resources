package transformers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"net/http"
	"resources/get-function/backend"
	"resources/types"
)

// GetResources takes in a json array of resource types, unmarshals them, and passes them onto the database
func GetResources(request events.APIGatewayProxyRequest, dynamodbClient dynamodbiface.DynamoDBAPI, dynamodbTable string) (*events.APIGatewayProxyResponse, error) {

	query := request.QueryStringParameters["query"]

	/*
		if query == "all" {

			result, err := backend.FetchResources(dynamodbClient, dynamodbTable)
			if err != nil {
				return createResponse(http.StatusBadRequest, types.ErrorBody{
					ErrorMsg: aws.String(err.Error()),
				})
			}
			return createResponse(http.StatusOK, result)

		}

	*/
	result, err := backend.FetchResource(query, dynamodbClient, dynamodbTable)
	if err != nil {
		return createResponse(http.StatusBadRequest, types.ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return createResponse(http.StatusOK, result)
}
