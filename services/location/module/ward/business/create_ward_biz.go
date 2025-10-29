package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (w *wardBusiness) CreateWardBiz(ctx context.Context, data *wardsmodel.CreateWardDTO) error {
	existingData, err := w.GetWardByCodeBiz(ctx, data.Code)
	if err != nil && !errors.Is(err, wardsmodel.ErrWardNotFound) {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrCannotCreateWard.Error())
	}
	if existingData != nil {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrIdDuplicated.Error())
	}

	if err := w.wardRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrCannotCreateWard.Error())
	}
	return nil
}
