package propertytypeapi

import (
	"context"

	propertytypemodel "github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyTypeService interface {
	GetPropertyTypeByIdBiz(ctx context.Context, id int) (*propertytypemodel.PropertyType, error)
	CreatePropertyTypeBiz(ctx context.Context, data *propertytypemodel.PropertyTypeCreateDto) error
	ListPropertyTypeBiz(ctx context.Context, filter *propertytypemodel.Filter,
		paging core.Paging,
		moreInfo ...string) ([]propertytypemodel.PropertyType, error)
}

type propertyTypeApi struct {
	propertyService PropertyTypeService
}

func NewPropertyTypeApi(propertyService PropertyTypeService) *propertyTypeApi {
	return &propertyTypeApi{propertyService: propertyService}
}
