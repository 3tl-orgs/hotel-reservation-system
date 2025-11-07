package roomtyperepo

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/pkg/errors"
)

func (r *roomTypeRepo) Create(ctx context.Context, data *roomtypemodel.RoomTypeCreateDTO) error {
	if err := r.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
