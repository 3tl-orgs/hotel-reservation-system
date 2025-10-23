package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (biz *business) GetCountryByCodeBiz(ctx context.Context, code string) (*model.Country, error) {
	data, err := biz.countryRepo.GetByCode(ctx, code)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound.
				WithError(model.ErrCountryNotFound.Error()).
				WithDebug(err.Error())
		}
		return nil, core.ErrInternalServerError.
			WithError(model.ErrCannotGetCountry.Error()).
			WithDebug(err.Error())
	}

	if data.Status == core.EntityDeleted {
		return nil, core.ErrNotFound.WithError(model.ErrCountryNotFound.Error())
	}

	return data, nil
}
