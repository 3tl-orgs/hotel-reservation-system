package facilitypropertiesbiz

import (
	"context"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (fp *facilityPropertiesBiz) ListFacilityPropBiz(ctx context.Context) ([]facilitypropertiesmodel.FacilityProperties, error) {
	data, err := fp.business.List(ctx)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotGetCountry.Error()).
			WithDebug(err.Error())
	}

	for i := range data {
		data[i].Mask()
	}
	return data, nil
}
