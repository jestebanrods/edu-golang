package internal

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/graph-gophers/graphql-go"
)

var ErrQueryNameNotProvided = errors.New("no query was provided in the HTTP body")

type graphqlBody struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

type Handler struct {
	schema *graphql.Schema
}

func (h *Handler) Handle(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("processing lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrQueryNameNotProvided
	}

	body := graphqlBody{}

	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Print("could not decode body", err)
	}

	result := h.schema.Exec(context, body.Query, body.OperationName, body.Variables)

	response, err := json.Marshal(result)
	if err != nil {
		log.Print("could not decode body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func NewHandler(schema *graphql.Schema) *Handler {
	return &Handler{
		schema: schema,
	}
}
