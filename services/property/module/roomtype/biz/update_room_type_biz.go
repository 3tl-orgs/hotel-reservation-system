package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

func (r *roomTypeBusiness) UpdateRoomTypeBiz(ctx context.Context, id int, data *roomtypemodel.RoomTypeUpdateDTO) error {
	existingData, err := r.GetRoomTypeByIdBiz(ctx, id)
	if err != nil {
		return roomtypemodel.ErrCannotCreateRoom
	}

	if existingData.Id != id {
		return roomtypemodel.ErrCannotCreateRoom
	}

	if err := r.roomTypeRepo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(roomtypemodel.ErrCannotUpdateRoom.Error())
	}

	updatedTime := time.Now()
	data.UpdatedAt = &updatedTime
	data.Mask()

	return nil
}
