package provincebusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (p *provinceBusiness) GetProvinceByIdBiz(ctx context.Context, id int) (*provincemodel.Province, error) {
	data, err := p.provinceRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound.
				WithError(provincemodel.ErrProvinceNotFound.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotGetProvince.Error())
	}

	return data, nil
}
