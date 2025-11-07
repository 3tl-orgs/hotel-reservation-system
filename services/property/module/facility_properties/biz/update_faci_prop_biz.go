package facilitypropertiesbiz

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

func (fp *facilityPropertiesBiz) UpdateFacilityPropBiz(ctx context.Context, id int, data *facilitypropertiesmodel.FacilityPropertiesUpdateDTO) error {
	existingData, err := fp.GetByIdFacilityPropBiz(ctx, id)
	if err != nil {
		return facilitypropertiesmodel.ErrFacilityPropertiesNotFound
	}

	if existingData.Id != id {
		return facilitypropertiesmodel.ErrFacilityPropertiesNotFound
	}

	if err := fp.business.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(facilitypropertiesmodel.ErrCannotUpdateFacilityProperties.Error()).
			WithDebug(err.Error())
	}

	updatedTime := time.Now()
	data.UpdatedAt = &updatedTime
	return nil

}
