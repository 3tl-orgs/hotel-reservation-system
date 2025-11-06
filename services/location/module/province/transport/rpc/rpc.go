package provincerpc

import (
	"context"

	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
)

type Business interface {
	GetProvinceByIdBiz(ctx context.Context, id int) (*provincemodel.Province, error)
}

type provinceRpc struct {
	business Business
}

func NewProvinceRpc(business Business) *provinceRpc {
	return &provinceRpc{business: business}
}

func (rpc *provinceRpc) GetProvinceByIdHdl(ctx context.Context, req *pb.GetProvinceByIdRequest) (*pb.ProvinceResponse, error) {
	province, err := rpc.business.GetProvinceByIdBiz(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(err.Error()).
			WithDebug(err.Error())
	}
	return &pb.ProvinceResponse{Id: int32(province.Id),
		Name:      province.Name,
		Code:      province.Code,
		CountryId: int32(province.CountryId)}, nil
}
