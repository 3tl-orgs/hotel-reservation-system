package api

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyTypeService interface {
	GetPropertyTypeByIdBiz(ctx context.Context, id int) (*model.PropertyType, error)
	CreatePropertyTypeBiz(ctx context.Context, data *model.PropertyTypeCreateDto) error
	ListPropertyTypeBiz(ctx context.Context, filter *model.Filter,
		paging core.Paging,
		moreInfo ...string) ([]model.PropertyType, error)
}

type propertyTypeApi struct {
	propertyService PropertyTypeService
}

func NewPropertyTypeApi(propertyService PropertyTypeService) *propertyTypeApi {
	return &propertyTypeApi{propertyService: propertyService}
}
