package facilitypropertiesbiz

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (fp *facilityPropertiesBiz) GetFacilityByPropBiz(ctx context.Context, id int) ([]facilitypropertiesmodel.FacilityResponse, error) {
	data, err := fp.business.GetFacility(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrNotFound.
				WithError(facilitypropertiesmodel.ErrCannotGetFacilityProperties.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(facilitypropertiesmodel.ErrCannotGetFacilityProperties.Error())
	}

	return data, nil
}
