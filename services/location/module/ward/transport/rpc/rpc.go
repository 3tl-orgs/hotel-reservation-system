package wardrpc

import (
	"context"

	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
)

type Business interface {
	GetWardByIdBiz(ctx context.Context, id int) (*wardsmodel.Ward, error)
}

type wardRpc struct {
	business Business
}

func NewWardRpc(business Business) *wardRpc {
	return &wardRpc{business: business}
}

func (w *wardRpc) GetWardByIdHdl(ctx context.Context, req *pb.GetWardByIdRequest) (*pb.WardResponse, error) {
	ward, err := w.business.GetWardByIdBiz(ctx, int(req.GetId()))
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(err.Error()).
			WithDebug(err.Error())
	}
	return &pb.WardResponse{Id: int32(ward.Id),
		Name:       ward.Name,
		Code:       ward.Code,
		ProvinceId: int32(ward.ProvinceId)}, nil
}
