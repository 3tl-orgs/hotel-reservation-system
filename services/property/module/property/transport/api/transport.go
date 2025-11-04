package propertyapi

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type PropertyBusiness interface {
	CreatePropertyBiz(ctx context.Context, data *propertymodel.PropertyCreateDTO) error
	DeletePropertyBiz(ctx context.Context, id int) error
	UpdatePropertyBiz(ctx context.Context, id int, data *propertymodel.PropertyUpdateDTO) error
	GetPropertyByIdBiz(ctx context.Context, id int) (*propertymodel.Property, error)
	ListPropertyBiz(ctx context.Context,
		pagingData *core.Paging,
		filter *propertymodel.Filter,
		moreKeys ...string,
	) ([]propertymodel.Property, error)
}

type propertyTransport struct {
	propertyBusiness PropertyBusiness
}

func NewPropertyTransport(propertyBusiness PropertyBusiness) *propertyTransport {
	return &propertyTransport{
		propertyBusiness: propertyBusiness,
	}
}
