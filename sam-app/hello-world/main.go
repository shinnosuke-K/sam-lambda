package main

import (
	"encoding/json"
	"fmt"
	"time"

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

type model struct {
	DB *gorm.DB
}

func (db *model) open() error {
	var err error
	db.DB, err = gorm.Open(mysql.Open(createConnect()), nil)
	if err != nil {
		return err
	}
	return nil
}

type Content struct {
	ID        int        `json:"id"`
	TicketID  int        `json:"ticket_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (db *model) hasTable() bool {
	return db.DB.Table("ticket_contents").Migrator().HasTable(&Content{})
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var db model

	if err := db.open(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	if !db.hasTable() {
		err := db.DB.Migrator().CreateTable(&Content{})
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprint("Hello"),
		StatusCode: 200,
	}, nil
}

type Response struct {
	ID int `json:"id"`
}

func httpHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var db model
	if err := db.open(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response := many()

	jsonBytes, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func many() []Response {

	res := make([]Response, 10000)

	for i := 0; i < 1000; i++ {
		res = append(res, Response{
			ID: i,
		})
	}
	return res
}

func main() {
	//lambda.Start(handler)

	lambda.Start(httpHandler)
}
