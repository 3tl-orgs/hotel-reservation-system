package roomtypebiz

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type RoomTypeRepo interface {
	Create(ctx context.Context, data *roomtypemodel.RoomTypeCreateDTO) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*roomtypemodel.RoomType, error)
	Update(ctx context.Context, id int, data *roomtypemodel.RoomTypeUpdateDTO) error
	List(ctx context.Context,
		paging *core.Paging,
		filter *roomtypemodel.Filter,
		moreKeys ...string,
	) ([]roomtypemodel.RoomType, error)
}

type roomTypeBusiness struct {
	roomTypeRepo RoomTypeRepo
}

func NewRoomTypeBusiness(roomTypeRepo RoomTypeRepo) *roomTypeBusiness {
	return &roomTypeBusiness{
		roomTypeRepo: roomTypeRepo,
	}
}
