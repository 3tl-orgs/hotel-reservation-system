package api

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
)

type PropertyTypeService interface {
	GetPropertyTypeByIdBiz(ctx context.Context, id int) (*model.PropertyType, error)
	CreatePropertyTypeBiz(ctx context.Context, data *model.PropertyTypeCreateDto) error
}

type propertyTypeApi struct {
	propertyService PropertyTypeService
}

func NewPropertyTypeApi(propertyService PropertyTypeService) *propertyTypeApi {
	return &propertyTypeApi{propertyService: propertyService}
}
