package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (w *wardBusiness) GetWardByCodeBiz(ctx context.Context, code string) (*wardsmodel.Ward, error) {
	data, err := w.wardRepo.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, wardsmodel.ErrWardNotFound) {
			return nil, wardsmodel.ErrWardNotFound
		}
		return nil, core.ErrInternalServerError.
			WithError(wardsmodel.ErrWardNotFound.Error())
	}

	data.Mask()
	return data, nil
}
