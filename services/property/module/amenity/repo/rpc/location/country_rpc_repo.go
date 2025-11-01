package location

import (
	"context"

	"github.com/ngleanhvu/go-booking/shared/dto"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	"github.com/pkg/errors"
)

type rpcClient struct {
	client pb.CountryServiceClient
}

func NewCountryClient(client pb.CountryServiceClient) *rpcClient {
	return &rpcClient{client: client}
}

func (c *rpcClient) GetCountryById(ctx context.Context, id int32) (*dto.CountryResponse, error) {
	country, err := c.client.GetCountryByIdHdl(ctx, &pb.GetCountryByIdRequest{Id: id})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := &dto.CountryResponse{int(country.Id),
		country.Code,
		country.Name,
	}

	return res, nil
}
