package facilitybiz

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (b *facilityBusiness) GetFacilityByIdBiz(ctx context.Context, id int) (*model.Facility, error) {
	data, err := b.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound.
				WithError(model.ErrCannotReadFacility.Error())
		}

		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotReadFacility.Error())
	}

	return data, nil
}