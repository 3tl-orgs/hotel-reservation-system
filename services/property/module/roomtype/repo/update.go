package roomtyperepo

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/pkg/errors"
)

func (r *roomTypeRepo) Update(ctx context.Context, id int, data *roomtypemodel.RoomTypeUpdateDTO) error {
	update := data.PatchToModel()

	if len(update) <= 0 {
		return nil
	}

	if err := r.db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(update).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
