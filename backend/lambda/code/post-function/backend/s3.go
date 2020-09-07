package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"resources/types"
)

// QueryS3 finds all S3 buckets
func QueryS3(region string, dynamodbClient dynamodbiface.DynamoDBAPI, tableName string) error {

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		return err
	}

	s3Client := s3.New(awsSession)

	result, err := s3Client.ListBuckets(nil)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// S3 is not region specific
	for _, b := range result.Buckets {

		bucket := types.Bucket{
			PK:   "s3",
			SK:   aws.StringValue(b.Name),
			Name: aws.StringValue(b.Name),
		}

		err = PutDatabase(bucket, dynamodbClient, tableName)

		if err != nil {
			return err
		}
	}
	return nil
}
