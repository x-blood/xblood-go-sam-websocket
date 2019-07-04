package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"xblood-go-sam-websocket/dynamodb"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) {
	fmt.Sprintf("start on_connect")
	connectionID := request.RequestContext.ConnectionID
	dynamodb.Put(connectionID)
	fmt.Sprintf("end on_connect")
}

func main() {
	lambda.Start(HandleRequest)
}
