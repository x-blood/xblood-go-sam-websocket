package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"xblood-go-sam-websocket/dynamodb"
	"xblood-go-sam-websocket/response"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) events.APIGatewayProxyResponse {
	fmt.Println("start on_disconnect")
	connectionID := request.RequestContext.ConnectionID
	err := dynamodb.Delete(connectionID)
	if err != nil {
		fmt.Println(err)
		return response.Create500response()
	}
	fmt.Println("end on_disconnect")
	return response.Create200response()
}

func main() {
	lambda.Start(HandleRequest)
}
