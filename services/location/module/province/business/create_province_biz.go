package provincebusiness

import (
	"context"
	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	provincemodel "github.com/ngleanhvu/go-booking/services/location/module/province/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (p *provinceBusiness) CreateProvinceBiz(ctx context.Context, data *provincemodel.CreateProvinceDTO) error {
	// Checking for exist Province
	existingData, err := p.GetProvinceByCodeBiz(ctx, data.Code)

	if err != nil && !errors.Is(err, provincemodel.ErrProvinceNotFound) {
		return err
	}
	if existingData != nil && existingData.Status == true {
		return core.ErrInternalServerError.
			WithError(model.ErrCountryCodeIsDuplicated.Error())
	}

	// Checking for exist Country
	_, err = p.countryRepo.GetById(ctx, data.CountryId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return core.ErrConflict.
				WithError(model.ErrCountryNotFound.Error())
		}
		return core.ErrInternalServerError.
			WithError(err.Error())
	}

	if err := p.provinceRepo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(provincemodel.ErrCannotCreateProvince.Error()).
			WithDebug(err.Error())
	}
	return nil
}
