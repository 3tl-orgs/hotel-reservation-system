package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) DeleteCountryBiz(ctx context.Context, id int) error {
	_, err := biz.GetCountryByIdBiz(ctx, id)

	if err != nil {
		return err
	}

	if err := biz.countryRepo.Delete(ctx, id); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotDeleteCountry.Error()).
			WithDebug(err.Error())
	}

	return nil
}
