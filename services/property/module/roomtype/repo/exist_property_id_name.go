package roomtyperepo

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *roomTypeRepo) ExistByPropertyAndName(ctx context.Context, propertyId int, name string) bool {
	var roomtype roomtypemodel.RoomType
	if err := r.db.WithContext(ctx).
		Where("property_id = ? AND name = ?", propertyId, name).
		First(&roomtype).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}
