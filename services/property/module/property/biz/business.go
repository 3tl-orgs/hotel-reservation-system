package propertybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	propertydetailmodel "github.com/ngleanhvu/go-booking/services/property/module/propertydetail/model"
	propertytypemodel "github.com/ngleanhvu/go-booking/services/property/module/propertytype/model"
	"github.com/ngleanhvu/go-booking/shared/dto"
	"github.com/ngleanhvu/go-booking/shared/uploadprovider"
)

type PropertyStore interface {
	Create(ctx context.Context, data *propertymodel.Property, data1 *propertydetailmodel.PropertyDetail) error
}

type PropertyTypeStore interface {
	GetById(ctx context.Context, id int) (*propertytypemodel.PropertyType, error)
}

type CountryClient interface {
	GetCountryById(ctx context.Context, id int32) (*dto.CountryResponse, error)
}

type ProvinceClient interface {
	GetProvinceById(ctx context.Context, id int32) (*dto.ProvinceResponse, error)
}

type WardClient interface {
	GetWardById(ctx context.Context, id int32) (*dto.WardResponse, error)
}

type AmenityStore interface {
	GetByIds(ctx context.Context, ids []int) ([]amenitymodel.Amenity, error)
}

type business struct {
	propertyStore     PropertyStore
	uploadProvider    uploadprovider.UploadProvider
	propertyTypeStore PropertyTypeStore
	amenityStore      AmenityStore
	countryClient     CountryClient
	provinceClient    ProvinceClient
	wardClient        WardClient
}

func NewPropertyBiz(propertyStore PropertyStore,
	uploadProvider uploadprovider.UploadProvider,
	propertyTypeStore PropertyTypeStore,
	amenityStore AmenityStore,
	countryClient CountryClient,
	provinceClient ProvinceClient,
	wardClient WardClient,
) *business {
	return &business{propertyStore: propertyStore,
		uploadProvider:    uploadProvider,
		propertyTypeStore: propertyTypeStore,
		amenityStore:      amenityStore,
		countryClient:     countryClient,
		provinceClient:    provinceClient,
		wardClient:        wardClient,
	}
}
