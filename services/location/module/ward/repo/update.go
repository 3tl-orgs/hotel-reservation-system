package wardsrepo

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/pkg/errors"
)

func (w *wardRepo) Update(ctx context.Context, id int, data *wardsmodel.UpdateWardDTO) error {
	if err := w.db.Table(wardsmodel.UpdateWardDTO{}.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
