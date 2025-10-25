package biz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) GetAmenityByIdsBiz(ctx context.Context, ids []int) ([]model.Amenity, error) {
	data, err := b.repo.GetByIds(ctx, ids)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotGetAmenity.Error()).
			WithDebug(err.Error())
	}
	return data, nil
}
