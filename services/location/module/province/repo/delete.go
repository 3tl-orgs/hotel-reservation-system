package provincerepo

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/pkg/errors"
)

func (p *provinceRepo) Delete(ctx context.Context, id int) error {
	if err := p.db.Table(provincemodel.Province{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": false}).
		Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
