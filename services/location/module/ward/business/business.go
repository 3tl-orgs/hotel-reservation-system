package wardsbusiness

import (
	"context"
	provincebusiness "github.com/ngleanhvu/go-booking/services/location/module/province/business"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type WardRepo interface {
	Create(ctx context.Context, data *wardsmodel.CreateWardDTO) error
	Update(ctx context.Context, id int, data *wardsmodel.UpdateWardDTO) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*wardsmodel.Ward, error)
	GetByCode(ctx context.Context, code string) (*wardsmodel.Ward, error)
	ListData(
		ctx context.Context,
		filter *wardsmodel.Filter,
		paging *core.Paging,
		moreKeys ...string,
	) ([]wardsmodel.Ward, error)
}

type wardBusiness struct {
	wardRepo     WardRepo
	provinceRepo provincebusiness.ProvinceRepo
}

func NewWardBusiness(wardRepo WardRepo, provinceRepo provincebusiness.ProvinceRepo) *wardBusiness {
	return &wardBusiness{
		wardRepo:     wardRepo,
		provinceRepo: provinceRepo,
	}
}
