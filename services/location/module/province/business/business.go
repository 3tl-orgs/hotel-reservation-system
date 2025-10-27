package provincebusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type ProvinceRepo interface {
	Create(ctx context.Context, data *provincemodel.CreateProvinceDTO) error
	Update(ctx context.Context, id int, data *provincemodel.UpdateProvinceDTO) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (*provincemodel.Province, error)
	GetByCode(ctx context.Context, code string) (*provincemodel.Province, error)
	ListData(ctx context.Context,
		filter *provincemodel.Filter,
		paging *core.Paging,
		moreKeys ...string) ([]provincemodel.Province, error)
}

type provinceBusiness struct {
	provinceRepo ProvinceRepo
}

func NewProvinceBusiness(provinceRepo ProvinceRepo) *provinceBusiness {
	return &provinceBusiness{provinceRepo: provinceRepo}
}
