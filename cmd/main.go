package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/grzelkowska/11projects/10go_serverless/pkg/handlers"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

// const tableName = "LambdaInGoUser"
const tableName = "go-serverless00"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}


// curl --header "Content-Type: application/json" --request POST --data '{"email": "grzelk@grze.lk", "firstName": "Grzelk", "lastName": "Lee"}' https://ccjigwgxf7.execute-api.ap-northeast-2.amazonaws.com/staging
// curl -X GET https://https://ccjigwgxf7.execute-api.ap-northeast-2.amazonaws.com/staging
// curl -X GET https://ccjigwgxf7.execute-api.ap-northeast-2.amazonaws.com/staging\?email\=grzelk@grze.lk
// curl --header "Content-Type: application/json" --request PUT --data '{"email": "grzelk@grze.lk", "firstName": "Phil", "lastName": "Rhee"}' https://ccjigwgxf7.execute-api.ap-northeast-2.amazonaws.com/staging
// curl -X DELETE https://ccjigwgxf7.execute-api.ap-northeast-2.amazonaws.com/staging\?email\=grzelk@grze.lk  