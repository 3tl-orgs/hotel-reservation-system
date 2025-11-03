package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
)

type PropertyTypeRepo interface {
	Create(ctx context.Context, data *model.PropertyTypeCreateDto) error
	GetById(ctx context.Context, id int) (*model.PropertyType, error)
	ExistByName(ctx context.Context, name string) bool
}

type business struct {
	propertyTypeRepo PropertyTypeRepo
}

func NewBusiness(propertyTypeRepo PropertyTypeRepo) *business {
	return &business{propertyTypeRepo: propertyTypeRepo}
}
