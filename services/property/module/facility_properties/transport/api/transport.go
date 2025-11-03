package facilitypropertiesapi

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
)

type facilityPropertiesBiz interface {
	CreateFacilityPropBiz(ctx context.Context, data *facilitypropertiesmodel.FacilityPropertiesCreateDTO) error
	UpdateFacilityPropBiz(ctx context.Context, id int, data *facilitypropertiesmodel.FacilityPropertiesUpdateDTO) error
	DeleteFacilityPropBiz(ctx context.Context, id int) error
	GetByIdFacilityPropBiz(ctx context.Context, id int) (*facilitypropertiesmodel.FacilityProperties, error)
	ListFacilityPropBiz(ctx context.Context) ([]facilitypropertiesmodel.FacilityProperties, error)
	GetFacilityByPropBiz(ctx context.Context, id int) ([]facilitypropertiesmodel.FacilityResponse, error)
}

type FacilityPropertiesTransport struct {
	transport facilityPropertiesBiz
}

func NewFacilityPropertiesTransport(transport facilityPropertiesBiz) *FacilityPropertiesTransport {
	return &FacilityPropertiesTransport{
		transport: transport,
	}
}
