package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"xblood-go-sam-websocket/dynamodb"
	"xblood-go-sam-websocket/response"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("start on_disconnect")
	connectionID := request.RequestContext.ConnectionID
	fmt.Printf("connectionId : %s ¥n", connectionID)
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
