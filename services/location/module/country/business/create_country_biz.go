package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (s *business) CreateCountryBiz(ctx context.Context, data *model.CountryCreateDto) error {

	existingData, _ := s.GetCountryByCodeBiz(ctx, data.Code)

	if existingData != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCountryCodeIsDuplicated.Error())
	}

	if err := s.countryRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotCreateCountry.Error()).
			WithDebug(err.Error())
	}

	return nil
}
