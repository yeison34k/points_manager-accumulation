package test

import (
	"accumulation/internal/adapter"
	"accumulation/internal/domain"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDynamoDB is a mock implementation of the DynamoDBAPI interface.
type MockDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	mock.Mock
}

func (m *MockDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func (m *MockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

func TestDynamoDBRepository_GetPointByID(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)
	repo := adapter.NewDynamoDBRepository(
		"TestTableName", mockDynamoDB,
	)

	mockDynamoDB.On("GetItem", mock.AnythingOfType("*dynamodb.GetItemInput")).Return(
		&dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"id":   {S: aws.String("123")},
				"name": {S: aws.String("TestPoint")},
			},
		},
		nil,
	)

	result, err := repo.GetPointByID("123")

	mockDynamoDB.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, "123", result.ID)
	assert.Equal(t, "TestUser", result.User)
}

func TestDynamoDBRepository_CreatePoint(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)

	repo := &adapter.DynamoDBRepository{
		DynamoDBClient: mockDynamoDB,
		TableName:      "TestTableName",
	}

	mockDynamoDB.On("PutItem", mock.AnythingOfType("*dynamodb.PutItemInput")).Return(
		&dynamodb.PutItemOutput{},
		nil,
	)

	err := repo.CreatePoint(&domain.Point{
		ID:    "123",
		User:  "TestUser",
		Total: 123.45,
	})

	mockDynamoDB.AssertExpectations(t)
	assert.NoError(t, err)
}
