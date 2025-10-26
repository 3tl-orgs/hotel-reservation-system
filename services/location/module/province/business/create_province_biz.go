package provincebusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceBusiness) CreateProvinceBiz(ctx context.Context, data *provincemodel.CreateProvinceDTO) error {
	if err := p.provinceRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotCreateProvince.Error()).
			WithDebug(err.Error())
	}
	return nil
}
