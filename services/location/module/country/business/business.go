package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
)

type CountryRepo interface {
	Create(ctx context.Context, data *model.CountryCreateDto) error
}

type business struct {
	countryRepo CountryRepo
}

func NewBusiness(countryRepo CountryRepo) *business {
	return &business{countryRepo: countryRepo}
}
