package main

import (
	"encoding/json"
	"fmt"
	"search/table"
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

func dbOpen() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(createConnect()), nil)
}

func createSQL(params map[string]string, db *gorm.DB) (*gorm.DB, error) {

	if p, ok := params["id"]; ok {
		db = db.Where("id = ?", p)
		if err := db.Error; err != nil {
			return nil, err
		}
	}

	if p, ok := params["subject"]; ok {
		db = db.Where("subject LIKE ?", "%"+p+"%")
		if err := db.Error; err != nil {
			return nil, err
		}
	}

	if _, ok := params["create"]; ok {
		db = db.Where("create_time > ?", time.RFC3339)
		if err := db.Error; err != nil {
			return nil, err
		}
	}

	if p, ok := params["type"]; ok {
		db = db.Where("type = ?", p)
		if err := db.Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}

func searchHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	db, err := dbOpen()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	params := request.QueryStringParameters

	db, err = createSQL(params, db)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
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
