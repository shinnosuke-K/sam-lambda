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

	//s := table.NewTicket()
	//s.

	response := many()

	jsonBytes, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(httpHandler)
}
