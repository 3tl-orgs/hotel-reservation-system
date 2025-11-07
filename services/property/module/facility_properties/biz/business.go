package facilitypropertiesbiz

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
)

type facilityPropertiesRepo interface {
	Create(ctx context.Context, data *facilitypropertiesmodel.FacilityPropertiesCreateDTO) error
	Update(ctx context.Context, id int, data *facilitypropertiesmodel.FacilityPropertiesUpdateDTO) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*facilitypropertiesmodel.FacilityProperties, error)
	List(ctx context.Context) ([]facilitypropertiesmodel.FacilityProperties, error)
	GetFacility(ctx context.Context, id int) ([]facilitypropertiesmodel.FacilityResponse, error)
}

type facilityPropertiesBiz struct {
	business facilityPropertiesRepo
}

func NewFacilityPropertiesRepo(business facilityPropertiesRepo) *facilityPropertiesBiz {
	return &facilityPropertiesBiz{
		business: business,
	}
}
