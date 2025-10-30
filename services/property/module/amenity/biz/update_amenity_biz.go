package biz

import (
	"context"
	"time"

	"github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) UpdateAmenityBiz(ctx context.Context, id int, data *model.AmenityUpdateDto) error {
	if _, err := biz.GetAmenityByIdBiz(ctx, id); err != nil {
		return err;
	}
	
	now := time.Now()
	data.UpdatedAt = &now

	if err := biz.repo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotUpdateAmenity.Error()).
			WithDebug(err.Error())
	}

	return nil
}