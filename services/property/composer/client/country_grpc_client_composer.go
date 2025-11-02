package client

import (
	"context"
	"log"

	"github.com/ngleanhvu/go-booking/services/property/common"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/dto"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CountryRPCClient interface {
	GetCountryById(ctx context.Context, id int32) (*dto.CountryResponse, error)
}

type rpcClient struct {
	client pb.CountryServiceClient
}

func (c *rpcClient) GetCountryById(ctx context.Context, id int32) (*dto.CountryResponse, error) {
	country, err := c.client.GetCountryByIdHdl(ctx, &pb.GetCountryByIdRequest{Id: id})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &dto.CountryResponse{
		Id:   int(country.Id),
		Code: country.Code,
		Name: country.Name,
	}, nil
}

func ComposerCountryRpcClient(sctx srvctx.ServiceContext) CountryRPCClient {
	grpcConfig := sctx.MustGet(core.KeyCompGrpcConf).(common.GrpcConfig)
	clientConn, err := grpc.Dial(
		grpcConfig.GetGRPCLocationServerAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewCountryServiceClient(clientConn)
	return &rpcClient{client: client}
}
