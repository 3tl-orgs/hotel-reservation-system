package facilitypropertiesrepo

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (fp *facilityPropertiesRepo) GetById(ctx context.Context, id int) (*facilitypropertiesmodel.FacilityProperties, error) {
	var data facilitypropertiesmodel.FacilityProperties

	if err := fp.db.Table(data.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, facilitypropertiesmodel.ErrCannotGetFacilityProperties
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
