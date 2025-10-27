package wardsrepo

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/pkg/errors"
)

func (w *wardRepo) Delete(ctx context.Context, id int) error {
	if err := w.db.Table(wardsmodel.Ward{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": false}).
		Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
