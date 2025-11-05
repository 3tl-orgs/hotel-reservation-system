package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *roomTypeBusiness) GetRoomTypeByProp(ctx context.Context, propertyId int) ([]roomtypemodel.RoomType, error) {
	data, err := r.roomTypeRepo.GetRoomTypes(ctx, propertyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, core.ErrNotFound.
				WithError(roomtypemodel.ErrYourPropertyHaveNoRoomTypes.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(err.Error())
	}

	for i := range data {
		data[i].Mask()
	}

	return data, nil
}
