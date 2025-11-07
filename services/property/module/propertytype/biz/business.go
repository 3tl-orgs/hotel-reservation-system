package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyTypeRepo interface {
	Create(ctx context.Context, data *model.PropertyTypeCreateDto) error
	GetById(ctx context.Context, id int) (*model.PropertyType, error)
	ExistByName(ctx context.Context, name string) bool
	List(ctx context.Context, filter *model.Filter, paging core.Paging, moreInfo ...string) ([]model.PropertyType, error)
}

type business struct {
	propertyTypeRepo PropertyTypeRepo
}

func NewBusiness(propertyTypeRepo PropertyTypeRepo) *business {
	return &business{propertyTypeRepo: propertyTypeRepo}
}
