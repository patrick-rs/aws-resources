package transformers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func createResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {

	response := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"}}
	response.StatusCode = status

	stringBody, _ := json.Marshal(body)
	response.Body = string(stringBody)

	return &response, nil
}

// "Access-Control-Allow-Credentials": "true", "Access-Control-Allow-Headers": "Content-Type, x-api-key"
