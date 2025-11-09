package amenitybiz

import (
	"context"

	amenitymodel "github.com/ngleanhvu/go-booking/services/property/module/amenity/model"
	"github.com/ngleanhvu/go-booking/shared/dto"
)

type AmenityRepo interface {
	Create(ctx context.Context, data *amenitymodel.AmenityCreateDto) error
	GetByIds(ctx context.Context, ids []int) ([]amenitymodel.Amenity, error)
	GetById(ctx context.Context, id int) (*amenitymodel.Amenity, error)
	Update(ctx context.Context, id int, data *amenitymodel.AmenityUpdateDto) error
	Delete(ctx context.Context, id int) error
	DeleteMany(ctx context.Context, ids []int) error
}

type CountryRepo interface {
	GetCountryById(ctx context.Context, id int32) (*dto.CountryResponse, error)
}

type business struct {
	repo        AmenityRepo
	countryRepo CountryRepo
}

func NewBusiness(repo AmenityRepo, countryRepo CountryRepo) *business {
	return &business{repo: repo, countryRepo: countryRepo}
}
