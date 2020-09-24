package main

import (
	"encoding/json"
	"fmt"
	"search/table"

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

func searchHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := dbOpen()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	queries := request.QueryStringParameters

	if q, ok := queries["id"]; ok {
		fmt.Println("id")
		db = db.Where("id = ?", q)
	}

	if q, ok := queries["subject"]; ok {
		fmt.Println("subject")
		db = db.Where("subject like %?%", q)
	}

	if q, ok := queries["create"]; ok {
		fmt.Println("create")
		db = db.Where("create_time > ", q)
	}

	if q, ok := queries["type"]; ok {
		fmt.Println("type")
		db = db.Where("type = ?", q)
	}

	tickets := make([]table.Ticket, 0)

	if err := db.Find(&tickets).Error; err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	jsonTicket, err := json.Marshal(&tickets)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:       string(jsonTicket),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(searchHandler)
}
