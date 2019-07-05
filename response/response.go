package response

import "github.com/aws/aws-lambda-go/events"

func Create200response() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string("ok"),
		StatusCode: 200,
	}
}

func Create500response() events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       string("ng"),
		StatusCode: 500,
	}
}