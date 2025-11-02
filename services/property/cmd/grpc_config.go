package cmd

import (
	"flag"

	"github.com/ngleanhvu/go-booking/shared/core"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
)

type config struct {
	grpcPort                  int
	grpcServerAddress         string
	grpcLocationServerAddress string
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
		3102,
		"gRPC Port. Default: 3102",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3102",
		"gRPC Address. localhost:3102",
	)
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

func (c *config) GetGRPCLocationServerAddress() string {
	return c.grpcLocationServerAddress
}

func (c *config) GetGRPCServerAddress() string {
	return c.grpcServerAddress
}
