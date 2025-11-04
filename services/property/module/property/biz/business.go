package propertybiz

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyRepo interface {
	Create(ctx context.Context, data *propertymodel.PropertyCreateDTO) error
	Update(ctx context.Context, id int, data *propertymodel.PropertyUpdateDTO) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*propertymodel.Property, error)
	List(ctx context.Context,
		pagingData *core.Paging,
		filter *propertymodel.Filter,
		moreKeys ...string,
	) ([]propertymodel.Property, error)
}

type propertyBusiness struct {
	propertyRepo PropertyRepo
}

func NewPropertyBusiness(propertyRepo PropertyRepo) *propertyBusiness {
	return &propertyBusiness{
		propertyRepo: propertyRepo,
	}
}
