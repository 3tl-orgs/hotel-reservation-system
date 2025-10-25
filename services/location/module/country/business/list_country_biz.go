package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (biz *business) ListCountryBiz(ctx context.Context,
	filter *model.Filter,
	paging *core.Paging,
	moreKey ...string) ([]model.Country, error) {

	result, err := biz.countryRepo.List(ctx, filter, paging, moreKey...)

	if err != nil {
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotGetCountry.Error()).
			WithDebug(err.Error())
	}

	for i := range result {
		result[i].Mask()
	}

	return result, nil
}
