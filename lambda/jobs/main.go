package main

import (
	"encoding/xml"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return handler(r), nil
	})
}

func handler(r events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var statusCode int
	if isValidXML([]byte(r.Body)) {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusBadRequest
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}
}

func isValidXML(data []byte) bool {
	return xml.Unmarshal(data, new(interface{})) == nil
}
