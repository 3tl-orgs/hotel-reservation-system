package propertyapi

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
)

type Business interface {
	CreatePropertyBiz(ctx context.Context, req *propertymodel.PropertyCreateDTO) error
}

type propertyApi struct {
	business Business
	upload   uploadprovider.UploadProvider
}

func NewPropertyApi(business Business, upload uploadprovider.UploadProvider) *propertyApi {
	return &propertyApi{business: business, upload: upload}
}
