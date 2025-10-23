package api

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type Business interface {
	CreateCountryBiz(ctx context.Context, data *model.CountryCreateDto) error
	GetCountryByIdBiz(ctx context.Context, id int) (*model.Country, error)
	GetCountryByCodeBiz(ctx context.Context, code string) (*model.Country, error)
	UpdateCountryBiz(ctx context.Context, id int, data *model.CountryUpdateDto) error
	ListCountryBiz(ctx context.Context,
		filter *model.Filter,
		paging *core.Paging,
		moreKey ...string) ([]model.Country, error)
}

type api struct {
	business Business
}

func NewApi(business Business) *api {
	return &api{
		business: business,
	}
}
