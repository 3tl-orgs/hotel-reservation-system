package provincebusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (p *provinceBusiness) GetProvinceByCodeBiz(ctx context.Context, code string) (*provincemodel.Province, error) {
	data, err := p.provinceRepo.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, provincemodel.ErrProvinceNotFound) {
			return nil, provincemodel.ErrProvinceNotFound
		}
		return nil, core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotGetProvince.Error())
	}

	data.Mask()
	return data, err
}
