package roomtyperepo

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *roomTypeRepo) GetById(ctx context.Context, id int) (*roomtypemodel.RoomType, error) {
	var data roomtypemodel.RoomType

	if err := r.db.Table(roomtypemodel.RoomType{}.TableName()).
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(roomtypemodel.ErrRoomNotFound)
		}
		return nil, errors.WithStack(err)
	}
	return &data, nil
}
