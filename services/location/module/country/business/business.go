package business

import (
	"context"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type CountryRepo interface {
	Create(ctx context.Context, data *model.CountryCreateDto) error
	GetById(ctx context.Context, id int) (*model.Country, error)
	GetByCode(ctx context.Context, code string) (*model.Country, error)
	Update(ctx context.Context, id int, data *model.CountryUpdateDto) error
	List(ctx context.Context, filter *model.Filter, paging *core.Paging, moreKey ...string) ([]model.Country, error)
}

type business struct {
	countryRepo CountryRepo
}

func NewBusiness(countryRepo CountryRepo) *business {
	return &business{countryRepo: countryRepo}
}
