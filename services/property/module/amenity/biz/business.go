package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
)

type AmenityRepo interface {
	Create(ctx context.Context, data *model.AmenityCreateDto) error
	GetByIds(ctx context.Context, ids []int) ([]model.Amenity, error)
}

type business struct {
	repo AmenityRepo
}

func NewBusiness(repo AmenityRepo) *business {
	return &business{repo: repo}
}
