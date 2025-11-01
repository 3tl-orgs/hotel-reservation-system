package client

import (
	"log"

	"github.com/ngleanhvu/go-booking/services/property/common"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	"github.com/ngleanhvu/go-booking/shared/srvctx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ComposerCountryRpcClient(sctx srvctx.ServiceContext) pb.CountryServiceClient {
	grpcConfig := sctx.MustGet(core.KeyCompGrpcConf).(common.GrpcConfig)
	clientConn, err := grpc.Dial(
		grpcConfig.GetGRPCLocationServerAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewCountryServiceClient(clientConn)
}
