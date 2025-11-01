package common

type GrpcConfig interface {
	GetGRPCPort() int
	GetGRPCServerAddress() string
	GetGRPCLocationServerAddress() string
}
