package provincetransport

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
)

type ProvinceBusiness interface {
	CreateProvinceBiz(ctx context.Context, data *provincemodel.CreateProvinceDTO) error
	UpdateProvinceBiz(ctx context.Context, id int, data *provincemodel.UpdateProvinceDTO) error
	DeleteProvinceBiz(ctx context.Context, id int) error
	GetProvinceByIdBiz(ctx context.Context, id int) (*provincemodel.Province, error)
	GetProvinceByCodeBiz(ctx context.Context, code string) (*provincemodel.Province, error)
}

type provinceTransport struct {
	provinceBusiness ProvinceBusiness
}

func NewProvinceTransport(provinceBusiness ProvinceBusiness) *provinceTransport {
	return &provinceTransport{
		provinceBusiness: provinceBusiness,
	}
}
