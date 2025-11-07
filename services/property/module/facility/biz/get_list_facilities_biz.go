package facilitybiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *facilityBusiness) GetListFacilitiesBiz(ctx context.Context, paging *core.Paging, moreKeys ...string) ([]model.Facility, error) {
	result, err := b.repo.List(ctx, paging, moreKeys...)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotReadFacility.Error()).
			WithDebug(err.Error())
	}

	return result, nil
}