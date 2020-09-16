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

type model struct {
	DB *gorm.DB
}

type Content struct {
	ID        int        `json:"id"`
	TicketID  int        `json:"ticket_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Response struct {
	ID int `json:"id"`
}

func createConnect() string {
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(zendesk-db:3306)"
	DBNAME := "zendesk"
	OPTION := "charset=utf8mb4&parseTime=True&loc=Local"

	return fmt.Sprintf("%s:%s@%s/%s?%s", USER, PASS, PROTOCOL, DBNAME, OPTION)
}

func (db *model) open() error {
	var err error
	db.DB, err = gorm.Open(mysql.Open(createConnect()), nil)
	if err != nil {
		return err
	}
	return nil
}

func (db *model) hasTable() bool {
	return db.DB.Table("ticket_contents").Migrator().HasTable(&Content{})
}

func httpHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var db model
	if err := db.open(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Println(db.hasTable())
	if !db.hasTable() {
		err := db.DB.Migrator().CreateTable(&Content{})
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
	lambda.Start(httpHandler)
}
