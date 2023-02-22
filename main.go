package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var person Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	msg := fmt.Sprintf("Hello %v %v", *person.FirstName, *person.SecondName)
	responseBody := ResponseBody{
		Message: &msg,
	}
	jbytes, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}
	return res, nil
}

type Person struct {
	FirstName  *string `json:"firstName"`
	SecondName *string `json:"secondName"`
}
type ResponseBody struct {
	Message *string `json:"message"`
}
