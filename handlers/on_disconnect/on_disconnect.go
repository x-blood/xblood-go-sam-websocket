package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"xblood-go-sam-websocket/dynamodb"
)

func HandleRequest(request events.APIGatewayWebsocketProxyRequest) {
	fmt.Sprintf("start on_disconnect")
	connectionID := request.RequestContext.ConnectionID
	dynamodb.Delete(connectionID)
	fmt.Sprintf("end on_disconnect")
}

func main() {
	lambda.Start(HandleRequest)
}
