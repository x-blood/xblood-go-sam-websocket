package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"xblood-go-sam-websocket/dynamodb"
	"xblood-go-sam-websocket/response"
)

type postData struct {
	Message string `json:"message"`
	Data    string `data:"message"`
}

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	var postData postData
	err = json.Unmarshal([]byte(request.Body), &postData)
	if err != nil {
		response.Create500response()
	}

	svc := apigatewaymanagementapi.New(newSession)
	svc.Endpoint = fmt.Sprintf("https://%s/%s", request.RequestContext.DomainName, request.RequestContext.Stage)

	for _, connection := range connections {
		connectionID := connection.ConnectionID
		svc.PostToConnection(&apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: &connectionID,
			Data:         []byte(postData.Data),
		})
	}

	fmt.Println("end send_message")
	return response.Create200response()
}

func main() {
	lambda.Start(HandleRequest)
}
