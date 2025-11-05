package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (r *roomTypeBusiness) CreateRoomTypeBiz(ctx context.Context, data *roomtypemodel.RoomTypeCreateDTO) error {
	valid := r.roomTypeRepo.ExistByPropertyAndName(ctx, data.PropertyId, data.Name)

	if valid == true {
		return core.ErrConflict.WithError(roomtypemodel.ErrRoomIsDuplicated.Error())
	}

	if err := r.roomTypeRepo.Create(ctx, data); err != nil {
		return errors.WithStack(core.ErrInternalServerError.
			WithError(roomtypemodel.ErrCannotGetRoom.Error()))
	}

	return nil
}
