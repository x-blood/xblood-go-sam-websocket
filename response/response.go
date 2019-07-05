package response

import "github.com/aws/aws-lambda-go/events"

func Create200response() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       string("ok"),
		StatusCode: 200,
	}, nil
}

func Create500response() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       string("ng"),
		StatusCode: 500,
	}, nil
}
