package provincebusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (p *provinceBusiness) DeleteProvinceBiz(ctx context.Context, id int) error {
	oldData, err := p.GetProvinceByIdBiz(ctx, id)

	if err != nil {
		return provincemodel.ErrProvinceNotFound
	}

	if oldData.Status == false {
		return provincemodel.ErrProvinceHasBeenDeleted
	}

	if err := p.provinceRepo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotDeleteProvince.Error())
	}

	return nil
}
