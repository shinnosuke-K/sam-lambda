package main

import (
	"encoding/json"
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

func open() error {
	db, err := gorm.Open(mysql.Open(createConnect()), nil)
	if err != nil {
		return err
	}
	return nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if err := open(); err != nil {
		return events.APIGatewayProxyResponse{}, err
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

	fmt.Println(request)

	if err := open(); err != nil {
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
