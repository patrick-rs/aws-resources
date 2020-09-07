package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/lambda"
	"resources/types"
)

// QueryLambda ..
func QueryLambda(requestRegions []string, region string, dynamodbClient dynamodbiface.DynamoDBAPI, tableName string) error {

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		fmt.Println("ERROR: unable to create aws session")
		return err
	}

	for _, reg := range requestRegions {
		svc := lambda.New(awsSession, &aws.Config{Region: aws.String(reg)})

		result, err := svc.ListFunctions(nil)

		if err != nil {
			fmt.Printf("ERROR: listing Lambda functions, region: %s, error: %s", reg, err)
			return err
		}

		for _, f := range result.Functions {
			function := types.LambdaFn{
				PK:          "lambda",
				SK:          fmt.Sprintf("REG#%s#NAME#%s", reg, aws.StringValue(f.FunctionName)),
				Arn:         aws.StringValue(f.FunctionArn),
				Name:        aws.StringValue(f.FunctionName),
				Region:      reg,
				Description: aws.StringValue(f.Description),
				Runtime:     aws.StringValue(f.Runtime),
			}

			err = PutDatabase(function, dynamodbClient, tableName)

			if err != nil {
				return err
			}
		}
	}

	return nil

}
