package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (r *roomTypeBusiness) GetRoomTypeByIdBiz(ctx context.Context, id int) (*roomtypemodel.RoomType, error) {
	data, err := r.roomTypeRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrNotFound) {
			return nil, core.ErrNotFound.
				WithError(roomtypemodel.ErrCannotGetRoom.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(roomtypemodel.ErrCannotGetRoom.Error())
	}
	
	data.Mask()
	return data, nil
}
