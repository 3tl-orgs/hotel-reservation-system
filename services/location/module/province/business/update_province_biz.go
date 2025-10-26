package provincebusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

func (p *provinceBusiness) UpdateProvinceBiz(ctx context.Context, id int, data *provincemodel.UpdateProvinceDTO) error {
	if err := p.provinceRepo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotUpdateProvince.Error())
	}

	timeUpdated := time.Now()
	data.UpdatedAt = &timeUpdated

	return nil
}
