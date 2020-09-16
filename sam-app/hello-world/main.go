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

	//record := table.Ticket{
	//	ID:             0,
	//	CreateTime:     time.Now().Local(),
	//	UpdateTime:     time.Now().Local(),
	//	Type:           "",
	//	Subject:        "",
	//	Priority:       "",
	//	Status:         "",
	//	Tag:            "",
	//	RequesterID:    0,
	//	AssigneeID:     0,
	//	OrganizationID: 0,
	//}

	return events.APIGatewayProxyResponse{
		Body:       string(1),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(httpHandler)
}
