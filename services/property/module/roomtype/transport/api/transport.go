package roomtypeapi

import (
	"context"
	roomtypemodel "github.com/ngleanhvu/go-booking/services/property/module/roomtype/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type roomTypeBusiness interface {
	UpdateRoomTypeBiz(ctx context.Context, id int, data *roomtypemodel.RoomTypeUpdateDTO) error
	CreateRoomTypeBiz(ctx context.Context, data *roomtypemodel.RoomTypeCreateDTO) error
	DeleteRoomTypeBiz(ctx context.Context, id int) error
	GetRoomTypeByIdBiz(ctx context.Context, id int) (*roomtypemodel.RoomType, error)
	ListRoomTypeBiz(ctx context.Context,
		paging *core.Paging,
		filter *roomtypemodel.Filter,
		moreKeys ...string,
	) ([]roomtypemodel.RoomType, error)
	GetRoomTypeByProp(ctx context.Context, propertyId int) ([]roomtypemodel.RoomType, error)
}

type RoomTypeTransport struct {
	roomTypeBusiness roomTypeBusiness
}

func NewRoomTypeTransport(roomTypeBusiness roomTypeBusiness) *RoomTypeTransport {
	return &RoomTypeTransport{
		roomTypeBusiness: roomTypeBusiness,
	}
}
