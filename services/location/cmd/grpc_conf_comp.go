package cmd

import (
	"flag"

	"github.com/ngleanhvu/go-booking/shared/core"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
)

type config struct {
	grpcPort          int    // for server port listening
	grpcServerAddress string // for client make grpc client connection
}

func NewGrpcConfig() *config {
	return &config{}
}

func (c *config) ID() string {
	return core.KeyCompGrpcConf
}

func (c *config) InitFlags() {
	flag.IntVar(
		&c.grpcPort,
		"grpc-port",
		3101,
		"gRPC Port. Default: 3101",
	)
	flag.StringVar(&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3101",
		"gRPC server address")
}

func (c *config) Activate(_ sctx.ServiceContext) error {
	return nil
}

func (c *config) Stop() error {
	return nil
}

func (c *config) GetGRPCPort() int {
	return c.grpcPort
}

func (c *config) GetGRPCServerAddress() string {
	return c.grpcServerAddress
}
