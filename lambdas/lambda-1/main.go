package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
)

var logger = log.Default()

func handler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Printf("%v", e)

	return buildSuccessResponse(map[string]string{"hello": "world"}), nil
}

func buildSuccessResponse(uri any) events.APIGatewayProxyResponse {
	b, err := json.Marshal(uri)
	if err != nil {
		logger.Printf("error marshalling success response: %s", err.Error())
		return buildErrorResponse(errors.New("unable to generate response"), http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        http.StatusOK,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              string(b),
		IsBase64Encoded:   false,
	}
}

func buildErrorResponse(e error, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode:        statusCode,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              fmt.Sprintf("{ \"errorMsg\": \"%s\" }", e.Error()),
		IsBase64Encoded:   false,
	}
}

func main() {
	lambda.Start(handler)
}
