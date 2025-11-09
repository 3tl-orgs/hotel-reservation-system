package propertytypebiz

import (
	"context"

	propertytypemodel "github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyTypeRepo interface {
	Create(ctx context.Context, data *propertytypemodel.PropertyTypeCreateDto) error
	GetById(ctx context.Context, id int) (*propertytypemodel.PropertyType, error)
	ExistByName(ctx context.Context, name string) bool
	List(ctx context.Context, filter *propertytypemodel.Filter, paging core.Paging, moreInfo ...string) ([]propertytypemodel.PropertyType, error)
}

type business struct {
	propertyTypeRepo PropertyTypeRepo
}

func NewBusiness(propertyTypeRepo PropertyTypeRepo) *business {
	return &business{propertyTypeRepo: propertyTypeRepo}
}
