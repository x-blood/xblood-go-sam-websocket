package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"xblood-go-sam-websocket/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) events.APIGatewayProxyResponse {
	fmt.Sprintf("start send_message")

	connections, err := dynamodb.GetAll()
	if err != nil {
		create500response()
	}

	var config *aws.Config
	newSession, err := session.NewSession(config)
	if err != nil {
		create500response()
	}

	svc := apigatewaymanagementapi.New(newSession)
	var testMessage = "test message"

	for _, connection := range connections {
		svc.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: &connection.ConnectionID,
			Data: []byte(testMessage),
		})
	}

	fmt.Sprintf("end send_message")
	return create200response()
}

func create200response() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string("ok"),
		StatusCode: 200,
	}
}

func create500response() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string("ng"),
		StatusCode: 500,
	}
}

func main() {
	lambda.Start(HandleRequest)
}
