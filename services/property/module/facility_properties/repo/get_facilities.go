package facilitypropertiesrepo

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (fp *facilityPropertiesRepo) GetFacility(ctx context.Context, id int) ([]facilitypropertiesmodel.FacilityResponse, error) {
	var data []facilitypropertiesmodel.FacilityResponse
	fp.db = fp.db.Debug()

	if err := fp.db.Table(facilitypropertiesmodel.FacilityProperties{}.TableName()).
		Where("property_id = ?", id).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, facilitypropertiesmodel.ErrCannotGetFacilityProperties
		}
		return nil, errors.WithStack(err)
	}

	return data, nil
}
