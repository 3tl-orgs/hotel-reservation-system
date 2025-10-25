package business

import (
	"context"
	"time"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) UpdateCountryBiz(ctx context.Context, id int, data *model.CountryUpdateDto) error {

	existingData, _ := biz.GetCountryByIdBiz(ctx, id)

	if existingData.Id != id {
		return core.ErrInternalServerError.
			WithError(model.ErrCountryCodeIsDuplicated.Error())
	}

	now := time.Now()
	data.UpdatedAt = &now

	if err := biz.countryRepo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotUpdateCountry.Error())
	}

	return nil
}
