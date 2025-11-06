package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/ngleanhvu/go-booking/services/location/composer"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
	"google.golang.org/grpc"
)

func StartGRPCServices(serviceCtx sctx.ServiceContext) {
	configComp := serviceCtx.MustGet(core.KeyCompGrpcConf).(core.GrpcConfig)
	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))

	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())

	s := grpc.NewServer()

	pb.RegisterCountryServiceServer(s, composer.NewCountryGrpcTransport(serviceCtx))
	pb.RegisterProvinceServiceServer(s, composer.NewProvinceGrpcTransport(serviceCtx))
	pb.RegisterWardServiceServer(s, composer.NewWardGrpcTransport(serviceCtx))

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
