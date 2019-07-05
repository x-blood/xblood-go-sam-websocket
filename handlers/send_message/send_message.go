package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"xblood-go-sam-websocket/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"xblood-go-sam-websocket/response"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) events.APIGatewayProxyResponse {
	fmt.Println("start send_message")

	connections, err := dynamodb.GetAll()
	if err != nil {
		response.Create500response()
	}

	var config *aws.Config
	newSession, err := session.NewSession(config)
	if err != nil {
		response.Create500response()
	}

	svc := apigatewaymanagementapi.New(newSession)
	var testMessage = "test message"

	for _, connection := range connections {
		svc.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: &connection.ConnectionID,
			Data: []byte(testMessage),
		})
	}

	fmt.Println("end send_message")
	return response.Create200response()
}

func main() {
	lambda.Start(HandleRequest)
}
