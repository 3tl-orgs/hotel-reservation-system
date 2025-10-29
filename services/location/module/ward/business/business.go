package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
)

type WardRepo interface {
	Create(ctx context.Context, data *wardsmodel.CreateWardDTO) error
	Update(ctx context.Context, id int, data *wardsmodel.UpdateWardDTO) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*wardsmodel.Ward, error)
	GetByCode(ctx context.Context, code string) (*wardsmodel.Ward, error)
}

type wardBusiness struct {
	wardRepo WardRepo
}

func NewWardBusiness(wardRepo WardRepo) *wardBusiness {
	return &wardBusiness{wardRepo: wardRepo}
}
