package wardsrepo

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/pkg/errors"
)

func (w *wardRepo) Create(ctx context.Context, data *wardsmodel.CreateWardDTO) error {
	if err := w.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
