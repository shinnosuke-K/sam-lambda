package main

import (
	"all/domain"
	"all/infrastructure/persistence"
	api "all/interface"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func allhandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := persistence.NewDB()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// ticket
	ticket := persistence.NewTicket(db)

	if !ticket.Has() {
		if err := ticket.CreateTable(); err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	ticketRes := new(api.TicketsResponse)
	byteBody, err := api.GetBody(ticketRes)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	parsedTicket, err := api.Parse(ticketRes, byteBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	ticketRes, ok := parsedTicket.(*api.TicketsResponse)
	if !ok {
		return events.APIGatewayProxyResponse{}, errors.New("error")
	}

	for _, t := range ticketRes.Tickets {
		ticket.Tickets = append(ticket.Tickets, domain.Ticket{
			ID:             t.ID,
			CreateTime:     t.CreateTime.Local(),
			UpdateTime:     t.UpdateTime.Local(),
			Type:           t.Type,
			Subject:        t.Subject,
			Priority:       t.Priority,
			Status:         t.Status,
			Tag:            strings.Join(t.Tags, ","),
			RequesterID:    t.RequesterID,
			AssigneeID:     t.AssigneeID,
			OrganizationID: t.OrganizationID,
		})
	}

	err = ticket.Insert()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// user
	user := persistence.NewUser(db)

	if !user.Has() {
		if err := user.CreateTable(); err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	userRes := new(api.UsersResponse)
	byteBody, err = api.GetBody(userRes)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	parsedUser, err := api.Parse(userRes, byteBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	userRes, ok = parsedUser.(*api.UsersResponse)
	if !ok {
		return events.APIGatewayProxyResponse{}, errors.New("invalid format")
	}

	for _, u := range userRes.Users {
		user.Users = append(user.Users, domain.User{
			ID:             u.ID,
			Name:           u.Name,
			Email:          u.Email,
			CreateTime:     u.CreateTime.Local(),
			OrganizationID: u.OrganizationID,
			Alias:          u.Alias,
			Role:           u.Role,
		})
	}

	if err := user.Insert(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// organization
	org := persistence.NewOrg(db)

	if !org.Has() {
		if err := org.CreateTable(); err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
	}

	orgRes := new(api.OrgsResponse)
	byteBody, err = api.GetBody(orgRes)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	parsedOrg, err := api.Parse(orgRes, byteBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	orgRes, ok = parsedOrg.(*api.OrgsResponse)
	if !ok {
		return events.APIGatewayProxyResponse{}, errors.New("invalid format")
	}

	for _, o := range orgRes.Orgs {
		org.Organizations = append(org.Organizations, domain.Organization{
			ID:         o.ID,
			Name:       o.Name,
			CreateTime: o.CreateTime.Local(),
		})
	}

	if err := org.Insert(); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err := persistence.Exec(db, persistence.SQL); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       "ok",
	}, nil
}

func main() {
	lambda.Start(allhandler)
}
