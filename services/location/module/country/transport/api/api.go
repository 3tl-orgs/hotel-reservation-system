package api

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
)

type Business interface {
	CreateCountryBiz(ctx context.Context, data *model.CountryCreateDto) error
}

type api struct {
	business Business
}

func NewApi(business Business) *api {
	return &api{
		business: business,
	}
}
