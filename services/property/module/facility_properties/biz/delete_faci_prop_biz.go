package facilitypropertiesbiz

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (fp *facilityPropertiesBiz) DeleteFacilityPropBiz(ctx context.Context, id int) error {
	existingData, err := fp.GetByIdFacilityPropBiz(ctx, id)
	if err != nil {
		return facilitypropertiesmodel.ErrFacilityPropertiesNotFound
	}

	if existingData.Status == core.EntityDeleted {
		return facilitypropertiesmodel.ErrRecordHasBeenDeleted
	}

	if err := fp.business.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(facilitypropertiesmodel.ErrCannotDeleteFacilityProperties.Error()).
			WithDebug(err.Error())
	}
	return nil
}
