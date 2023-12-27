package domain

type PointRepository interface {
	GetPointByID(id string) (*Point, error)
	CreatePoint(point *Point) error
}
