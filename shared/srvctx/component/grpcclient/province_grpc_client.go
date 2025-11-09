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

type ProvinceRPCComponent struct {
	id       string
	prefix   string
	addr     string
	client   pb.ProvinceServiceClient
	conn     *grpc.ClientConn
	logger   sctx.Logger
	isActive bool
}

// NewCountryRPCClientComponent khởi tạo component gRPC client cho service Country
func NewProvinceRPCClientComponent(id, prefix string) *ProvinceRPCComponent {
	return &ProvinceRPCComponent{
		id:     id,
		prefix: prefix,
	}
}

// ID trả về định danh component (phải duy nhất)
func (c *ProvinceRPCComponent) ID() string {
	return c.id
}

// InitFlags khai báo cờ cấu hình cho địa chỉ gRPC server
func (c *ProvinceRPCComponent) InitFlags() {
	if flag.Lookup("grpc-province-server-address") == nil {
		flag.StringVar(
			&c.addr,
			"grpc-province-server-address",
			"localhost:3101",
			"province gRPC server address. Default: localhost:3101",
		)
	}
}

// Activate khởi tạo kết nối gRPC
func (c *ProvinceRPCComponent) Activate(_ sctx.ServiceContext) error {
	c.logger = sctx.GlobalLogger().GetLogger(c.id)
	c.logger.Infof("Connecting to Country gRPC server at %s...", c.addr)

	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return errors.Wrap(err, "failed to connect to Country gRPC server")
	}

	c.conn = conn
	c.client = pb.NewProvinceServiceClient(conn)
	c.isActive = true
	c.logger.Info("Country gRPC client activated successfully.")

	return nil
}

// Stop đóng kết nối khi service dừng
func (c *ProvinceRPCComponent) Stop() error {
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
func (c *ProvinceRPCComponent) GetClient() pb.ProvinceServiceClient {
	if !c.isActive {
		log.Fatal("CountryRPCComponent is not active")
	}
	return c.client
}

// GetCountryById ví dụ hàm tiện ích gọi gRPC
func (c *ProvinceRPCComponent) GetProvinceById(ctx context.Context, id int32) (*dto.ProvinceResponse, error) {
	resp, err := c.client.GetProvinceByIdHdl(ctx, &pb.GetProvinceByIdRequest{Id: id})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &dto.ProvinceResponse{
		Id:        int(resp.Id),
		Code:      resp.Code,
		Name:      resp.Name,
		CountryId: int(resp.CountryId),
	}, nil
}
