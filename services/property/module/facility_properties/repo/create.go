package facilitypropertiesrepo

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/pkg/errors"
)

func (fp *facilityPropertiesRepo) Create(ctx context.Context, data *facilitypropertiesmodel.FacilityPropertiesCreateDTO) error {
	if err := fp.db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
