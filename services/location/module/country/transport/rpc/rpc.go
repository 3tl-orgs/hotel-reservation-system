package rpc

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
)

type Business interface {
	GetCountryByIdBiz(ctx context.Context, id int) (*model.Country, error)
}

type countryRpc struct {
	business Business
}

func NewRpc(business Business) *countryRpc {
	return &countryRpc{business: business}
}

func (rpc *countryRpc) GetCountryByIdHdl(ctx context.Context, req *pb.GetCountryByIdRequest) (*pb.CountryResponse, error) {
	country, err := rpc.business.GetCountryByIdBiz(ctx, int(req.Id))
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(err.Error()).
			WithDebug(err.Error())
	}
	return &pb.CountryResponse{
		Code: country.Code,
		Id:   int32(country.Id),
		Name: country.Name,
	}, nil
}
