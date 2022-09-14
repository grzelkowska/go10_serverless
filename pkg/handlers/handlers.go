package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/grzelkowska/11projects/10go_serverless/pkg/user"
)

var ErrorMethodNotAllowed = "method not allowed"

func GetUser() {

}

func CreateUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}

func UnhandleMethod()(*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}