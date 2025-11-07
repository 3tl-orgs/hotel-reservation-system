package facilitypropertiesrepo

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/pkg/errors"
	"time"
)

func (fp *facilityPropertiesRepo) Delete(ctx context.Context, id int) error {

	if err := fp.db.Table(facilitypropertiesmodel.FacilityProperties{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     false,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
		
}
