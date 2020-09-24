package main

import (
	"fmt"

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

	return events.APIGatewayProxyResponse{
		Body:       "ok",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(searchHandler)
}