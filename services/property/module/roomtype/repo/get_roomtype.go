package roomtyperepo

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *roomTypeRepo) GetRoomTypes(ctx context.Context, propertyId int) ([]roomtypemodel.RoomType, error) {
	var data []roomtypemodel.RoomType

	if err := r.db.Table(roomtypemodel.RoomType{}.TableName()).
		Where("property_id = ?", propertyId).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, roomtypemodel.ErrYourPropertyHaveNoRoomTypes
		}
		return nil, errors.WithStack(err)
	}

	return data, nil
}
