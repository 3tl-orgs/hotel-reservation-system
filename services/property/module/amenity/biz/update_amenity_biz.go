package amenitybiz

import (
	"context"
	"time"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) UpdateAmenityBiz(ctx context.Context, id int, data *amenitymodel.AmenityUpdateDto) error {
	if _, err := biz.GetAmenityByIdBiz(ctx, id); err != nil {
		return err
	}

	now := time.Now()
	data.UpdatedAt = &now

	if err := biz.repo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(amenitymodel.ErrCannotUpdateAmenity.Error()).
			WithDebug(err.Error())
	}

	return nil
}
