package adapter

import (
	"accumulation/internal/domain"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBRepository struct {
	DynamoDBClient DynamoDBAPI
	TableName      string
}

type DynamoDBAPI interface {
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func NewDynamoDBRepository(tableName string, client DynamoDBAPI) *DynamoDBRepository {
	return &DynamoDBRepository{
		DynamoDBClient: client,
		TableName:      tableName,
	}
}

func (r *DynamoDBRepository) GetPointByID(id string) (*domain.Point, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	result, err := r.DynamoDBClient.GetItem(input)
	if err != nil {
		return nil, err
	}

	point := &domain.Point{
		ID:   id,
		Name: aws.StringValue(result.Item["name"].S),
	}

	return point, nil
}

func (r *DynamoDBRepository) CreatePoint(point *domain.Point) error {
	totalAsString := fmt.Sprintf("%.2f", point.Total)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(point.ID),
			},
			"user": {
				S: aws.String(point.User),
			},
			"name": {
				S: aws.String(point.Name),
			},
			"total": {
				N: aws.String(totalAsString),
			},
		},
	}

	_, err := r.DynamoDBClient.PutItem(input)
	return err
}
