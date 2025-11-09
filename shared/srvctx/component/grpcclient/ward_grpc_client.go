package grpcclient

import (
	"flag"
	"log"

	"context"

	"github.com/ngleanhvu/go-booking/shared/dto"
	"github.com/ngleanhvu/go-booking/shared/proto/pb"
	sctx "github.com/ngleanhvu/go-booking/shared/srvctx"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WardRPCComponent struct {
	id       string
	prefix   string
	addr     string
	client   pb.WardServiceClient
	conn     *grpc.ClientConn
	logger   sctx.Logger
	isActive bool
}

// NewCountryRPCClientComponent khởi tạo component gRPC client cho service Country
func NewWardRPCClientComponent(id, prefix string) *WardRPCComponent {
	return &WardRPCComponent{
		id:     id,
		prefix: prefix,
	}
}

// ID trả về định danh component (phải duy nhất)
func (c *WardRPCComponent) ID() string {
	return c.id
}

// InitFlags khai báo cờ cấu hình cho địa chỉ gRPC server
func (c *WardRPCComponent) InitFlags() {
	if flag.Lookup("grpc-ward-server-address") == nil {
		flag.StringVar(
			&c.addr,
			"grpc-ward-server-address",
			"localhost:3101",
			"ward gRPC server address. Default: localhost:3101",
		)
	}
}

// Activate khởi tạo kết nối gRPC
func (c *WardRPCComponent) Activate(_ sctx.ServiceContext) error {
	c.logger = sctx.GlobalLogger().GetLogger(c.id)
	c.logger.Infof("Connecting to Country gRPC server at %s...", c.addr)

	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "failed to connect to Country gRPC server")
	}

	c.conn = conn
	c.client = pb.NewWardServiceClient(conn)
	c.isActive = true
	c.logger.Info("Country gRPC client activated successfully.")

	return nil
}

// Stop đóng kết nối khi service dừng
func (c *WardRPCComponent) Stop() error {
	if c.conn != nil {
		c.logger.Info("Closing Country gRPC connection...")
		if err := c.conn.Close(); err != nil {
			return errors.Wrap(err, "failed to close gRPC connection")
		}
	}
	c.isActive = false
	return nil
}

// GetClient trả về pb.CountryServiceClient đã được khởi tạo
func (c *WardRPCComponent) GetClient() pb.WardServiceClient {
	if !c.isActive {
		log.Fatal("CountryRPCComponent is not active")
	}
	return c.client
}

// GetCountryById ví dụ hàm tiện ích gọi gRPC
func (c *WardRPCComponent) GetWardById(ctx context.Context, id int32) (*dto.WardResponse, error) {
	resp, err := c.client.GetWardByIdHdl(ctx, &pb.GetWardByIdRequest{Id: id})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &dto.WardResponse{
		Id:         int(resp.Id),
		Code:       resp.Code,
		Name:       resp.Name,
		ProvinceId: int(resp.ProvinceId),
	}, nil
}
