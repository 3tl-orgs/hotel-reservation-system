package roomtyperepo

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/pkg/errors"
	"time"
)

func (r *roomTypeRepo) Delete(ctx context.Context, id int) error {
	if err := r.db.Table(roomtypemodel.RoomType{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     false,
			"updated_at": time.Now(),
		}).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
