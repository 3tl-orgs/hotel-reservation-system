package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (w *wardBusiness) DeleteWardBiz(ctx context.Context, id int) error {
	if err := w.wardRepo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(wardsmodel.ErrCannotDeleteWard.Error())
	}

	return nil
}
