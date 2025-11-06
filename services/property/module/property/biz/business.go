package propertybiz

import (
	"context"

	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
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
}

func NewPropertyBiz(propertyStore PropertyStore) *business {
	return &business{propertyStore: propertyStore}
}
