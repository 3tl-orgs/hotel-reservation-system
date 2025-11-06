package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (r *roomTypeBusiness) ListRoomTypeBiz(ctx context.Context,
	paging *core.Paging,
	filter *roomtypemodel.Filter,
	moreKeys ...string,
) ([]roomtypemodel.RoomType, error) {
	data, err := r.roomTypeRepo.List(ctx, paging, filter, moreKeys...)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(roomtypemodel.ErrCannotGetRoom.Error()).
			WithDebug(err.Error())
	}

	for i := range data {
		data[i].Mask()
	}

	return data, nil
}
