package wardsbusiness

import (
	"context"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (w *wardBusiness) CreateWardBiz(ctx context.Context, data *wardsmodel.CreateWardDTO) error {
	// Checking exist Ward
	existingData, err := w.GetWardByCodeBiz(ctx, data.Code)
	if err != nil && !errors.Is(err, wardsmodel.ErrWardNotFound) {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrCannotCreateWard.Error())
	}
	if existingData != nil {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrIdDuplicated.Error())
	}

	// Checking exist Province
	_, err = w.provinceRepo.GetById(ctx, data.ProvinceId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrConflict.
				WithError(provincemodel.ErrProvinceNotFound.Error())
		}
		return core.ErrInternalServerError.
			WithError(err.Error())
	}

	if err := w.wardRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrCannotCreateWard.Error())
	}
	return nil
}
