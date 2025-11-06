package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (r *roomTypeBusiness) DeleteRoomTypeBiz(ctx context.Context, id int) error {
	existingData, err := r.GetRoomTypeByIdBiz(ctx, id)
	if err != nil {
		return roomtypemodel.ErrRoomNotFound
	}

	if existingData.Status == core.EntityDeleted {
		return core.ErrBadRequest.
			WithError(roomtypemodel.ErrRoomHasBeenDeleted.Error())
	}

	if err := r.roomTypeRepo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(roomtypemodel.ErrCannotDeleteRoom.Error())
	}

	return nil
}
