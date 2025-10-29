package wardsbusiness

import (
	"context"
	wardsmodel "github.com/ngleanhvu/go-booking/services/location/module/ward/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (w *wardBusiness) ListWardBiz(ctx context.Context,
	filter *wardsmodel.Filter,
	paging *core.Paging,
	moreKeys ...string,
) ([]wardsmodel.Ward, error) {
	result, err := w.wardRepo.ListData(ctx, filter, paging, moreKeys...)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(wardsmodel.ErrCannotGetWard.Error()).
			WithDebug(err.Error())
	}
	for i := range result {
		result[i].Mask()
	}
	
	return result, nil
}
