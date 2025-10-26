package provincerepo

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/pkg/errors"
)

func (p *provinceRepo) Create(ctx context.Context, data *provincemodel.CreateProvinceDTO) error {
	if err := p.db.Table(provincemodel.CreateProvinceDTO{}.TableName()).
		Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
