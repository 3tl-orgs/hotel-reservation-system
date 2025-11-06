package facilitybiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *facilityBusiness) DeleteFacilityBiz(ctx context.Context, id int) error {
	if err := b.repo.Delete(ctx, id); err != nil {
		return err
	}

	if err := b.repo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError. 
			WithError(model.ErrCannotDeleteFacility.Error()). 
			WithDebug(err.Error())
	}

	return nil
}