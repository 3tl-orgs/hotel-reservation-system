package facilitypropertiesrepo

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/pkg/errors"
)

func (fp *facilityPropertiesRepo) List(ctx context.Context) ([]facilitypropertiesmodel.FacilityProperties, error) {
	var result []facilitypropertiesmodel.FacilityProperties
	if err := fp.db.Table(facilitypropertiesmodel.FacilityProperties{}.TableName()).
		Where("status is true").
		Find(&result).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return result, nil
}
