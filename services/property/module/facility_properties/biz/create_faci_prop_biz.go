package facilitypropertiesbiz

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (fp *facilityPropertiesBiz) CreateFacilityPropBiz(ctx context.Context, data *facilitypropertiesmodel.FacilityPropertiesCreateDTO) error {

	if err := fp.business.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(facilitypropertiesmodel.ErrCannotCreateFacilityProperties.Error()).
			WithDebug(err.Error())
	}
	return nil
}
