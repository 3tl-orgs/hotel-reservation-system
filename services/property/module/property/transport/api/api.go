package propertyapi

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
)

type Business interface {
	CreatePropertyBiz(ctx context.Context, req *propertymodel.PropertyCreateDTO) error
}

type propertyApi struct {
	business Business
}

func NewPropertyApi(business Business) *propertyApi {
	return &propertyApi{business: business}
}
