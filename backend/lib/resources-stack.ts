import cdk = require ("@aws-cdk/core")
import lambda = require("@aws-cdk/aws-lambda")
import assets = require("@aws-cdk/aws-s3-assets")
import path = require("path")
import iam = require('@aws-cdk/aws-iam')
import dynamodb = require('@aws-cdk/aws-dynamodb')
import apigw = require('@aws-cdk/aws-apigateway')

export class ResourcesStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const dynamoTable = new dynamodb.Table(this, 'resourcesTable', {
      partitionKey: {
        name: 'PK',
        type: dynamodb.AttributeType.STRING
      },
      sortKey: {
        name: 'SK',
        type: dynamodb.AttributeType.STRING
      },
      tableName: 'resourcesTable',
    })

    const getFunctionCode = new assets.Asset(this, "GetFunction", {
      path: path.join(__dirname, "../lambda/build/get-function")
    })

    const postFunctionCode = new assets.Asset(this, "PostFunction", {
      path: path.join(__dirname,"../lambda/build/post-function" )
    })

    /*
    const role = new iam.Role(this, 'LambdaExecutionRole', {
      assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com')
    })
    */

    const dynamodbLambdaPolicy = new iam.PolicyStatement()
    dynamodbLambdaPolicy.addActions("dynamodb:*")
    dynamodbLambdaPolicy.addResources(dynamoTable.tableArn)

    const lambdaLambdaPolicy = new iam.PolicyStatement()
    lambdaLambdaPolicy.addActions("lambda:ListFunctions")
    lambdaLambdaPolicy.addAllResources()

    const ec2LambdaPolicy = new iam.PolicyStatement()
    ec2LambdaPolicy.addActions("ec2:DescribeInstances")
    ec2LambdaPolicy.addAllResources()

    const s3LambdaPolicy = new iam.PolicyStatement()
    s3LambdaPolicy.addActions("s3:ListAllMyBuckets")
    s3LambdaPolicy.addAllResources()

    const listDynamoPolicy  = new iam.PolicyStatement()
    listDynamoPolicy.addActions("dynamodb:ListTables")
    listDynamoPolicy.addAllResources()

    const listAPIGW = new iam.PolicyStatement()
    listAPIGW.addActions("apigateway:GET")
    listAPIGW.addAllResources()
    
    const lambdaGetFunction = new lambda.Function(this, "LambdaGetFunction", {
      code: lambda.Code.fromBucket(
        getFunctionCode.bucket,
        getFunctionCode.s3ObjectKey
      ),
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
    })

    const lambdaPostFunction = new lambda.Function(this, "LambdaPostFunction", {
      code: lambda.Code.fromBucket(postFunctionCode.bucket,postFunctionCode.s3ObjectKey),
      runtime: lambda.Runtime.GO_1_X,
      handler: "main",
    })

    lambdaGetFunction.addEnvironment("DYNAMODB_TABLE_NAME", dynamoTable.tableName)
    lambdaPostFunction.addEnvironment("DYNAMODB_TABLE_NAME", dynamoTable.tableName)

    lambdaGetFunction.addToRolePolicy(dynamodbLambdaPolicy)

    lambdaPostFunction.addToRolePolicy(s3LambdaPolicy)
    lambdaPostFunction.addToRolePolicy(dynamodbLambdaPolicy)
    lambdaPostFunction.addToRolePolicy(lambdaLambdaPolicy)
    lambdaPostFunction.addToRolePolicy(ec2LambdaPolicy)
    lambdaPostFunction.addToRolePolicy(listDynamoPolicy)
    lambdaPostFunction.addToRolePolicy(listAPIGW)

    const api = new apigw.RestApi(this, 'resources-api', {
      deploy: false,
    })


    const apiKey = new apigw.ApiKey(this, 'ResourcesAPIKey', {
      enabled: true,
      value: "d2e5183861bb4c72b89acde4a2c3c3af",
    })

    const APIDeployment = new apigw.Deployment(this, "Deployment", {
      api
    })

    const devStage = new apigw.Stage(this, "DevStage", {
      stageName: "dev",
      deployment: APIDeployment,
      
    })

    const defaultUsagePlan = new apigw.UsagePlan(this, "UsagePlan", {
      apiKey: apiKey,
    })

    defaultUsagePlan.addApiStage({
      api: api,
      stage: devStage
    })

    const resources = api.root.addResource('resources')

    resources.addCorsPreflight({
      allowOrigins: apigw.Cors.ALL_ORIGINS,
      allowHeaders: apigw.Cors.DEFAULT_HEADERS,
      allowCredentials: true,
      allowMethods: apigw.Cors.ALL_METHODS
    })

    // const apiKey = new apigw.CfnApiKey(this, 'resources api key')

    const getResourcesIntegration = new apigw.LambdaIntegration(lambdaGetFunction)
    const postResourcesIntegration = new apigw.LambdaIntegration(lambdaPostFunction)

    resources.addMethod('POST', postResourcesIntegration, {
      apiKeyRequired: true
    })
    resources.addMethod('GET', getResourcesIntegration, {
      apiKeyRequired: true,
    })



    new cdk.CfnOutput(this, 'ResponseMessage', {
      description: 'URL',
      value: `https://${api.restApiId}.execute-api.${this.region}.amazonaws.com/dev`
    })
}
}
