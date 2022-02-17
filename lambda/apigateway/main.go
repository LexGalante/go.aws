package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Result struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	weigth, err := strconv.ParseFloat(request.QueryStringParameters["weigth"], 32)
	if err != nil {
		return badRequest("the parameter [weight] cannot be convert to int", err)
	}

	height, err := strconv.ParseFloat(request.QueryStringParameters["height"], 32)
	if err != nil {
		return badRequest("the parameter [height] cannot be convert to int", err)
	}

	imc := weigth / height * 2

	switch {
	case imc > 40.0:
		return ok(fmt.Sprintf("IMC = %v, obesity III", imc))
	case imc <= 39.9 && imc >= 30.0:
		return ok(fmt.Sprintf("IMC = %v, obesity II", imc))
	case imc <= 29.9 && imc >= 25.0:
		return ok(fmt.Sprintf("IMC = %v, obesity I", imc))
	case imc < 18.5:
		return ok(fmt.Sprintf("IMC = %v, very slim", imc))
	default:
		return ok(fmt.Sprintf("IMC = %v, normal", imc))
	}
}

func badRequest(message string, err error) (events.APIGatewayProxyResponse, error) {
	result := Result{false, message}

	body, _ := json.Marshal(&result)

	log.Println("error:", err.Error())

	return events.APIGatewayProxyResponse{StatusCode: 400, Body: string(body)}, err
}

func ok(message string) (events.APIGatewayProxyResponse, error) {
	result := Result{true, message}

	body, _ := json.Marshal(&result)

	return events.APIGatewayProxyResponse{StatusCode: 400, Body: string(body)}, nil
}
