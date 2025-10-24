package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) UpdateCountryBiz(ctx context.Context, id int, data *model.CountryUpdateDto) error {
	existingData, err := biz.countryRepo.GetByCode(ctx, *data.Code)

	if err != nil {
		return err
	}

	if existingData != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCountryCodeIsDuplicated.Error())
	}

	if err := biz.countryRepo.Update(ctx, id, data); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotUpdateCountry.Error())
	}

	return nil
}
