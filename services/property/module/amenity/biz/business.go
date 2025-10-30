package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
)

type AmenityRepo interface {
	Create(ctx context.Context, data *model.AmenityCreateDto) error
	GetByIds(ctx context.Context, ids []int) ([]model.Amenity, error)
	GetById(ctx context.Context, id int) (*model.Amenity, error)
	Update(ctx context.Context, id int, data *model.AmenityUpdateDto) error
	Delete(ctx context.Context, id int) error
	DeleteMany(ctx context.Context, ids []int) error
}

type business struct {
	repo AmenityRepo
}

func NewBusiness(repo AmenityRepo) *business {
	return &business{repo: repo}
}
