package main

import (
	"hello-world/model"
	"hello-world/table"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func httpHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db := model.New()

	has, err := db.HasTable(table.Ticket{})
	if !has && err == nil {
		err := db.DB.Migrator().CreateTable(&table.Ticket{})
		if err != nil {
			return events.APIGatewayProxyResponse{Body: err.Error()}, err
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       string(1),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(httpHandler)
}
