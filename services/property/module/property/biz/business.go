package propertybiz

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
)

type PropertyStore interface {
	Create(ctx context.Context, data *propertymodel.Property) error
}

type PropertyDetailStore interface {
	Create(ctx context.Context, data *propertydetailmodel.PropertyDetail) error
}

type business struct {
	propertyStore       PropertyStore
	propertyDetailStore PropertyDetailStore
	uploadProvider      uploadprovider.UploadProvider
}

func NewPropertyBiz(propertyStore PropertyStore, propertyDetailStore PropertyDetailStore,
	uploadProvider uploadprovider.UploadProvider) *business {
	return &business{propertyStore: propertyStore,
		propertyDetailStore: propertyDetailStore,
		uploadProvider:      uploadProvider}
}
