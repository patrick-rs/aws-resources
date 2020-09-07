package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"resources/types"
)

// QueryTables Lists tables and returns them
func QueryTables(requestRegions []string, region string, dynamodbClient dynamodbiface.DynamoDBAPI, tableName string) error {
	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		fmt.Println("ERROR: unable to create aws session")
		return err
	}

	for _, reg := range requestRegions {
		dynamodbClient := dynamodb.New(awsSession, &aws.Config{Region: aws.String(reg)})

		result, err := dynamodbClient.ListTables(nil)

		if err != nil {
			fmt.Printf("ERROR: listing dynamodb tables, region: %s, error: %s", reg, err)
			return err
		}

		for _, n := range result.TableNames {
			input := dynamodb.DescribeTableInput{TableName: n}
			output, err := dynamodbClient.DescribeTable(&input)

			if err != nil {
				return err
			}
			table := types.DynamoTable{
				PK:     "dynamodb",
				SK:     fmt.Sprintf("REG#%s#NAME#%s", reg, aws.StringValue(output.Table.TableName)),
				Arn:    aws.StringValue(output.Table.TableArn),
				Name:   aws.StringValue(output.Table.TableName),
				Region: reg,
				RCU:    aws.Int64Value(output.Table.ProvisionedThroughput.ReadCapacityUnits),
				WCU:    aws.Int64Value(output.Table.ProvisionedThroughput.WriteCapacityUnits),
			}

			err = PutDatabase(table, dynamodbClient, tableName)

			if err != nil {
				return err
			}
		}
	}
	return nil
}
