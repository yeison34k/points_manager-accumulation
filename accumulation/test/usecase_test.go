package test

import (
	"accumulation/internal/domain"
	"accumulation/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockPointRepository struct {
}

func (m *MockPointRepository) GetPointByID(id string) (*domain.Point, error) {
	return &domain.Point{ID: "1", User: "Compra TC"}, nil
}

func (m *MockPointRepository) CreatePoint(point *domain.Point) error {
	return nil
}

func TestPointUsecase_GetPointByID(t *testing.T) {
	mockRepository := &MockPointRepository{}
	usecase := usecase.NewPointUsecase(mockRepository)

	point, err := usecase.GetPointByID("1")

	assert.NoError(t, err)
	assert.NotNil(t, point)
	assert.Equal(t, "Compra TC", point.User)
}

func TestPointUsecase_CreatePoint(t *testing.T) {
	mockRepository := &MockPointRepository{}
	usecase := usecase.NewPointUsecase(mockRepository)

	err := usecase.CreatePoint(&domain.Point{ID: "1", User: "Jane Doe"})

	assert.NoError(t, err)
}
