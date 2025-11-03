package facilitypropertiesrepo

import (
	"context"
	facilitypropertiesmodel "github.com/ngleanhvu/go-booking/services/property/module/facility_properties/model"
	"github.com/pkg/errors"
)

func (fp *facilityPropertiesRepo) Update(ctx context.Context, id int, data *facilitypropertiesmodel.FacilityPropertiesUpdateDTO) error {
	update := map[string]interface{}{}

	//if data.AmenityIds != nil {
	//	jSon, _ := json.Marshal(data.AmenityIds)
	//	update["amenity_ids"] = datatypes.JSON(jSon)
	//}
	if data.AmenityIds != nil {
		update["amenity_ids"] = *data.AmenityIds
	}

	if data.Status != nil {
		update["status"] = *data.Status
	}

	if len(update) > 0 {
		if err := fp.db.WithContext(ctx).
			Table(data.TableName()).
			Where("id = ?", id).
			Updates(&update).Error; err != nil {
			return errors.WithStack(err)
		}

	}

	return nil

}
