package amenitybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) CreateAmenityBiz(ctx context.Context,
	data *amenitymodel.AmenityCreateDto) error {
	if err := b.repo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(amenitymodel.ErrCannotCreateAmenity.Error()).
			WithDebug(err.Error())
	}
	return nil
}
