package main

import (
    "os"
    "encoding/json"
    "log"

    "gopkg.in/go-playground/validator.v9"

    "github.com/uudashr/iso8601"
    "github.com/satori/go.uuid"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
    Guid string`json:"guid" validate:"required,uuid4"`
    Device string`json:"device" validate:"required,uuid4"`
    AppVersion string`json:"appVersion" validate:"required"`
    Start iso8601.Time`json:"start" validate:"required"`
    End iso8601.Time`json:"end" validate:"required"`
    Zone string`json:"zone" validate:"required"`
    Latitude float64`json:"latitude,omitempty"`
    Longitude float64`json:"longitude,omitempty"`
}

var validate *validator.Validate
var dynamo *dynamodb.DynamoDB

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    // Parse the request body
    var i Item
    err := json.Unmarshal([]byte(request.Body), &i)

    if err != nil {
        log.Print("Error parsing request body")
        log.Print(err)
        return events.APIGatewayProxyResponse{
            Body: "{ \"message\": \"Invalid request body\"}",
            StatusCode: 400,
        }, nil
    }

    // Assign a GUID for the item
    i.Guid = uuid.NewV4().String()

    // Read the app version from the header
    i.AppVersion = request.Headers["x-app-version"]

    // Validate the item
    err = validate.Struct(i)
    if err != nil {
        log.Print("Error validating request body")
        log.Print(err)
        return events.APIGatewayProxyResponse{
            Body: "{ \"message\": \"Invalid request body\"}",
            StatusCode: 400,
        }, nil
    }

    // Write to Dynamo
    av, err := dynamodbattribute.MarshalMap(i)
    input := &dynamodb.PutItemInput{
        Item: av,
        TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
    }

    _, err = dynamo.PutItem(input)

    if err != nil {
        log.Print("Error inserting item to dynamo")
        log.Print(err)
        return events.APIGatewayProxyResponse{
            Body: "{ \"message\": \"Internal error\"}",
            StatusCode: 500,
        }, nil
    }

    rb, err := json.Marshal(i)

    if err != nil {
        log.Print("Error creating response JSON")
        log.Print(err)
        return events.APIGatewayProxyResponse{
            Body: "{ \"message\": \"Internal error\"}",
            StatusCode: 500,
        }, nil
    }

    return events.APIGatewayProxyResponse{
        Body: string(rb),
        StatusCode: 200,
    }, nil
}

func main() {
    validate = validator.New()
    sess, _ := session.NewSession()
    dynamo = dynamodb.New(sess)

	lambda.Start(Handler)
}
