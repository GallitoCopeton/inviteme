package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the name parameter from the request body
	name := request.QueryStringParameters["name"]

	// Create the response body
	message := fmt.Sprintf("Hello, %s!", name)

	// Create the API Gateway response
	response := events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 200,
	}

	return response, nil
}

func main() {
	// Start the Lambda function handler
	lambda.Start(handler)
}
