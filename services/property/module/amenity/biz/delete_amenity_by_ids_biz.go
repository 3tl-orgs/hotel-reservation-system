package amenitybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) DeleteAmenityByIdsBiz(ctx context.Context, ids []int) error {
	if _, err := biz.GetAmenityByIdsBiz(ctx, ids); err != nil {
		return err
	}

	// for _, id := range ids {
	// 	if err := biz.repo.Delete(ctx, id); err != nil {
	// 		return core.ErrInternalServerError.
	// 			WithError(model.ErrCannotDeleteAmenity.Error()).
	// 			WithDebug(fmt.Sprintf("failed to delete amenity id=%d: %v", id, err.Error()))
	// 	}
	// }

	if err := biz.repo.DeleteMany(ctx, ids); err != nil {
		return core.ErrInternalServerError.
			WithError(amenitymodel.ErrCannotDeleteManyAmenities.Error()).
			WithDebug(err.Error())
	}

	return nil
}
