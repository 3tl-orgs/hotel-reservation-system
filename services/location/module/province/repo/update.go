package provincerepo

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/pkg/errors"
)

func (p *provinceRepo) Update(ctx context.Context, id int, data *provincemodel.UpdateProvinceDTO) error {
	if err := p.db.Table(provincemodel.UpdateProvinceDTO{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
