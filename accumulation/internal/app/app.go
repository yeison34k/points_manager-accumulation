package app

import (
	"accumulation/internal/domain"
	"accumulation/internal/usecase"
)

type MyApp struct {
	pointUsecase usecase.PointUsecase
}

func NewMyApp(pointUsecase usecase.PointUsecase) *MyApp {
	return &MyApp{
		pointUsecase: pointUsecase,
	}
}



func (a *MyApp) HandleRequest(req *domain.Point) error {
	err := a.pointUsecase.CreatePoint(req)
	if err != nil {
		return err
	}

	return nil
}
