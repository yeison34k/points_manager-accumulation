package main

import (
	"accumulation/internal/adapter"
	"accumulation/internal/app"
	"accumulation/internal/domain"
	"accumulation/internal/usecase"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LambdaHandler struct {
	myApp app.MyApp
}

func NewLambdaHandler(pointRepository domain.PointRepository) *LambdaHandler {
	pointUsecase := usecase.NewPointUsecase(pointRepository)
	myApp := app.NewMyApp(*pointUsecase)
	return &LambdaHandler{
		myApp: *myApp,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := []byte(request.Body)
	var body domain.Point
	err := json.Unmarshal(b, &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	date := time.Now()
	layout := "01-02-2006 15:04:05"
	dateString := date.Format(layout)
	dateTime, err := time.Parse(layout, dateString)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error parsing date-time",
		}, nil
	}

	body.CreateDate = dateTime.String()

	err = h.myApp.HandleRequest(&body)
	if err != nil {
		log.Fatal(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "point: success create",
	}, nil
}

func main() {
	tableName := "Point"
	sess := session.Must(session.NewSession())
	actualDynamoDBClient := dynamodb.New(sess)
	pointRepository := adapter.NewDynamoDBRepository(tableName, actualDynamoDBClient)
	handler := NewLambdaHandler(pointRepository)
	lambda.Start(handler.HandleRequest)
}
