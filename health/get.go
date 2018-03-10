package main

import (
    "encoding/json"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    body := map[string]string{"status": "OK"}
    bodyJson, _ := json.Marshal(body)
    return events.APIGatewayProxyResponse{Body: string(bodyJson), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
