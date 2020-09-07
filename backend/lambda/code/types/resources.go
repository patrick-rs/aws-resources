package types

// LambdaFn .
type LambdaFn struct {
	PK          string // lambda
	SK          string // #REG<region>#NAME#<name>
	Arn         string
	Name        string
	Region      string
	Description string
	Runtime     string
}

// Ec2Instance ..
type Ec2Instance struct {
	PK           string // ec2
	SK           string // #REG<region>#ID#<instance-id>
	Region       string
	InstanceType string
	PublicDNS    string
	IPv4PublicIP string
}

// Bucket ..
type Bucket struct {
	PK   string // s3
	SK   string // name
	Name string
}

// DynamoTable struct
type DynamoTable struct {
	PK     string
	SK     string
	Arn    string
	Name   string
	Region string
	RCU    int64
	WCU    int64
}

// APIGW .
type APIGW struct {
	PK          string
	SK          string
	Name        string
	Description string
	ID          string
}
