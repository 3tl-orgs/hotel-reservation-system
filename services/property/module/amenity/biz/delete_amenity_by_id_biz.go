package amenitybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *business) DeleteAmenityByIdBiz(ctx context.Context, id int) error {
	if _, err := b.GetAmenityByIdBiz(ctx, id); err != nil {
		return err
	}

	if err := b.repo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(amenitymodel.ErrCannotDeleteAmenity.Error()).
			WithDebug(err.Error())
	}

	return nil
}
