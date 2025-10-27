package wardstransport

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
)

type WardBusiness interface {
	CreateWardBiz(ctx context.Context, data *wardsmodel.CreateWardDTO) error
	UpdateWardBiz(ctx context.Context, id int, data *wardsmodel.UpdateWardDTO) error
	DeleteWardBiz(ctx context.Context, id int) error
	GetWardByIdBiz(ctx context.Context, id int) (*wardsmodel.Ward, error)
	GetWardByCodeBiz(ctx context.Context, code string) (*wardsmodel.Ward, error)
}

type wardTransport struct {
	wardBusiness WardBusiness
}

func NewWardTransport(wardBusiness WardBusiness) *wardTransport {
	return &wardTransport{
		wardBusiness: wardBusiness,
	}
}
