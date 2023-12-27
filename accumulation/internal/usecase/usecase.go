package usecase

import (
	"accumulation/internal/domain"
)

type PointUsecase struct {
	pointRepository domain.PointRepository
}

func NewPointUsecase(pointRepository domain.PointRepository) *PointUsecase {
	return &PointUsecase{
		pointRepository: pointRepository,
	}
}

func (u *PointUsecase) GetPointByID(id string) (*domain.Point, error) {
	return u.pointRepository.GetPointByID(id)
}

func (u *PointUsecase) CreatePoint(point *domain.Point) error {
	return u.pointRepository.CreatePoint(point)
}
