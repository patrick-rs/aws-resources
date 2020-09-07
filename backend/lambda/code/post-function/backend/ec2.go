package backend

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"resources/types"
)

// QueryEc2  ..
func QueryEc2(requestRegions []string, region string, dynamodbClient dynamodbiface.DynamoDBAPI, tableName string) error {

	awsSession, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		fmt.Println("ERROR1: unable to create aws session")
		return err
	}

	// lambdaFns, err := lambdaResource(awsSession, region, regions, dynamodbTable)
	for _, reg := range requestRegions {

		svc := ec2.New(awsSession, &aws.Config{Region: aws.String(reg)})

		result, err := svc.DescribeInstances(nil)

		if err != nil {
			fmt.Printf("ERROR2 listing EC2 instances, region: %s, error: %s", reg, err)
			return err
		}

		for _, res := range result.Reservations {
			for _, inst := range res.Instances {
				instance := types.Ec2Instance{
					PK:           "ec2",
					SK:           fmt.Sprintf("#REG%s#ID#%s", reg, aws.StringValue(inst.InstanceId)),
					Region:       reg,
					InstanceType: aws.StringValue(inst.InstanceType),
					PublicDNS:    aws.StringValue(inst.PublicDnsName),
					IPv4PublicIP: aws.StringValue(inst.PublicIpAddress),
				}

				err = PutDatabase(instance, dynamodbClient, tableName)

				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}
