package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (w *wardBusiness) GetWardByIdBiz(ctx context.Context, id int) (*wardsmodel.Ward, error) {
	data, err := w.wardRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, wardsmodel.ErrWardNotFound) {
			return nil, core.ErrNotFound.
				WithError(wardsmodel.ErrWardNotFound.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(wardsmodel.ErrWardNotFound.Error())
	}

	return data, nil
}
