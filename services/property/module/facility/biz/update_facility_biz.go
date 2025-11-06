package facilitybiz

import (
	"context"
	"time"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *facilityBusiness) UpdateFacilityBiz(ctx context.Context, id int, data *model.FacilityUpdateDto) error {
	if _, err := b.repo.GetById(ctx, id); err != nil {
		return err;
	}
	
	now := time.Now()
	data.UpdatedAt = &now

	if err := b.repo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotUpdateFacility.Error()).
			WithDebug(err.Error())
	}

	return nil
}