package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"time"
)

func (w *wardBusiness) UpdateWardBiz(ctx context.Context, id int, data *wardsmodel.UpdateWardDTO) error {
	existingData, err := w.GetWardByIdBiz(ctx, id)
	if err != nil {
		return wardsmodel.ErrWardNotFound
	}

	if existingData.Id != id {
		return wardsmodel.ErrWardNotFound
	}

	if err := w.wardRepo.Update(ctx, id, data); err != nil {
		return core.ErrBadRequest.
			WithError(wardsmodel.ErrCannotUpdateWard.Error())
	}
	updatedTime := time.Now()
	data.UpdatedAt = &updatedTime

	data.Mask()
	return nil
}
