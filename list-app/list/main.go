package main

import (
	"encoding/json"
	"fmt"
	"list/table"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func createConnect() string {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(zendesk-db:3306)"
	DBNAME := "zendesk"
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"

	return fmt.Sprintf("%s:%s@%s/%s?%s", USER, PASS, PROTOCOL, DBNAME, OPTION)
}

func dbOpen() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(createConnect()), nil)
}

func listHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := dbOpen()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	tickets := make([]table.Ticket, 0)

	if err := db.Find(&tickets).Error; err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	jsonTicket, err := json.Marshal(&tickets)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(jsonTicket),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(listHandler)
}
