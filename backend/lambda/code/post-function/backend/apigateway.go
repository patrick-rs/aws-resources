package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"resources/types"
)

//QueryAPIGW lists api gateway
func QueryAPIGW(requestRegions []string, region string, dynamodbClient dynamodbiface.DynamoDBAPI, tableName string) error {
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		fmt.Println("ERROR: unable to create aws session")
		return err
	}

	for _, reg := range requestRegions {
		apigwClient := apigateway.New(awsSession, &aws.Config{Region: aws.String(reg)})

		result, err := apigwClient.GetRestApis(nil)

		if err != nil {
			fmt.Printf("ERROR: unable to list apigateways, region: %s, error: %s", reg, err)
			return err
		}

		for _, a := range result.Items {
			apigw := types.APIGW{
				PK:          "apigw",
				SK:          fmt.Sprintf("#REG%s#NAME#%s", reg, aws.StringValue(a.Name)),
				Name:        aws.StringValue(a.Name),
				Description: aws.StringValue(a.Description),
				ID:          aws.StringValue(a.Id),
			}

			err = PutDatabase(apigw, dynamodbClient, tableName)

			if err != nil {
				return nil
			}
		}
	}

	return nil
}
