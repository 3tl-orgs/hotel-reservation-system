package amenitybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) GetAmenityByIdsBiz(ctx context.Context, ids []int) ([]amenitymodel.Amenity, error) {
	data, err := b.repo.GetByIds(ctx, ids)
	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(amenitymodel.ErrCannotGetAmenity.Error()).
			WithDebug(err.Error())
	}
	return data, nil
}
