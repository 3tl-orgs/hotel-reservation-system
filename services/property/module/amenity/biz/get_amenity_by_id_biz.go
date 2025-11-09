package amenitybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) GetAmenityByIdBiz(ctx context.Context, id int) (*amenitymodel.Amenity, error) {
	data, err := b.repo.GetById(ctx, id)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(amenitymodel.ErrCannotGetAmenity.Error()).
			WithDebug(err.Error())
	}

	if data == nil {
		return nil, core.ErrInternalServerError.
			WithError(amenitymodel.ErrAmenityNotFound.Error())
	}

	return data, nil
}
