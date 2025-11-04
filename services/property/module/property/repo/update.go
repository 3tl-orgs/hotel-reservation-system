package propertyrepo

import (
	"context"
	propertymodel "github.com/ngleanhvu/go-booking/services/property/module/property/model"
	"github.com/pkg/errors"
)

func (p *propertyRepo) Update(ctx context.Context, id int, data *propertymodel.PropertyUpdateDTO) error {
	update := map[string]interface{}{}

	if data.Name != nil {
		update["name"] = *data.Name
	}

	if data.Address != nil {
		update["address"] = *data.Address
	}

	if data.ProvinceId != nil {
		update["province_id"] = *data.ProvinceId
	}

	if data.WardId != nil {
		update["ward_id"] = *data.WardId
	}

	if len(update) > 0 {
		if err := p.db.WithContext(ctx).
			Table(data.TableName()).
			Where("id = ?", id).
			Updates(&update).Error; err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
